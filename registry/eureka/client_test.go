package eureka

import (
	"testing"

	conf "github.com/chunyu224/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/stretchr/testify/assert"
)

func TestNewEurekaRegistry(t *testing.T) {
	var cfg conf.Registry
	cfg.Eureka.Endpoints = []string{"https://127.0.0.1:18761"}

	reg := NewRegistry(&cfg)
	assert.Nil(t, reg)
}
