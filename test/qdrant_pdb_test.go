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
	policyv1 "k8s.io/api/policy/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func renderPodDisruptionBudget(t *testing.T, values map[string]string) policyv1.PodDisruptionBudget {
	helmChartPath, err := filepath.Abs("../charts/qdrant")
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueID())
	values["podDisruptionBudget.enabled"] = "true"
	options := &helm.Options{
		SetValues:      values,
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplateContext(t, context.Background(), options, helmChartPath, "qdrant", []string{"templates/pdb.yaml"})

	var pdb policyv1.PodDisruptionBudget
	helm.UnmarshalK8SYaml(t, output, &pdb)
	return pdb
}

func TestPodDisruptionBudgetDefaultsToMaxUnavailable(t *testing.T) {
	t.Parallel()

	pdb := renderPodDisruptionBudget(t, map[string]string{})

	require.Equal(t, intstr.FromInt32(1), *pdb.Spec.MaxUnavailable)
	require.Nil(t, pdb.Spec.MinAvailable)
}

func TestPodDisruptionBudgetKeepsMaxUnavailableZero(t *testing.T) {
	t.Parallel()

	pdb := renderPodDisruptionBudget(t, map[string]string{
		"podDisruptionBudget.maxUnavailable": "0",
	})

	require.NotNil(t, pdb.Spec.MaxUnavailable)
	require.Equal(t, intstr.FromInt32(0), *pdb.Spec.MaxUnavailable)
	require.Nil(t, pdb.Spec.MinAvailable)
}

func TestPodDisruptionBudgetMinAvailableReplacesDefaultMaxUnavailable(t *testing.T) {
	t.Parallel()

	pdb := renderPodDisruptionBudget(t, map[string]string{
		"podDisruptionBudget.minAvailable": "2",
	})

	require.NotNil(t, pdb.Spec.MinAvailable)
	require.Equal(t, intstr.FromInt32(2), *pdb.Spec.MinAvailable)
	require.Nil(t, pdb.Spec.MaxUnavailable)
}

func TestPodDisruptionBudgetKeepsPercentages(t *testing.T) {
	t.Parallel()

	pdb := renderPodDisruptionBudget(t, map[string]string{
		"podDisruptionBudget.maxUnavailable": "50%",
	})

	require.Equal(t, intstr.FromString("50%"), *pdb.Spec.MaxUnavailable)
	require.Nil(t, pdb.Spec.MinAvailable)
}
