package vacancystorage

import (
	"testing"

	"github.com/hh-parser/internal/models/vacancy"
	"github.com/stretchr/testify/assert"
)

/*
*1. Добавить новую вакансию
	*1.1. Добавить валидатор добавляемой вакансии
*2. Получить вакансию и удалить из хранилища
*3. Количество вакансий в хранилище
*4. Получить N вакансий и удалить из хранилища
*/

func TestPopVacancy(t *testing.T) {
	vacancies := []vacancy.Vacancy{
		{},
		{},
		{},
		{},
	}

	this_storage := Storage
	this_storage.Pop(vacancies...)

	assert.Len(t, this_storage.vacancies, len(vacancies))
	Storage = &storage{}
}

func TestPushCount(t *testing.T) {
	tests := []struct {
		vacancies []vacancy.Vacancy
		want      int
		expected  int
	}{
		{
			vacancies: []vacancy.Vacancy{{}, {}, {}, {}},
			want:      2,
			expected:  2,
		},
		{
			vacancies: []vacancy.Vacancy{{}, {}, {}, {}},
			want:      5,
			expected:  4,
		},
		{
			vacancies: []vacancy.Vacancy{},
			want:      5,
			expected:  0,
		},
	}

	for _, test := range tests {
		Storage = &storage{}

		this_storage := Storage
		this_storage.Pop(test.vacancies...)
		pushed := this_storage.PushCount(uint(test.want))
		assert.Len(t, pushed, test.expected)
	}

	Storage = &storage{}
}
