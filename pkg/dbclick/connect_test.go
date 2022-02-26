package dbclick

import (
	"testing"

	"github.com/hh-parser/pkg/dbclick/clickconfig"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	config := clickconfig.ReadConfig("/home/tim/.config/hh-parser/config.json")
	answer := Ping(config)
	assert.EqualValues(t, "row_data\n", answer)
}
