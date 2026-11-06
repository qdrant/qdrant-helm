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

func TestIngressWithIngressClassName(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueID())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJSONValues: map[string]string{
			"ingress.enabled":          `true`,
			"ingress.ingressClassName": `"nginx"`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, releaseName, []string{"templates/ingress.yaml"})

	var ingress networkingv1.Ingress
	helm.UnmarshalK8SYaml(t, output, &ingress)

	require.Equal(t, "nginx", *ingress.Spec.IngressClassName)
}

func TestIngressWithTls(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueID())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJSONValues: map[string]string{
			"ingress.enabled": `true`,
			"ingress.hosts":   `[{"host":"test.qdrant.local","paths":[{"path":"/","pathType":"Prefix","servicePort": 6333}]}]`,
			"ingress.tls":     `[{"hosts":["test.qdrant.local"],"secretName":"test-secret"}]`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, releaseName, []string{"templates/ingress.yaml"})

	var ingress networkingv1.Ingress
	helm.UnmarshalK8SYaml(t, output, &ingress)

	require.Equal(t, "test.qdrant.local", ingress.Spec.Rules[0].Host)
	require.Equal(t, "/", ingress.Spec.Rules[0].HTTP.Paths[0].Path)
	require.Equal(t, networkingv1.PathType("Prefix"), *ingress.Spec.Rules[0].HTTP.Paths[0].PathType)
	require.Equal(t, int32(6333), ingress.Spec.Rules[0].HTTP.Paths[0].Backend.Service.Port.Number)
	require.Equal(t, "test.qdrant.local", ingress.Spec.TLS[0].Hosts[0])
	require.Equal(t, "test-secret", ingress.Spec.TLS[0].SecretName)
}

func TestIngressBackendServiceNames(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	require.NoError(t, err)
	options := &helm.Options{
		SetJSONValues: map[string]string{
			"ingress.enabled":  `true`,
			"fullnameOverride": `"release-backend"`,
			"ingress.hosts": `[
				{"host":"first.example.com","paths":[
					{"path":"/custom","serviceName":"custom-qdrant-service","servicePort":6333},
					{"path":"/default","servicePort":6334},
					{"path":"/empty","serviceName":"","servicePort":6333}
				]},
				{"host":"second.example.com","paths":[
					{"path":"/other","serviceName":"other-service","servicePort":8080}
				]}
			]`,
		},
	}
	output := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, "qdrant", []string{"templates/ingress.yaml"})
	var ingress networkingv1.Ingress
	helm.UnmarshalK8SYaml(t, output, &ingress)

	require.Len(t, ingress.Spec.Rules, 2)
	paths := ingress.Spec.Rules[0].HTTP.Paths
	require.Len(t, paths, 3)
	require.Equal(t, "custom-qdrant-service", paths[0].Backend.Service.Name)
	require.Equal(t, int32(6333), paths[0].Backend.Service.Port.Number)
	require.Equal(t, "release-backend", paths[1].Backend.Service.Name)
	require.Equal(t, int32(6334), paths[1].Backend.Service.Port.Number)
	require.Equal(t, "release-backend", paths[2].Backend.Service.Name)
	otherPaths := ingress.Spec.Rules[1].HTTP.Paths
	require.Len(t, otherPaths, 1)
	require.Equal(t, "other-service", otherPaths[0].Backend.Service.Name)
	require.Equal(t, int32(8080), otherPaths[0].Backend.Service.Port.Number)
}
