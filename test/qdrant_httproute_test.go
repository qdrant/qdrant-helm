package test

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
)

func TestHTTPRouteDisabledByDefault(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	require.NoError(t, err)

	options := &helm.Options{
		KubectlOptions: k8s.NewKubectlOptions("", "", "qdrant-"+strings.ToLower(random.UniqueID())),
	}

	_, err = helm.RenderTemplateContextE(t, context.Background(), options, helmChartPath, "qdrant", []string{"templates/httproute.yaml"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "could not find template")
}

func TestHTTPRoute(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	require.NoError(t, err)

	options := &helm.Options{
		SetJSONValues: map[string]string{
			"httpRoute.enabled":     `true`,
			"httpRoute.parentRefs":  `[{"name":"gw","namespace":"gateway-system","sectionName":"https"}]`,
			"httpRoute.hostnames":   `["test.qdrant.local"]`,
			"httpRoute.annotations": `{"a":"b"}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", "qdrant-"+strings.ToLower(random.UniqueID())),
	}

	output := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, "qdrant", []string{"templates/httproute.yaml"})

	var route map[string]any
	helm.UnmarshalK8SYaml(t, output, &route)

	require.Equal(t, "gateway.networking.k8s.io/v1", route["apiVersion"])
	require.Equal(t, "HTTPRoute", route["kind"])
	require.Equal(t, "b", route["metadata"].(map[string]any)["annotations"].(map[string]any)["a"])

	spec := route["spec"].(map[string]any)
	parent := spec["parentRefs"].([]any)[0].(map[string]any)
	require.Equal(t, "gw", parent["name"])
	require.Equal(t, "gateway-system", parent["namespace"])
	require.Equal(t, "https", parent["sectionName"])
	require.Equal(t, []any{"test.qdrant.local"}, spec["hostnames"])

	rule := spec["rules"].([]any)[0].(map[string]any)
	path := rule["matches"].([]any)[0].(map[string]any)["path"].(map[string]any)
	require.Equal(t, "PathPrefix", path["type"])
	require.Equal(t, "/", path["value"])
	backend := rule["backendRefs"].([]any)[0].(map[string]any)
	require.Equal(t, "qdrant", backend["name"])
	require.EqualValues(t, 6333, backend["port"])
}

func TestHTTPRouteRequiresParentRefs(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	require.NoError(t, err)

	options := &helm.Options{
		SetJSONValues:  map[string]string{"httpRoute.enabled": `true`},
		KubectlOptions: k8s.NewKubectlOptions("", "", "qdrant-"+strings.ToLower(random.UniqueID())),
	}

	_, err = helm.RenderTemplateContextE(t, context.Background(), options, helmChartPath, "qdrant", []string{"templates/httproute.yaml"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "httpRoute.parentRefs")
}

func TestHTTPRouteCustomBackendAndLabels(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	require.NoError(t, err)

	options := &helm.Options{
		SetJSONValues: map[string]string{
			"httpRoute.enabled":          `true`,
			"httpRoute.parentRefs":       `[{"name":"gw"}]`,
			"httpRoute.pathPrefix":       `"/qdrant"`,
			"httpRoute.servicePort":      `6334`,
			"httpRoute.additionalLabels": `{"route":"label"}`,
			"additionalAnnotations":      `{"global":"anno"}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", "qdrant-"+strings.ToLower(random.UniqueID())),
	}

	output := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, "qdrant", []string{"templates/httproute.yaml"})

	var route map[string]any
	helm.UnmarshalK8SYaml(t, output, &route)

	metadata := route["metadata"].(map[string]any)
	require.Equal(t, "label", metadata["labels"].(map[string]any)["route"])
	require.Equal(t, "anno", metadata["annotations"].(map[string]any)["global"])

	spec := route["spec"].(map[string]any)
	require.NotContains(t, spec, "hostnames")

	rule := spec["rules"].([]any)[0].(map[string]any)
	path := rule["matches"].([]any)[0].(map[string]any)["path"].(map[string]any)
	require.Equal(t, "/qdrant", path["value"])
	require.EqualValues(t, 6334, rule["backendRefs"].([]any)[0].(map[string]any)["port"])
}
