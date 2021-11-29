package kitex_k8s_native_discovery

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/discovery"
	"net"
)

type HostPort struct {
	net.Addr
	network string
	address string
}

func (hp HostPort) Network() string {
	return hp.network
}

func (hp HostPort) String() string {
	return hp.address
}

type ServiceEndpoint struct {
	discovery.Instance
	network     string
	serviceName string
}

func NewTcpEndpoint(svcName string) ServiceEndpoint {
	return ServiceEndpoint{
		network:     "tcp",
		serviceName: svcName,
	}
}

func (e ServiceEndpoint) Address() net.Addr {
	return HostPort{
		network: e.network,
		// 由于address的格式要求为 host:port / ip:port
		// 约定80端口，不需要修改
		address: fmt.Sprintf("%s:80", e.serviceName),
	}
}

func (e ServiceEndpoint) Weight() int {
	return 10
}

func (e ServiceEndpoint) Tag(key string) (value string, exist bool) {
	// 没有任何tag
	return "", false
}
