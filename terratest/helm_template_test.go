package test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	k8score "k8s.io/api/core/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
)

func TestConfigMap(t *testing.T) {

	helmChartPath, err := filepath.Abs("../helm-charts/example")
	releaseName := "example-release"
	require.NoError(t, err)
	namespaceName := "test-namespace" + strings.ToLower(random.UniqueId())
	logger.Logf(t, "Namespace: %s\n", namespaceName)

	options := &helm.Options{
		SetValues: map[string]string{
			"image.registry":     "commonregistry",
			"image.tag":          "latest",
			"myKeyValueMap.key0": "val0",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/configmap.yaml"})
	// render
	var configMap k8score.ConfigMap
	helm.UnmarshalK8SYaml(t, output, &configMap)
	// validate
	require.Equal(t, releaseName+"-configmap", configMap.GetName())
	configMapData := configMap.Data
	require.Equal(t, 1, len(configMapData))
	require.Equal(t, "key0=val0", configMapData["application.properties"])

	outputDeployment := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/alpine-deployment.yaml"})
	var deployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, outputDeployment, &deployment)

}
