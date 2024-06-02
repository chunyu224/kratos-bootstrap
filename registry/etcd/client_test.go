package etcd

import (
	"testing"

	conf "github.com/chunyu224/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/stretchr/testify/assert"
)

func TestNewEtcdRegistry(t *testing.T) {
	var cfg conf.Registry
	cfg.Etcd.Endpoints = []string{"127.0.0.1:2379"}

	reg := NewRegistry(&cfg)
	assert.Nil(t, reg)
}
