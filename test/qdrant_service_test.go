package test

import (
	"k8s.io/apimachinery/pkg/util/intstr"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func TestNodePortsAreSetOnNodePortService(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"service": `{"type": "NodePort", "ports": [{"name": "http", "port": 6333, "targetPort": 6333, "protocol": "TCP", "nodePort": 30333}, {"name": "grpc", "port": 6334, "targetPort": 6334, "protocol": "TCP", "nodePort": 30334}]}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	statefulsetOutput := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, statefulsetOutput, &statefulSet)

	require.Equal(t, "http", statefulSet.Spec.Template.Spec.Containers[0].Ports[0].Name)
	require.EqualValues(t, 6333, statefulSet.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort)
	require.Equal(t, corev1.ProtocolTCP, statefulSet.Spec.Template.Spec.Containers[0].Ports[0].Protocol)
	require.Equal(t, "grpc", statefulSet.Spec.Template.Spec.Containers[0].Ports[1].Name)
	require.EqualValues(t, 6334, statefulSet.Spec.Template.Spec.Containers[0].Ports[1].ContainerPort)
	require.Equal(t, corev1.ProtocolTCP, statefulSet.Spec.Template.Spec.Containers[0].Ports[1].Protocol)

	serviceOutput := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/service.yaml"})

	var service corev1.Service
	helm.UnmarshalK8SYaml(t, serviceOutput, &service)

	require.Equal(t, corev1.ServiceType("NodePort"), service.Spec.Type)
	require.Equal(t, "http", service.Spec.Ports[0].Name)
	require.EqualValues(t, 6333, service.Spec.Ports[0].Port)
	require.EqualValues(t, intstr.IntOrString{IntVal: 6333}, service.Spec.Ports[0].TargetPort)
	require.EqualValues(t, 30333, service.Spec.Ports[0].NodePort)
	require.Equal(t, corev1.ProtocolTCP, service.Spec.Ports[0].Protocol)
	require.Equal(t, "grpc", service.Spec.Ports[1].Name)
	require.EqualValues(t, 6334, service.Spec.Ports[1].Port)
	require.EqualValues(t, intstr.IntOrString{IntVal: 6334}, service.Spec.Ports[1].TargetPort)
	require.EqualValues(t, 30334, service.Spec.Ports[1].NodePort)
	require.Equal(t, corev1.ProtocolTCP, service.Spec.Ports[1].Protocol)
}

func TestNodePortsAreDefaultsOnClusterIpService(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"service": `{"type": "ClusterIP", "ports": [{"name": "http", "port": 6333, "targetPort": 6333, "protocol": "TCP"}, {"name": "grpc", "port": 6334, "targetPort": 6334, "protocol": "TCP"}]}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	statefulSetOutput := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, statefulSetOutput, &statefulSet)

	require.Equal(t, "http", statefulSet.Spec.Template.Spec.Containers[0].Ports[0].Name)
	require.EqualValues(t, 6333, statefulSet.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort)
	require.Equal(t, corev1.ProtocolTCP, statefulSet.Spec.Template.Spec.Containers[0].Ports[0].Protocol)
	require.Equal(t, "grpc", statefulSet.Spec.Template.Spec.Containers[0].Ports[1].Name)
	require.EqualValues(t, 6334, statefulSet.Spec.Template.Spec.Containers[0].Ports[1].ContainerPort)
	require.Equal(t, corev1.ProtocolTCP, statefulSet.Spec.Template.Spec.Containers[0].Ports[1].Protocol)

	serviceOutput := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/service.yaml"})

	var service corev1.Service
	helm.UnmarshalK8SYaml(t, serviceOutput, &service)

	require.Equal(t, corev1.ServiceType("ClusterIP"), service.Spec.Type)
	require.Equal(t, "http", service.Spec.Ports[0].Name)
	require.EqualValues(t, 6333, service.Spec.Ports[0].Port)
	require.EqualValues(t, intstr.IntOrString{IntVal: 6333}, service.Spec.Ports[0].TargetPort)
	require.EqualValues(t, 0, service.Spec.Ports[0].NodePort)
	require.Equal(t, corev1.ProtocolTCP, service.Spec.Ports[0].Protocol)
	require.Equal(t, "grpc", service.Spec.Ports[1].Name)
	require.EqualValues(t, 6334, service.Spec.Ports[1].Port)
	require.EqualValues(t, intstr.IntOrString{IntVal: 6334}, service.Spec.Ports[1].TargetPort)
	require.EqualValues(t, 0, service.Spec.Ports[1].NodePort)
	require.Equal(t, corev1.ProtocolTCP, service.Spec.Ports[1].Protocol)
}
