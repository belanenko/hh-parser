package db

type UserInfo struct {
	Username string `env:"DB_USERNAME,notEmpty"`
	Password string `env:"DB_PASSWORD,notEmpty"`
}

type Config struct {
	Host       string `env:"DB_HOST,notEmpty"`
	DbName     string `env:"DB_NAME,notEmpty"`
	UserInfo   UserInfo
	CaCertPath string `env:"DB_CERT_PATH,notEmpty"`
}
