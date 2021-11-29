package kitex_k8s_native_discovery

import (
	"github.com/cloudwego/kitex/client"
	"time"
)

type ProDeployEnv struct {
	client.Suite
	namespace string
}

func NewDiscoveryWithNamespace(namespace string) *ProDeployEnv {
	return &ProDeployEnv{
		namespace: namespace,
	}
}

func NewDiscovery() *ProDeployEnv {
	return &ProDeployEnv{
	}
}

func (e *ProDeployEnv) Options() []client.Option {
	resolver := NewResolverWithNamespace(e.namespace)
	return []client.Option{
		client.WithResolver(resolver),
		client.WithRPCTimeout(10 * time.Second),
		client.WithConnectTimeout(3 * time.Second),
	}
}
