package dbclick

import (
	"testing"

	"github.com/hh-parser/pkg/dbclick/clickconfig"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	config := clickconfig.ReadConfig("/home/tim/.config/hh-parser/config.json")
	answer := Send(config, "SHOW TABLES FROM hh1;")
	assert.EqualValues(t, "row_data\n", answer)
}
