package kitex_k8s_native_discovery

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"strings"
)

// Resolver 使用k8s集群内service做服务发现
type Resolver struct {
	discovery.Resolver
	namespace string
	domainSfx string // {svc}.{ns}.{domainSfx}
}

func NewResolver() *Resolver {
	return &Resolver{
		domainSfx: "svc.cluster.local",
	}
}

func NewResolverWithNamespace(ns string) *Resolver {
	return &Resolver{
		namespace: ns,
		domainSfx: "svc.cluster.local",
	}
}

func (r *Resolver) SetDomainSuffix(sfx string) {
	r.domainSfx = sfx
}

// Target get target key, we need to hashable key for cache,
// but EndpointInfo is a interface, so just exchange it to string(or? no hashable type with golang)
func (r *Resolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) string {
	svcName := target.ServiceName()
	targetSvcPfx := strings.ReplaceAll(svcName, ".", "-")
	targetSvcPfx = strings.ReplaceAll(targetSvcPfx, "_", "-")
	svcHost := fmt.Sprintf("%s", targetSvcPfx) // a-b-c
	if r.namespace != "" {
		svcHost = fmt.Sprintf("%s.%s.%s", svcHost, r.namespace, r.domainSfx)
	}
	return svcHost
}

// Resolve find instances from target key
func (r *Resolver) Resolve(ctx context.Context, key string) (discovery.Result, error) {
	result := discovery.Result{
		Cacheable: false,
		Instances: []discovery.Instance{
			NewTcpEndpoint(key),
		},
	}
	return result, nil
}

// Diff support watcher by given diff function
func (r *Resolver) Diff(key string, prev, next discovery.Result) (discovery.Change, bool) {
	// return false to indicate unsupported of diff
	return discovery.Change{}, false
}

// Name unique key for cache and reuse
func (r *Resolver) Name() string {
	return "kitex-k8s-native-discovery-resolver"
}
