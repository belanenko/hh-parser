package jsonparser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJsonFromHTML(t *testing.T) {
	wait := `{"@context": "https://schema.org/"}`
	html := fmt.Sprintf(`</div></div></div><script type="application/ld+json">%s</script></div>`, wait)
	expect := GetJsonFromHTML(html)
	assert.EqualValues(t, wait, expect)
}
