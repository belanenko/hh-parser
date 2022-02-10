package main

import (
	"fmt"

	"github.com/hh-parser/internal/db"
)

func main() {
	var config db.Config
	db.Configuration("./../conf.json", &config)

	client := db.CreateClient(config.CertPath)

	fmt.Println("Выполняем -> SELECT * FROM hh1.test_table:")
	answer := db.SendRequest(db.GET, client, config, "SELECT * FROM hh1.test_table")
	fmt.Println(answer)

	fmt.Println("Выполняем -> insert into hh1.test_table (Id, Description) values (1, 'тестовое описание вакансии под номером 1'):")
	answer = db.SendRequest(db.POST, client, config, "insert into hh1.test_table (Id, Description) values (1, 'тестовое описание вакансии под номером 1')")
	fmt.Println(answer)

	fmt.Println("Выполняем -> SELECT * FROM hh1.test_table:")
	answer = db.SendRequest(db.GET, client, config, "SELECT * FROM hh1.test_table")
	fmt.Println(answer)
}
