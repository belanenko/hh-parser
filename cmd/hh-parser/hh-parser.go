package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/hh-parser/internal/fioutil"
	"github.com/hh-parser/internal/sourceparser"
)

const PATH_TO_DIR = "./../../../../../../resp"

func main() {
	startTime := time.Now()

	files, err := ioutil.ReadDir(PATH_TO_DIR)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		strjson := fioutil.ReadFileToString(fmt.Sprintf("%s/%s", PATH_TO_DIR, file.Name()))
		vac, err := sourceparser.UnmarshalFromHTML(strjson)
		if err != nil {
			continue
		}
		fmt.Println("--------------------------------------")
		fmt.Printf("Название вакансии: %s\nID: %d\n", vac.Title, vac.Identifier.Value)
		fmt.Printf("Дата публикации: %s\nНаниматель: %s\n", vac.ValidThrough, vac.HiringOrganization.Name)
	}

	stopTime := time.Now()

	fmt.Printf("\n\nВыполнено за %f сек.", stopTime.Sub(startTime).Seconds())

}
