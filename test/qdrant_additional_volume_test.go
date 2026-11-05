package test

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func TestAdditionalVolumeAndVolumeMount(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueID())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJSONValues: map[string]string{
			"additionalVolumes":      `[{"name":"volumeName","emptyDir":{"sizeLimit":"500Mi"}}]`,
			"additionalVolumeMounts": `[{"name":"volumeName","mountPath":"/mount/path"}]`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, output, &statefulSet)

	container, _ := lo.Find(statefulSet.Spec.Template.Spec.Containers, func(container corev1.Container) bool {
		return container.Name == "qdrant"
	})

	volumeMount, hasAdditionalVolumeMount := lo.Find(container.VolumeMounts, func(volumeMount corev1.VolumeMount) bool {
		return volumeMount.Name == "volumeName"
	})

	volume, hasAdditionalVolume := lo.Find(statefulSet.Spec.Template.Spec.Volumes, func(volume corev1.Volume) bool {
		return volume.Name == "volumeName"
	})

	require.Equal(t, hasAdditionalVolumeMount, true)
	require.Equal(t, hasAdditionalVolume, true)

	require.Equal(t, "/mount/path", volumeMount.MountPath)
	require.Equal(t, "500Mi", volume.EmptyDir.SizeLimit.String())
	require.NotContains(t, statefulSet.Spec.Template.Spec.InitContainers[0].Command, "/mount/path")
	require.Len(t, statefulSet.Spec.Template.Spec.InitContainers[0].VolumeMounts, 2)
}

func TestAdditionalVolumeOwnership(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	require.NoError(t, err)

	options := &helm.Options{
		SetJSONValues: map[string]string{
			"additionalVolumes": `[
				{"name":"snapshots","nfs":{"server":"nfs.example.com","path":"/"}},
				{"name":"credentials","secret":{"secretName":"credentials"}},
				{"name":"untouched","emptyDir":{}}
			]`,
			"additionalVolumeMounts": `[
				{"name":"snapshots","mountPath":"/qdrant/efs snapshots","subPath":"snapshots","updateFsOwnership":true},
				{"name":"credentials","mountPath":"/credentials","readOnly":true,"updateFsOwnership":true},
				{"name":"untouched","mountPath":"/untouched","updateFsOwnership":false}
			]`,
		},
	}

	output := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, "qdrant", []string{"templates/statefulset.yaml"})
	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, output, &statefulSet)

	initContainer := statefulSet.Spec.Template.Spec.InitContainers[0]
	require.Equal(t, []string{"chown", "-R", "1000:3000", "/qdrant/storage", "/qdrant/snapshots", "/qdrant/efs snapshots"}, initContainer.Command)
	require.Len(t, initContainer.VolumeMounts, 3)
	selectedMount := corev1.VolumeMount{
		Name: "snapshots", MountPath: "/qdrant/efs snapshots", SubPath: "snapshots",
	}
	require.Contains(t, initContainer.VolumeMounts, selectedMount)

	container, found := lo.Find(statefulSet.Spec.Template.Spec.Containers, func(container corev1.Container) bool {
		return container.Name == "qdrant"
	})
	require.True(t, found)
	require.Contains(t, container.VolumeMounts, selectedMount)
	require.Contains(t, container.VolumeMounts, corev1.VolumeMount{
		Name: "credentials", MountPath: "/credentials", ReadOnly: true,
	})
	require.Contains(t, container.VolumeMounts, corev1.VolumeMount{
		Name: "untouched", MountPath: "/untouched",
	})
	// This chart option must never leak into Kubernetes VolumeMount objects.
	require.NotContains(t, output, "updateFsOwnership")

	testOutput := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, "qdrant", []string{"templates/tests/test-db-interaction.yaml"})
	var testPods []corev1.Pod
	helm.UnmarshalK8SYamls(t, testOutput, &testPods, func(pod corev1.Pod) bool {
		return pod.Kind == "Pod"
	})
	require.Len(t, testPods, 1)
	require.Len(t, testPods[0].Spec.Containers, 1)
	require.Contains(t, testPods[0].Spec.Containers[0].VolumeMounts, selectedMount)
	require.NotContains(t, testOutput, "updateFsOwnership")
}

func TestAdditionalVolumeOwnershipRespectsSecuritySettings(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	require.NoError(t, err)

	for name, settings := range map[string]map[string]string{
		"ownership disabled": {"updateVolumeFsOwnership": "false"},
		"unprivileged image": {"image.useUnprivilegedImage": "true"},
		"root user":          {"containerSecurityContext.runAsUser": "0"},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			settings["additionalVolumes"] = `[{"name":"snapshots","emptyDir":{}}]`
			settings["additionalVolumeMounts"] = `[{"name":"snapshots","mountPath":"/qdrant/extra-snapshots","updateFsOwnership":true}]`
			options := &helm.Options{SetJSONValues: settings}
			output := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, "qdrant", []string{"templates/statefulset.yaml"})
			var statefulSet appsv1.StatefulSet
			helm.UnmarshalK8SYaml(t, output, &statefulSet)
			require.Empty(t, statefulSet.Spec.Template.Spec.InitContainers)
			require.NotContains(t, output, "updateFsOwnership")
		})
	}
}

func TestAdditionalVolumeOwnershipRejectsSubPathExpr(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	require.NoError(t, err)
	options := &helm.Options{
		SetJSONValues: map[string]string{
			"additionalVolumes":      `[{"name":"snapshots","emptyDir":{}}]`,
			"additionalVolumeMounts": `[{"name":"snapshots","mountPath":"/snapshots","subPathExpr":"$(POD_NAME)","updateFsOwnership":true}]`,
		},
	}
	_, err = helm.RenderTemplateE(t, options, helmChartPath, "qdrant", []string{"templates/statefulset.yaml"})
	require.ErrorContains(t, err, "updateFsOwnership does not support subPathExpr")
}
