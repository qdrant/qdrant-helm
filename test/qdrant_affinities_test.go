package test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAffinities(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"affinity": `
			{
				"podAntiAffinity":
				{
					"requiredDuringSchedulingIgnoredDuringExecution":
					[
						{
							"labelSelector":
							{
								"matchExpressions":
								[
									{"key": "app.kubernetes.io/name", "operator":"In", "values":["{{ include \"qdrant.name\" . }}"]},
									{"key": "app.kubernetes.io/instance", "operator":"In", "values":["{{ .Release.Name }}"]}
								]
							},
							"topologyKey": "kubernetes.io/hostname"
						}
					]
				}
			}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, output, &statefulSet)

	matchExpressions := statefulSet.Spec.Template.Spec.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution[0].LabelSelector.MatchExpressions

	require.Equal(t, 2, len(matchExpressions))

	// first match expression
	require.Equal(t, v1.LabelSelectorOperator("In"), matchExpressions[0].Operator)
	require.Equal(t, "app.kubernetes.io/name", matchExpressions[0].Key)
	require.Equal(t, "qdrant", matchExpressions[0].Values[0])

	// second match expression
	require.Equal(t, "app.kubernetes.io/instance", matchExpressions[1].Key)
	require.Equal(t, "qdrant", matchExpressions[1].Values[0])
}
