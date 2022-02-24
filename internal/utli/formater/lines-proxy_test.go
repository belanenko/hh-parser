package formater

import (
	"testing"

	"github.com/hh-parser/internal/models/proxy"
	"github.com/stretchr/testify/assert"
)

func TestStringToProxy(t *testing.T) {
	lines := []string{
		"194.32.229.244:3001:ogJIFk:pdpGGSjLDa",
		"212.115.49.114:3001:ogJIFk:pdpGGSjLDa",
		"212.115.49.146:3001:ogJIFk:pdpGGSjLDa",
	}

	assert.Len(t, StringToProxy(proxy.SOCKS5, lines...), len(lines))
}
