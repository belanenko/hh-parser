package fioutil

import (
	"fmt"
	"io/ioutil"
)

func ReadFileToString(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(content)
}
