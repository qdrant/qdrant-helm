package test

import (
	"path/filepath"
	"strings"
	"testing"

	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
)

func TestAdditionalLabelsAreSetOnStatefulset(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"additionalLabels":         `{"test": "additionalLabels"}`,
			"podLabels":                `{"test": "podLabels"}`,
			"service.additionalLabels": `{"test": "serviceAdditionalLabels"}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, output, &statefulSet)

	require.Contains(t, statefulSet.ObjectMeta.Labels, "test")
	require.Equal(t, "additionalLabels", statefulSet.ObjectMeta.Labels["test"])
	require.Contains(t, statefulSet.Spec.Template.ObjectMeta.Labels, "test")
	require.Equal(t, "podLabels", statefulSet.Spec.Template.ObjectMeta.Labels["test"])
}

func TestAdditionalLabelsAreSetOnService(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"service.additionalLabels": `{"test": "serviceAdditionalLabels"}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/service.yaml"})

	var service corev1.Service
	helm.UnmarshalK8SYaml(t, output, &service)

	require.Contains(t, service.ObjectMeta.Labels, "test")
	require.Equal(t, "serviceAdditionalLabels", service.ObjectMeta.Labels["test"])
}

func TestAdditionalLabelsAreSetOnServiceHeadless(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"service.additionalLabels": `{"test": "serviceAdditionalLabels"}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/service-headless.yaml"})

	var service corev1.Service
	helm.UnmarshalK8SYaml(t, output, &service)

	require.Contains(t, service.ObjectMeta.Labels, "test")
	require.Equal(t, "serviceAdditionalLabels", service.ObjectMeta.Labels["test"])
}

func TestAdditionalLabelsAreSetOnServiceMonitor(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"metrics.serviceMonitor.enabled":          `true`,
			"metrics.serviceMonitor.additionalLabels": `{"test": "serviceMonitorAdditionalLabels"}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/servicemonitor.yaml"})

	var serviceMonitor monitoringv1.ServiceMonitor
	helm.UnmarshalK8SYaml(t, output, &serviceMonitor)

	require.Contains(t, serviceMonitor.ObjectMeta.Labels, "test")
	require.Equal(t, "serviceMonitorAdditionalLabels", serviceMonitor.ObjectMeta.Labels["test"])
}

func TestAdditionalLabelsAreSetOnIngress(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"ingress.enabled":          `true`,
			"ingress.additionalLabels": `{"test": "ingressAdditionalLabels"}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/ingress.yaml"})

	var ingress networkingv1.Ingress
	helm.UnmarshalK8SYaml(t, output, &ingress)

	require.Contains(t, ingress.ObjectMeta.Labels, "test")
	require.Equal(t, "ingressAdditionalLabels", ingress.ObjectMeta.Labels["test"])
}
