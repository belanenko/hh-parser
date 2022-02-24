package interfaces

import "github.com/hh-parser/internal/models/vacancy"

type VacancyStorage interface {
	Pop(vacancies ...vacancy.Vacancy)
	PushCount(count uint) []vacancy.Vacancy
	Len() int
}
