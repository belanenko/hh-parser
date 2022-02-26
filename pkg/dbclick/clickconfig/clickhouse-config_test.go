package clickconfig

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	filePath := "./test.json"
	os.Remove(filePath)

	text := `{
		"DB_USERNAME": "username",
		"DB_PASSWORD": "passwd",
		"DB_HOST": "hosst",
		"DB_NAME": "dbname",
		"DB_CERT_PATH": "/ddddd/d"
	}`

	if err := os.WriteFile(filePath, []byte(text), 0644); err != nil {
		log.Fatalln(err)
	}

	config := ReadConfig(filePath)
	assert.EqualValues(t, "username", config.DbUsername)
	os.Remove(filePath)

}
