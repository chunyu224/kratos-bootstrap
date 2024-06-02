package consul

import (
	"testing"

	conf "github.com/chunyu224/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/stretchr/testify/assert"
)

func TestNewConsulRegistry(t *testing.T) {
	var cfg conf.Registry
	cfg.Consul.Scheme = "http"
	cfg.Consul.Address = "localhost:8500"
	cfg.Consul.HealthCheck = false

	reg := NewRegistry(&cfg)
	assert.Nil(t, reg)
}
