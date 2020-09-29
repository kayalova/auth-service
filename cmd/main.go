package main

import (
	"github.com/kayalova/auth-service/router"
)

func main() {
	engine := router.InitRoutes()
	router.Run(engine)
}
