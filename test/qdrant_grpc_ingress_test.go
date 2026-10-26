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
	networkingv1 "k8s.io/api/networking/v1"
)

func TestGrpcIngressDisabledByDefault(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	require.NoError(t, err)

	options := &helm.Options{
		KubectlOptions: k8s.NewKubectlOptions("", "", "qdrant-"+strings.ToLower(random.UniqueID())),
	}

	_, err = helm.RenderTemplateE(t, options, helmChartPath, "qdrant", []string{"templates/grpc-ingress.yaml"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "could not find template")
}

func TestGrpcIngress(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueID())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJSONValues: map[string]string{
			"grpcIngress.enabled":          `true`,
			"grpcIngress.ingressClassName": `"nginx"`,
			"grpcIngress.annotations":      `{"nginx.ingress.kubernetes.io/backend-protocol":"GRPC"}`,
			"grpcIngress.tls":              `[{"hosts":["grpc.qdrant.local"],"secretName":"grpc-secret"}]`,
			"grpcIngress.hosts":            `[{"host":"grpc.qdrant.local","paths":[{"path":"/","pathType":"Prefix","servicePort": 6334}]}]`,
			"ingress.enabled":              `true`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, releaseName, []string{"templates/grpc-ingress.yaml"})

	var ingress networkingv1.Ingress
	helm.UnmarshalK8SYaml(t, output, &ingress)

	require.Equal(t, "qdrant-grpc", ingress.Name)
	require.Equal(t, "nginx", *ingress.Spec.IngressClassName)
	require.Equal(t, "GRPC", ingress.Annotations["nginx.ingress.kubernetes.io/backend-protocol"])
	require.Equal(t, "grpc.qdrant.local", ingress.Spec.Rules[0].Host)
	require.Equal(t, "qdrant", ingress.Spec.Rules[0].HTTP.Paths[0].Backend.Service.Name)
	require.Equal(t, int32(6334), ingress.Spec.Rules[0].HTTP.Paths[0].Backend.Service.Port.Number)
	require.Equal(t, "grpc.qdrant.local", ingress.Spec.TLS[0].Hosts[0])
	require.Equal(t, "grpc-secret", ingress.Spec.TLS[0].SecretName)
}
