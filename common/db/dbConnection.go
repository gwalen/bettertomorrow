package db

import "sync"

type DbConnection struct {
	port    int
	address string
	user    string
	pass    string
}

/* -- singleton for DI -- */

var dbConnectionInstance *DbConnection
var once sync.Once

func ProvideDbConnection() *DbConnection {
	once.Do(func() {
		dbConnectionInstance = &DbConnection{
			port:    123,
			address: "127.0.0.1",
			user:    "gw",
			pass:    "gw",
		}
	})
	return dbConnectionInstance
}

/* ---- */
