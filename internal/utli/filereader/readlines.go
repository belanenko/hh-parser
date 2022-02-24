package filereader

import (
	"log"
	"os"
	"strings"
)

func ReadAllLines(path string) []string {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.ReplaceAll(string(dat), "\r\n", "\n"), "\n")
}
