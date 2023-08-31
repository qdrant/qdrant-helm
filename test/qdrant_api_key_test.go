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

func TestStringApiKey(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"apiKey": `"test_api_key"`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, output, &statefulSet)

	output = helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/secret.yaml"})

	var secret corev1.Secret
	helm.UnmarshalK8SYaml(t, output, &secret)

	container, _ := lo.Find(statefulSet.Spec.Template.Spec.Containers, func(container corev1.Container) bool {
		return container.Name == "qdrant"
	})

	secretVolumeMount, hasSecretVolumeMount := lo.Find(container.VolumeMounts, func(volumeMount corev1.VolumeMount) bool {
		return volumeMount.Name == "qdrant-secret"
	})

	_, hasSecretVolume := lo.Find(statefulSet.Spec.Template.Spec.Volumes, func(volume corev1.Volume) bool {
		return volume.Name == secretVolumeMount.Name
	})

	require.Equal(t, hasSecretVolumeMount, true)
	require.Equal(t, hasSecretVolume, true)
	require.Contains(t, lo.Keys(secret.Data), "local.yaml")
	require.Contains(t, lo.Keys(secret.Data), "api-key")

	require.Equal(t, "service:\n  api_key: test_api_key", string(secret.Data["local.yaml"]))
}

func TestRandomApiKey(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"apiKey": `true`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, output, &statefulSet)

	output = helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/secret.yaml"})

	var secret corev1.Secret
	helm.UnmarshalK8SYaml(t, output, &secret)

	container, _ := lo.Find(statefulSet.Spec.Template.Spec.Containers, func(container corev1.Container) bool {
		return container.Name == "qdrant"
	})

	secretVolumeMount, hasSecretVolumeMount := lo.Find(container.VolumeMounts, func(volumeMount corev1.VolumeMount) bool {
		return volumeMount.Name == "qdrant-secret"
	})

	_, hasSecretVolume := lo.Find(statefulSet.Spec.Template.Spec.Volumes, func(volume corev1.Volume) bool {
		return volume.Name == secretVolumeMount.Name
	})

	require.Equal(t, hasSecretVolumeMount, true)
	require.Equal(t, hasSecretVolume, true)
	require.Contains(t, lo.Keys(secret.Data), "local.yaml")
	require.Contains(t, lo.Keys(secret.Data), "api-key")

	require.Regexp(t, "^service:\n  api_key: [a-zA-Z0-9]+$", string(secret.Data["local.yaml"]))
}

func TestNoApiKey(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues:  map[string]string{},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, output, &statefulSet)

	container, _ := lo.Find(statefulSet.Spec.Template.Spec.Containers, func(container corev1.Container) bool {
		return container.Name == "qdrant"
	})

	_, hasSecretVolumeMount := lo.Find(container.VolumeMounts, func(volumeMount corev1.VolumeMount) bool {
		return volumeMount.Name == "qdrant-secret"
	})

	_, hasSecretVolume := lo.Find(statefulSet.Spec.Template.Spec.Volumes, func(volume corev1.Volume) bool {
		return volume.Name == "qdrant-secret"
	})

	require.Equal(t, hasSecretVolumeMount, false)
	require.Equal(t, hasSecretVolume, false)
}
