package kubernetes

import (
	"testing"

	conf "github.com/chunyu224/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/stretchr/testify/assert"
)

func TestNewKubernetesRegistry(t *testing.T) {
	var cfg conf.Registry
	reg := NewRegistry(&cfg)
	assert.Nil(t, reg)
}
