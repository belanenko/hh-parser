package db

type Config struct {
	Host     string `json:"host"`
	CertPath string `json:"certPath"`
	DbName   string `json:"dbName"`
	User     struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	} `json:"user"`
}
