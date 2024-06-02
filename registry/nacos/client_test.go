package nacos

import (
	"testing"

	conf "github.com/chunyu224/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/stretchr/testify/assert"
)

func TestNewNacosRegistry(t *testing.T) {
	var cfg conf.Registry
	cfg.Nacos.Address = "127.0.0.1"
	cfg.Nacos.Port = 8848

	reg := NewRegistry(&cfg)
	assert.Nil(t, reg)
}
