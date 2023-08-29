package test

import (
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/stretchr/testify/require"
	"path/filepath"
	"strings"
	"testing"
)

func TestRelabelings(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"metrics.serviceMonitor.enabled":     `true`,
			"metrics.serviceMonitor.relabelings": `[{"sourceLabels": ["source"], "targetLabel": "target"}]`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/servicemonitor.yaml"})

	var serviceMonitor monitoringv1.ServiceMonitor
	helm.UnmarshalK8SYaml(t, output, &serviceMonitor)

	require.Equal(t, "source", string(serviceMonitor.Spec.Endpoints[0].RelabelConfigs[0].SourceLabels[0]))
	require.Equal(t, "target", serviceMonitor.Spec.Endpoints[0].RelabelConfigs[0].TargetLabel)
}

func TestMetricRelabelings(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"metrics.serviceMonitor.enabled":           `true`,
			"metrics.serviceMonitor.metricRelabelings": `[{"sourceLabels": ["source"], "action": "drop"}]`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/servicemonitor.yaml"})

	var serviceMonitor monitoringv1.ServiceMonitor
	helm.UnmarshalK8SYaml(t, output, &serviceMonitor)

	require.Equal(t, "source", string(serviceMonitor.Spec.Endpoints[0].MetricRelabelConfigs[0].SourceLabels[0]))
	require.Equal(t, "drop", serviceMonitor.Spec.Endpoints[0].MetricRelabelConfigs[0].Action)
}
