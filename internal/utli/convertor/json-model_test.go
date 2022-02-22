package convertor

import (
	"fmt"
	"testing"

	"github.com/hh-parser/internal/models/vacancy"
	"github.com/stretchr/testify/assert"
)

func TestJsonToModel(t *testing.T) {
	description := "description"
	wait_vacancy := vacancy.Vacancy{
		Description: description,
	}
	expect_vacancy := vacancy.Vacancy{}
	this_json := fmt.Sprintf(`{"description": "%s"}`, description)
	JsonToModel(this_json, &expect_vacancy)

	assert.EqualValues(t, wait_vacancy.Description, expect_vacancy.Description)
}
