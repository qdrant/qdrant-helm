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
	appsv1 "k8s.io/api/apps/v1"
)

func TestTopologySpreadConstraints(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueID())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJSONValues: map[string]string{
			"topologySpreadConstraints": `
			[
                {
                    "maxSkew": 1,
					"topologyKey": "topology.kubernetes.io/zone",
				    "labelSelector": {
                        "matchLabels": {
                            "app.kubernetes.io/name": "{{ include \"qdrant.name\" . }}"
                        }
                    },
                    "whenUnsatisfiable": "DoNotSchedule"
                }
            ]`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderRemoteTemplateContext(t, context.Background(), options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, output, &statefulSet)

	require.Equal(t, 1, len(statefulSet.Spec.Template.Spec.TopologySpreadConstraints))
	require.Equal(t, int32(1), statefulSet.Spec.Template.Spec.TopologySpreadConstraints[0].MaxSkew)
}
