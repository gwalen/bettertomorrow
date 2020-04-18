package dbsqlx

import (
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
		dbUrl := generateDbURL(dbConfig)
		//TODO test onyly
		fmt.Printf("dbUrl : %v \n", dbUrl)
		dbHandle, err := sqlx.Open("mysql", generateDbURL(dbConfig))
		if err != nil {
			fmt.Printf("%v \n", fmt.Errorf("error in connectDatabase(): %v", err))
		}
		db = dbHandle
		// db.ShowSQL(true) //TODO: howto in sqlx //https://stackoverflow.com/questions/33041063/how-can-i-log-all-outgoing-sql-statements-from-go-mysql
	})

	return db
}

/* ---- */


func createDbConfig() *DbConfig {
	dbConfig = &DbConfig{
		port:   3306,
		host:   "127.0.0.1",
		user:   "gw",
		pass:   "gw",
		dbName: "bettertomorrow",
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