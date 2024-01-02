package test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
)

func TestDefaultProbesOnStatefulset(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	tests := []struct {
		version           string
		readynessEndpoint string
	}{
		{
			version:           "v1.6.0",
			readynessEndpoint: "/",
		},
		{
			version:           "v1.7.3",
			readynessEndpoint: "/readyz",
		},
		{
			version:           "v1.7.4",
			readynessEndpoint: "/readyz",
		},
	}

	for _, test := range tests {
		t.Run("/readyz with "+test.version, func(t *testing.T) {
			options := &helm.Options{
				SetValues: map[string]string{
					"image.tag":                  test.version,
					"image.repository":           "test/repo",
					"image.useUnprivilegedImage": "true",
				},
				KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
			}

			output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

			var statefulSet appsv1.StatefulSet
			helm.UnmarshalK8SYaml(t, output, &statefulSet)

			container, _ := lo.Find(statefulSet.Spec.Template.Spec.Containers, func(container corev1.Container) bool {
				return container.Name == "qdrant"
			})

			require.Empty(t, container.StartupProbe)
			require.Equal(t, test.readynessEndpoint, container.ReadinessProbe.HTTPGet.Path)
			require.Empty(t, container.LivenessProbe)
		})
	}
}
