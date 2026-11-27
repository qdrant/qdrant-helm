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
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
)

const dbInteractionTemplate = "templates/tests/test-db-interaction.yaml"

func renderDBInteractionTemplate(t *testing.T, values map[string]string) (string, error) {
	t.Helper()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueID())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJSONValues:  values,
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	return helm.RenderTemplateE(t, options, helmChartPath, "qdrant", []string{dbInteractionTemplate})
}

func unmarshalDBInteractionResources(t *testing.T, output string) (corev1.Pod, corev1.ConfigMap) {
	t.Helper()

	var pod corev1.Pod
	var configMap corev1.ConfigMap
	var foundPod, foundConfigMap bool

	for _, doc := range strings.Split(output, "\n---\n") {
		switch {
		case strings.Contains(doc, "kind: Pod"):
			helm.UnmarshalK8SYaml(t, doc, &pod)
			foundPod = true
		case strings.Contains(doc, "kind: ConfigMap"):
			helm.UnmarshalK8SYaml(t, doc, &configMap)
			foundConfigMap = true
		}
	}

	require.True(t, foundPod, "expected the test Pod to be rendered")
	require.True(t, foundConfigMap, "expected the test ConfigMap to be rendered")

	return pod, configMap
}

func TestChartTestsDBInteractionEnabledByDefault(t *testing.T) {
	t.Parallel()

	output, err := renderDBInteractionTemplate(t, nil)
	require.NoError(t, err)

	pod, configMap := unmarshalDBInteractionResources(t, output)

	require.Equal(t, "qdrant-test-db-interaction", pod.Name)
	require.Equal(t, "test", pod.Annotations["helm.sh/hook"])
	require.Equal(t, "qdrant-test-db-interaction", configMap.Name)
	require.Equal(t, "test", configMap.Annotations["helm.sh/hook"])
	require.Contains(t, configMap.Data, "entrypoint.sh")
}

func TestChartTestsDBInteractionExplicitlyEnabled(t *testing.T) {
	t.Parallel()

	output, err := renderDBInteractionTemplate(t, map[string]string{
		"chartTests.dbInteraction.enabled": `true`,
	})
	require.NoError(t, err)

	pod, _ := unmarshalDBInteractionResources(t, output)
	require.Equal(t, "qdrant-test-db-interaction", pod.Name)
}

func TestChartTestsDBInteractionDisabled(t *testing.T) {
	t.Parallel()

	output, err := renderDBInteractionTemplate(t, map[string]string{
		"chartTests.dbInteraction.enabled": `false`,
	})

	// Helm reports a template that renders to nothing as missing.
	require.Error(t, err)
	require.Contains(t, err.Error(), "could not find template "+dbInteractionTemplate)
	require.NotContains(t, output, "kind: Pod")
	require.NotContains(t, output, "kind: ConfigMap")
}

func TestChartTestsDBInteractionDisabledRendersNoHooks(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueID())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJSONValues: map[string]string{
			"chartTests.dbInteraction.enabled": `false`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, "qdrant", nil)

	require.Contains(t, output, "kind: StatefulSet")
	require.NotContains(t, output, "test-db-interaction")
	require.NotContains(t, output, "helm.sh/hook")
}

func TestChartTestsDBInteractionOverrides(t *testing.T) {
	t.Parallel()

	output, err := renderDBInteractionTemplate(t, map[string]string{
		"chartTests.dbInteraction.image":                   `"example.com/test-image:1.2.3"`,
		"chartTests.dbInteraction.resources.limits.cpu":    `"250m"`,
		"chartTests.dbInteraction.resources.requests.cpu":  `"50m"`,
		"chartTests.dbInteraction.resources.limits.memory": `"300Mi"`,
	})
	require.NoError(t, err)

	pod, _ := unmarshalDBInteractionResources(t, output)

	require.Len(t, pod.Spec.Containers, 1)
	container := pod.Spec.Containers[0]
	require.Equal(t, "example.com/test-image:1.2.3", container.Image)
	require.Equal(t, "250m", container.Resources.Limits.Cpu().String())
	require.Equal(t, "300Mi", container.Resources.Limits.Memory().String())
	require.Equal(t, "50m", container.Resources.Requests.Cpu().String())
}
