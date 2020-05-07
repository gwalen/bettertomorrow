package main

import (
	_ "bettertomorrow/common/configuration"  // run init in configration package
	"bettertomorrow/route"
)

func main()  {
	router := route.Init()
	router.Logger.Fatal(router.Start(":8000"))
}

