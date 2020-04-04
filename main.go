package main

import (
	"bettertomorrow/route"
)

func main()  {

	router := route.Init()
	router.Logger.Fatal(router.Start(":8000"))
}
