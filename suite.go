package kitex_k8s_native_discovery

import (
	"github.com/cloudwego/kitex/client"
)

type SvcDiscovery struct {
	client.Suite
	namespace string
}

func NewDiscoveryWithNamespace(namespace string) *SvcDiscovery {
	return &SvcDiscovery{
		namespace: namespace,
	}
}

func NewDiscovery() *SvcDiscovery {
	return &SvcDiscovery{}
}

func (e *SvcDiscovery) Options() []client.Option {
	resolver := NewResolverWithNamespace(e.namespace)
	return []client.Option{
		client.WithResolver(resolver),
	}
}
