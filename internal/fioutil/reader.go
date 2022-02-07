package fioutil

import (
	"io/ioutil"
	"log"
)

func ReadFileToString(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	return string(content)
}
