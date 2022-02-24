package filereader

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAllLines(t *testing.T) {
	filepath := "./test.txt"
	os.Remove(filepath)
	wait := "test"
	file, _ := os.Create(filepath)
	file.WriteString(wait)
	out := ReadAllLines(filepath)
	assert.EqualValues(t, wait, out[0])
	os.Remove(filepath)

}
