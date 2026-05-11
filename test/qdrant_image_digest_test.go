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

const testDigest = "sha256:94728574965d17c6485dd361aa3c0818b325b9016dac5ea6afec7b4b2700865f"

func TestImageDigestPinning(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	tests := []struct {
		name              string
		setValues         map[string]string
		expectedImage     string
		readinessEndpoint string
	}{
		{
			name: "digest only renders repository@digest",
			setValues: map[string]string{
				"image.repository": "test/repo",
				"image.digest":     testDigest,
			},
			expectedImage:     "test/repo@" + testDigest,
			readinessEndpoint: "/readyz",
		},
		{
			name: "digest with tag still renders repository@digest",
			setValues: map[string]string{
				"image.repository": "test/repo",
				"image.tag":        "v1.17.1",
				"image.digest":     testDigest,
			},
			expectedImage:     "test/repo@" + testDigest,
			readinessEndpoint: "/readyz",
		},
		{
			name: "digest takes precedence over useUnprivilegedImage suffix",
			setValues: map[string]string{
				"image.repository":           "test/repo",
				"image.digest":               testDigest,
				"image.useUnprivilegedImage": "true",
			},
			expectedImage:     "test/repo@" + testDigest,
			readinessEndpoint: "/readyz",
		},
		{
			name: "tag@sha256 notation strips digest from semverCompare",
			setValues: map[string]string{
				"image.repository": "test/repo",
				"image.tag":        "v1.17.1@" + testDigest,
			},
			expectedImage:     "test/repo:v1.17.1@" + testDigest,
			readinessEndpoint: "/readyz",
		},
		{
			name: "tag@sha256 notation respects pre-1.7.3 readiness path",
			setValues: map[string]string{
				"image.repository": "test/repo",
				"image.tag":        "v1.6.0@" + testDigest,
			},
			expectedImage:     "test/repo:v1.6.0@" + testDigest,
			readinessEndpoint: "/",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
			logger.Log(t, "Namespace: %s\n", namespaceName)

			options := &helm.Options{
				SetValues:      test.setValues,
				KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
			}

			output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

			var statefulSet appsv1.StatefulSet
			helm.UnmarshalK8SYaml(t, output, &statefulSet)

			container, _ := lo.Find(statefulSet.Spec.Template.Spec.Containers, func(container corev1.Container) bool {
				return container.Name == "qdrant"
			})

			require.Equal(t, test.expectedImage, container.Image)
			require.Equal(t, test.readinessEndpoint, container.ReadinessProbe.HTTPGet.Path)
		})
	}
}
