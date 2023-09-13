package test

import (
	"github.com/ghodss/yaml"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
	helmv3 "helm.sh/helm/v3/pkg/chart"
	appsv1 "k8s.io/api/apps/v1"
)

func TestDefaultImage(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, output, &statefulSet)

	container, _ := lo.Find(statefulSet.Spec.Template.Spec.Containers, func(container corev1.Container) bool {
		return container.Name == "qdrant"
	})

	chartYaml, err := os.ReadFile("../charts/qdrant/Chart.yaml")
	if err != nil {
		log.Fatalf("unable to chart yaml file: %v", err)
	}
	var chart helmv3.Metadata
	err = yaml.Unmarshal(chartYaml, &chart)

	if err != nil {
		log.Fatalf("unable to decode chart yaml: %v", err)
	}
	require.Equal(t, "qdrant/qdrant:"+chart.AppVersion, container.Image)
}

func TestOverwriteImage(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetValues: map[string]string{
			"image.tag":                  "test-tag",
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

	require.Equal(t, "test/repo:test-tag-unprivileged", container.Image)
}
