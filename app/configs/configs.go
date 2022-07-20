package configs

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	Name string
	Env  string
}

type DbConfig struct {
	Host     string
	Port     string
	Dbname   string
	Username string
	Password string
}

type Configs struct {
	Appconfig AppConfig
	Dbconfig  DbConfig
}

var lock = &sync.Mutex{}
var configs *Configs

func GetInstance() *Configs {
	fmt.Println("configs", configs)
	if configs == nil {
		lock.Lock()
		configs = &Configs{
			Appconfig: AppConfig{
				Name: os.Getenv("APP_NAME"),
				Env:  os.Getenv("APP_ENV"),
			},
			Dbconfig: DbConfig{
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				Dbname:   os.Getenv("DB_NAME"),
				Username: os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASS"),
			},
		}
		lock.Unlock()
	}
	return configs
}
