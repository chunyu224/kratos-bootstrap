package servicecomb

import (
	"testing"

	conf "github.com/chunyu224/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/stretchr/testify/assert"
)

func TestNewServicecombRegistry(t *testing.T) {
	var cfg conf.Registry
	cfg.Servicecomb.Endpoints = []string{"127.0.0.1:30100"}

	reg := NewRegistry(&cfg)
	assert.Nil(t, reg)
}
