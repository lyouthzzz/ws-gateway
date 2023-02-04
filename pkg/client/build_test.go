package client

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig_Build(t *testing.T) {
	c := DefaultConfig().WithAddr("localhost:8080")
	require.Equal(t, "localhost:8080", c.Addr)

	_, err := DefaultConfig().WithAddr("localhost:8080").BuildGRPCClient(context.Background())
	require.Nil(t, err)
}
