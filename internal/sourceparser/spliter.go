package sourceparser

import (
	"errors"
	"strings"
)

func GetJsonFromHtmlString(str string) (string, error) {
	startIndex := strings.Index(str, `<script type="application/ld+json">`)
	if startIndex == -1 {
		return "", errors.New("json not finded in page")
	}

	startIndex += len(`<script type="application/ld+json">`)
	stopIndex := strings.Index(str[startIndex:], `</script>`)
	return str[startIndex : startIndex+stopIndex], nil

}
