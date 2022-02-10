package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/hh-parser/internal/db"
	"github.com/hh-parser/internal/fioutil"
	"github.com/hh-parser/internal/models"
	"github.com/hh-parser/internal/sourceparser"
)

const PATH_TO_DIR = "./../../../../../../resp/"

func main() {
	startTime := time.Now()

	files, err := ioutil.ReadDir(PATH_TO_DIR)
	if err != nil {
		panic(err)
	}

	var sb strings.Builder
	sb.WriteString("INSERT INTO hh1.test_table (Id, Description) VALUES")

	var countFiles int32

	for i := 0; i < len(files); i++ {
		this_file := PATH_TO_DIR + files[i].Name()
		if !fioutil.FileExists(this_file) {
			continue
		}

		content := fioutil.ReadFileToString(this_file)
		j, err := sourceparser.GetJsonFromHtmlString(content)
		if err != nil {
			continue
		}
		var out models.Vacancy
		json.Unmarshal([]byte(j), &out)

		countFiles++

		if i+1 == len(files) {
			sb.WriteString(fmt.Sprintf(" (%d, '%s');", out.Identifier.Value, out.Description))
			break
		}
		sb.WriteString(fmt.Sprintf(" (%d, '%s'),", out.Identifier.Value, out.Description))
	}

	fmt.Println("Кол-во файлов:", countFiles)
	var config db.Config
	db.Configuration("./conf.json", &config)

	client := db.CreateClient(config.CertPath)

	fmt.Println("Выполняем -> SELECT * FROM hh1.test_table:")
	answer := db.SendRequest(db.GET, client, config, "SELECT * FROM hh1.test_table")
	fmt.Println(answer)

	fmt.Println("Выполняем -> insert into hh1.test_table (Id, Description) values ................:")
	answer = db.SendRequest(db.POST, client, config, sb.String())
	fmt.Println(answer)

	fmt.Println("Выполняем -> SELECT * FROM hh1.test_table:")
	answer = db.SendRequest(db.GET, client, config, "SELECT * FROM hh1.test_table")
	fmt.Println(answer)

	stopTime := time.Now()

	fmt.Printf("\n\nВыполнено за %f сек.", stopTime.Sub(startTime).Seconds())

}
