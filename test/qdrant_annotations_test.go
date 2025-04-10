package test

import (
	v1 "k8s.io/api/core/v1"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	networkingv1 "k8s.io/api/networking/v1"
)

func TestPvcAnnotations(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"persistence.annotations": `{"test": "value"}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, output, &statefulSet)

	require.Contains(t, statefulSet.Spec.VolumeClaimTemplates[0].Annotations, "test")
	require.Equal(t, statefulSet.Spec.VolumeClaimTemplates[0].Annotations["test"], "value")
}

func TestIngressAnnotations(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"ingress.enabled":     `true`,
			"ingress.annotations": `{"test": "value"}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/ingress.yaml"})

	var ingress networkingv1.Ingress
	helm.UnmarshalK8SYaml(t, output, &ingress)

	require.Contains(t, ingress.Annotations, "test")
	require.Equal(t, ingress.Annotations["test"], "value")
}

func TestServiceAccountAnnotations(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"serviceAccount.annotations": `{"test": "value"}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/serviceaccount.yaml"})

	var serviceAccount v1.ServiceAccount
	helm.UnmarshalK8SYaml(t, output, &serviceAccount)

	require.Contains(t, serviceAccount.Annotations, "test")
	require.Equal(t, serviceAccount.Annotations["test"], "value")
}
