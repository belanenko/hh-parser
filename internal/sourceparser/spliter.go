package sourceparser

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/hh-parser/internal/models"
)

func UnmarshalFromHTML(str string) (models.Vacancy, error) { // TODO: Прикрутить дженерики
	arr := strings.Split(str, `<script type="application/ld+json">`)
	if len(arr) == 1 {
		return models.Vacancy{}, errors.New("Bad html file")
	}
	this_json := strings.Split(arr[1], `</script>`)[0]

	this_json = strings.ReplaceAll(this_json, "\n", "")
	var out models.Vacancy
	badd := []byte(this_json)
	err := json.Unmarshal(badd, &out)
	if err != nil {
		log.Fatalln(err)
	}
	return out, nil
}
