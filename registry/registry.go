package registry

import (
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/chunyu224/kratos-bootstrap/registry/consul"
	"github.com/chunyu224/kratos-bootstrap/registry/etcd"
	"github.com/chunyu224/kratos-bootstrap/registry/eureka"
	"github.com/chunyu224/kratos-bootstrap/registry/kubernetes"
	"github.com/chunyu224/kratos-bootstrap/registry/nacos"
	"github.com/chunyu224/kratos-bootstrap/registry/servicecomb"
	"github.com/chunyu224/kratos-bootstrap/registry/zookeeper"

	conf "github.com/chunyu224/kratos-bootstrap/api/gen/go/conf/v1"
)

// NewRegistry 创建一个注册客户端
func NewRegistry(cfg *conf.Registry) registry.Registrar {
	if cfg == nil {
		return nil
	}

	switch Type(cfg.Type) {
	case Consul:
		return consul.NewRegistry(cfg)
	case Etcd:
		return etcd.NewRegistry(cfg)
	case ZooKeeper:
		return zookeeper.NewRegistry(cfg)
	case Nacos:
		return nacos.NewRegistry(cfg)
	case Kubernetes:
		return kubernetes.NewRegistry(cfg)
	case Eureka:
		return eureka.NewRegistry(cfg)
	case Polaris:
		return nil
	case Servicecomb:
		return servicecomb.NewRegistry(cfg)
	}

	return nil
}

// NewDiscovery 创建一个发现客户端
func NewDiscovery(cfg *conf.Registry) registry.Discovery {
	if cfg == nil {
		return nil
	}

	switch Type(cfg.Type) {
	case Consul:
		return consul.NewRegistry(cfg)
	case Etcd:
		return etcd.NewRegistry(cfg)
	case ZooKeeper:
		return zookeeper.NewRegistry(cfg)
	case Nacos:
		return nacos.NewRegistry(cfg)
	case Kubernetes:
		return kubernetes.NewRegistry(cfg)
	case Eureka:
		return eureka.NewRegistry(cfg)
	case Polaris:
		return nil
	case Servicecomb:
		return servicecomb.NewRegistry(cfg)
	}

	return nil
}
