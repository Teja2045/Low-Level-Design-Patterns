package functionalOptions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerOptions(t *testing.T) {
	maxConns := 10
	server := NewServer(WithMaxConns(maxConns))

	assert.Equal(t, server.Opts.maxConns, maxConns)

	server2 := NewServer(WithCustomID("custom"))
	assert.Equal(t, server2.Opts.id, "custom")
	assert.Equal(t, server.Opts.id, "default")
}
