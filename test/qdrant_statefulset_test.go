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
)

func TestContainerSecurityContextUserAndGroupDefault(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues:  map[string]string{},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	statefulsetOutput := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, statefulsetOutput, &statefulSet)
	require.Equal(t, "1000:3000", statefulSet.Spec.Template.Spec.InitContainers[0].Command[2])
	require.Equal(t, int64(1000), *statefulSet.Spec.Template.Spec.Containers[0].SecurityContext.RunAsUser)
	require.Equal(t, int64(3000), *statefulSet.Spec.Template.Spec.SecurityContext.FSGroup)
	require.Equal(t, int64(2000), *statefulSet.Spec.Template.Spec.Containers[0].SecurityContext.RunAsGroup)
}

func TestContainerSecurityContextUserAndGroupLargeValues(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"containerSecurityContext": `{
				"runAsUser": 1000640000,
				"runAsGroup": 1000840000
			}`,
			"podSecurityContext": `{
				"fsGroup": 1000740000
			}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	statefulsetOutput := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, statefulsetOutput, &statefulSet)
	require.Equal(t, "1000640000:1000740000", statefulSet.Spec.Template.Spec.InitContainers[0].Command[2])
	require.Equal(t, int64(1000640000), *statefulSet.Spec.Template.Spec.Containers[0].SecurityContext.RunAsUser)
	require.Equal(t, int64(1000740000), *statefulSet.Spec.Template.Spec.SecurityContext.FSGroup)
	require.Equal(t, int64(1000840000), *statefulSet.Spec.Template.Spec.Containers[0].SecurityContext.RunAsGroup)
}

func TestCanAttachVolumeAttributesClassToPersistence(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../charts/qdrant")
	releaseName := "qdrant"
	require.NoError(t, err)

	namespaceName := "qdrant-" + strings.ToLower(random.UniqueId())
	logger.Log(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetJsonValues: map[string]string{
			"persistence": `{
				"volumeAttributesClassName": "my-class"
			}`,
			"snapshotPersistence": `{
				"enabled": true,
				"volumeAttributesClassName": "my-class"
			}`,
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	statefulsetOutput := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, statefulsetOutput, &statefulSet)
	require.Equal(t, "my-class", *statefulSet.Spec.VolumeClaimTemplates[0].Spec.VolumeAttributesClassName)
	require.Equal(t, "my-class", *statefulSet.Spec.VolumeClaimTemplates[1].Spec.VolumeAttributesClassName)
}
