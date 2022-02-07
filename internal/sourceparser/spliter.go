package sourceparser

import (
	"encoding/json"
	"strings"

	"github.com/hh-parser/internal/models"
)

func UnmarshalString(str string) models.Vacancy { // TODO: Прикрутить дженерики
	arr := strings.Split(str, `<script type="application/ld+json">`)
	this_json := strings.Split(arr[1], `</script>`)[0]

	var out models.Vacancy
	json.Unmarshal([]byte(this_json), &out)
	return out
}
