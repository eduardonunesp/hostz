package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadHostsFile(t *testing.T) {
	hosts := NewHostsParser()
	result, err := hosts.ReadHostsFile("../../fixtures/hosts")
	assert.NotNil(t, result)
	assert.Equal(t, err, nil)
}

func TestParseHosts(t *testing.T) {
	hosts := NewHostsParser()
	hostBytes, err := hosts.ReadHostsFile("../../fixtures/hosts")
	assert.Nil(t, err)
	results := hosts.ParseHosts(hostBytes)
	assert.Equal(t, len(results), 3)
}
