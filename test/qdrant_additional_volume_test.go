package test

import (
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

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"additionalVolumes":      `[{"name":"volumeName","emptyDir":{"sizeLimit":"500Mi"}}]`,
			"additionalVolumeMounts": `[{"name":"volumeName","mountPath":"/mount/path"}]`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

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
}
