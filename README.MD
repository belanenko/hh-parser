# hh.ru parse tool 

---

## Start flags

1. `-threads` sets parsing-gorutines count. <br/>Example: `-threads 4`.
2. `-config` path to *.env config file. <br/>Example: `-config "~/app/config.env"`
3. `-startIndex` start vacancy id position  <br/>Example: `-startIndex 50000000`
4. `-conutVacancy` count vacancies which will be parsed<br/>Example: `-conutVacancy 10000`

### Example flags: 
> `app.go -threads 4 -config "~/app/config.env" -startIndex 50000000 -conutVacancy 10000`

---

## Config file params   

### for clickhouse 

1. `DB_NAME`
2. `DB_HOST`<br/>Example: `DB_HOST=https://rc1a-awh1aqqz9v7110od.mdb.yandexcloud.net:8080/`
3. `DB_CERT_PATH` <br/>Example: `DB_CERT_PATH=/usr/local/share/ca-certificates/Yandex/YandexInternalRootCA.crt`
4. `DB_USERNAME`
5. `DB_PASSWORD`

---

