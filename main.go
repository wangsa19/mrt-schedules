package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wangsa19/mrt-schedules/modules/station"
)

func main() {
	InitialRouter()
}

func InitialRouter() {
	var (
		router = gin.Default()
		api    = router.Group("/v1/api")
	)

	station.Initiate(api)

	router.Run(":8080")
}
