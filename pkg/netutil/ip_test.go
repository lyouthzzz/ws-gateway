package netutil

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	ip, err := LocalIP()
	require.NoError(t, err)

	fmt.Println(ip.String())
}
