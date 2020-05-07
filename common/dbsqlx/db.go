package dbsqlx

import (
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type DbConfig struct {
	host   string
	port   int
	user   string
	pass   string
	dbName string
}

/* -- singleton for DI -- */

var dbConfig *DbConfig
var db *sqlx.DB
var once sync.Once

func DB() *sqlx.DB {
	once.Do(func() {
		dbConfig = createDbConfig()
		dbHandle, err := sqlx.Open("mysql", generateDbURL(dbConfig))
		if err != nil {
			fmt.Printf("%v \n", fmt.Errorf("Sqlx error in connectDatabase(): %v", err))
		}
		db = dbHandle
	})

	return db
}

/* ---- */


func createDbConfig() *DbConfig {
	dbConfig = &DbConfig{
		port:   viper.GetInt("db.port"),
		host:   viper.GetString("db.host"),
		user:   viper.GetString("db.user"),
		pass:   viper.GetString("db.pass"),
		dbName: viper.GetString("db.name"),
	}

	return dbConfig
}

func generateDbURL(dbConfig *DbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.user,
		dbConfig.pass,
		dbConfig.host,
		dbConfig.port,
		dbConfig.dbName,
	)
}