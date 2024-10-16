package mysql

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/the-arcade-01/go-es-search/server/internal/config"
)

var (
	once        sync.Once
	mysqlClient *sql.DB
)

func NewMysqlConnection(appConfig *config.AppConfig) *sql.DB {
	var err error
	once.Do(func() {
		mysqlClient, err = sql.Open(appConfig.MysqlDriverName, appConfig.MysqlDataSource)
		if err != nil {
			log.Fatalln("[NewMysqlConnection] error initializing db connection, err: ", err)
		}
		if err := mysqlClient.Ping(); err != nil {
			log.Fatalln("[NewMysqlConnection] error at ping, err: ", err)
		}
		log.Println("[NewMysqlConnection] Mysql initialized")
	})
	return mysqlClient
}
