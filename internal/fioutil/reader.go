package fioutil

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadFileToString(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	return string(content)
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
