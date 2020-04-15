package dbxorm

import (
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
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
var db *xorm.Engine
var once sync.Once

func DB() *xorm.Engine {
	once.Do(func() {
		dbConfig = createDbConfig()
		dbUrl := generateDbURL(dbConfig)
		//TODO test onyly
		fmt.Printf("dbUrl : %v \n", dbUrl)
		dbCon, err := xorm.NewEngine("mysql", generateDbURL(dbConfig))
		if err != nil {
			fmt.Printf("%v \n", fmt.Errorf("error in connectDatabase(): %v", err))
		}
		db = dbCon
		db.ShowSQL(true)
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