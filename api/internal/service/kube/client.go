package kube

import (
	"encoding/json"
	"fmt"

	"github.com/gotomicro/ego/core/elog"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	clientcmdlatest "k8s.io/client-go/tools/clientcmd/api/latest"
	clientcmdapiv1 "k8s.io/client-go/tools/clientcmd/api/v1"
)

func buildClient(apiServerAddr string, kubeconfig string) (*kubernetes.Clientset, *rest.Config, error) {
	configV1 := clientcmdapiv1.Config{}
	err := json.Unmarshal([]byte(kubeconfig), &configV1)
	if err != nil {
		elog.Error("json unmarshal kubeconfig error.", elog.String("kubeconfig", kubeconfig), elog.FieldErr(err))
		return nil, nil, err
	}
	var configObject runtime.Object
	configObject, err = clientcmdlatest.Scheme.ConvertToVersion(&configV1, clientcmdapi.SchemeGroupVersion)
	if err != nil {
		elog.Error("build client config error. ", zap.Error(err))
		return nil, nil, err
	}
	configInternal := configObject.(*clientcmdapi.Config)
	var clientConfig *rest.Config
	clientConfig, err = clientcmd.NewDefaultClientConfig(*configInternal, &clientcmd.ConfigOverrides{
		ClusterDefaults: clientcmdapi.Cluster{Server: apiServerAddr}, // InsecureSkipTLSVerify: true
	}).ClientConfig()
	if err != nil {
		elog.Error("build client config error. ", zap.Error(err))
		return nil, nil, err
	}
	clientConfig.QPS = defaultQPS
	clientConfig.Burst = defaultBurst

	clientSet, err := kubernetes.NewForConfig(clientConfig)

	if err != nil {
		elog.Error(fmt.Sprintf("apiServerAddr(%s) kubernetes.NewForConfig(%v) error.", apiServerAddr, clientConfig), zap.Error(err))
		return nil, nil, err
	}

	return clientSet, clientConfig, nil
}
