package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadHostsFile(t *testing.T) {
	hosts := NewHostsParser()
	result, err := hosts.ReadHostsFile("../../fixtures/hosts")
	require.NotNil(t, result)
	require.Equal(t, err, nil)
}

func TestParseHosts(t *testing.T) {
	hosts := NewHostsParser()
	hostBytes, err := hosts.ReadHostsFile("../../fixtures/hosts")
	require.Nil(t, err)
	results := hosts.ParseHosts(hostBytes)
	require.Equal(t, len(results), 3)
}
