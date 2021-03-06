package vacancystorage

import (
	"strings"
	"sync"

	"github.com/hh-parser/internal/models/vacancy"
)

func init() {
	Storage = &storage{}
}

type storage struct {
	m         sync.Mutex
	vacancies []vacancy.Vacancy
}

var Storage *storage

func (s *storage) Pop(vacancies ...vacancy.Vacancy) {
	s.m.Lock()
	defer s.m.Unlock()
	for _, v := range vacancies {
		if !s.vacancyValidate(&v) {
			continue
		}
		s.vacancies = append(s.vacancies, v)
	}
}

func (s *storage) vacancyValidate(v *vacancy.Vacancy) bool {
	// TODO: Impliment validation logic
	v.Description = strings.Replace(v.Description, "'", "''", -1)
	return true
}

func (s *storage) PushCount(count uint) []vacancy.Vacancy {
	s.m.Lock()
	defer s.m.Unlock()
	vacanciesLen := len(s.vacancies)
	if count > uint(vacanciesLen) {
		count = uint(vacanciesLen)
	}
	this_vacancies := s.vacancies[:count]
	s.vacancies = s.vacancies[count:]
	return this_vacancies
}

func (s *storage) Len() int {
	s.m.Lock()
	defer s.m.Unlock()
	return len(s.vacancies)
}
