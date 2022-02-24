package filereader

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAllLines(t *testing.T) {
	filepath := "./test.txt"
	os.Remove(filepath)
	wait := []string{
		"test",
		"test1",
		"test2",
	}
	file, _ := os.Create(filepath)
	file.WriteString(strings.Join(wait, "\n"))
	out := ReadAllLines(filepath)
	assert.EqualValues(t, wait[0], out[0])
	assert.EqualValues(t, wait[1], out[1])
	assert.EqualValues(t, wait[2], out[2])
	os.Remove(filepath)

}
