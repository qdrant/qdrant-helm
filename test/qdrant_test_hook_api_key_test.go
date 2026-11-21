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
	corev1 "k8s.io/api/core/v1"
)

func TestTestHookApiKey(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	tests := []struct {
		name         string
		values       map[string]string
		secretKeyRef corev1.SecretKeySelector
	}{
		{
			name:         "generated api key",
			values:       map[string]string{"apiKey": "true"},
			secretKeyRef: corev1.SecretKeySelector{LocalObjectReference: corev1.LocalObjectReference{Name: "qdrant-apikey"}, Key: "api-key"},
		},
		{
			name:         "string api key",
			values:       map[string]string{"apiKey": `"test_api_key"`},
			secretKeyRef: corev1.SecretKeySelector{LocalObjectReference: corev1.LocalObjectReference{Name: "qdrant-apikey"}, Key: "api-key"},
		},
		{
			name:         "generated read-only api key",
			values:       map[string]string{"readOnlyApiKey": "true"},
			secretKeyRef: corev1.SecretKeySelector{LocalObjectReference: corev1.LocalObjectReference{Name: "qdrant-apikey"}, Key: "read-only-api-key"},
		},
		{
			name:         "api key from an existing secret",
			values:       map[string]string{"apiKey": `{"valueFrom":{"secretKeyRef":{"name":"external","key":"key"}}}`},
			secretKeyRef: corev1.SecretKeySelector{LocalObjectReference: corev1.LocalObjectReference{Name: "external"}, Key: "key"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			namespaceName := "qdrant-" + strings.ToLower(random.UniqueID())
			logger.Log(t, "Namespace: %s\n", namespaceName)

			options := &helm.Options{
				SetJSONValues:  tt.values,
				KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
			}

			output := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, releaseName, []string{"templates/tests/test-db-interaction.yaml"})
			var pod corev1.Pod
			var configMap corev1.ConfigMap
			for _, doc := range strings.Split(output, "\n---\n") {
				if strings.Contains(doc, "\nkind: Pod\n") {
					helm.UnmarshalK8SYaml(t, doc, &pod)
				} else if strings.Contains(doc, "\nkind: ConfigMap\n") {
					helm.UnmarshalK8SYaml(t, doc, &configMap)
				}
			}
			require.Len(t, pod.Spec.Containers, 1)

			env := pod.Spec.Containers[0].Env
			require.Len(t, env, 1)
			require.Equal(t, "QDRANT_API_KEY", env[0].Name)
			require.NotNil(t, env[0].ValueFrom)
			require.Equal(t, &tt.secretKeyRef, env[0].ValueFrom.SecretKeyRef)
			require.Contains(t, configMap.Data["entrypoint.sh"], `API_KEY_HEADER="Api-key: ${QDRANT_API_KEY}"`)
		})
	}
}
