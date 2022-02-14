package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

//	Структура хранит данные для подключения к БД - Для изменения базы данных нужно изменить поля
//	в структуре Config и задать соответствующие
//	переменные окружения.
type Config struct {
	PgHost        string `env:"PG_HOST,required"`
	PgPort        int    `env:"PG_PORT,required"`
	PgDB          string `env:"PG_DB_NAME,required"`
	PgUser        string `env:"PG_DB_USER,required"`
	PgPassword    string `env:"PG_DB_PASSWORD,required"`
	PgSSLMode     string `env:"PG_DB_SSL_MODE,required"`
	PgSSLRootCert string `env:"PG_DB_SSL_ROOT_CERT,required"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(
		func() {
			myEnv, err := godotenv.Read()
			if err != nil {
				log.Fatal(err)
			}
			cfg := &config
			opts := &env.Options{
				Environment: myEnv,
			}

			if err := env.Parse(cfg, *opts); err != nil {
				log.Fatal(err)
			}
		})
	return &config
}

