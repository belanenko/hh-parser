package convertor

import (
	"encoding/json"

	"github.com/hh-parser/internal/models/vacancy"
)

func JsonToModel(jsonStr string, vacancy *vacancy.Vacancy) error {
	return json.Unmarshal([]byte(jsonStr), vacancy)
}
