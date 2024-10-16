package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
	"github.com/the-arcade-01/go-es-search/server/internal/utils"
)

type AppConfig struct {
	DbType          string
	MysqlDriverName string
	MysqlDataSource string
}

var (
	once      sync.Once
	appConfig *AppConfig
)

func init() {
	FetchAndLoadConfig()
}

func FetchAndLoadConfig() *AppConfig {
	once.Do(func() {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalln("[FetchAndLoadConfig] error getting current working directory, err:", err)
		}
		envPath := filepath.Join(cwd, ".env")
		log.Println("[FetchAndLoadConfig] Loading env file from:", envPath)

		err = godotenv.Load(envPath)
		if err != nil {
			log.Fatalln("[FetchAndLoadConfig] error loading env props, err: ", err)
		}
		appConfig = &AppConfig{
			DbType:          os.Getenv(utils.DB_TYPE),
			MysqlDriverName: os.Getenv(utils.MYSQL_DRIVER_NAME),
			MysqlDataSource: os.Getenv(utils.MYSQL_DATA_SOURCE),
		}
		log.Println("[FetchAndLoadConfig] AppConfig initialized")
	})
	return appConfig
}
