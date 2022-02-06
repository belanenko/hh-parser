package storage

import (
	"strings"

	"github.com/hh-parser/structs"
)

type infoStorage []structs.Jobs

func CreateInfoStorage() infoStorage {
	return make(infoStorage, 0)
}

func (is *infoStorage) AddJobs(jobs ...structs.Jobs) {
	*is = append(*is, jobs...)
}

func (is *infoStorage) GetCount() int {
	return len(*is)
}

func (is *infoStorage) FindAllByCompany(companyName string) []structs.Jobs {
	var out []structs.Jobs
	for _, v := range *is {
		if strings.ToLower(v.Company.Name) == strings.ToLower(companyName) {
			out = append(out, v)
		}
	}
	return out
}
