package main

import (
	"github.com/gin-gonic/gin"
	"github.com/seew0/jobscheduler/router"
)

func main() {
	Router := gin.Default()
	Server := router.NewServer(":8080", Router)
	Server.Start()
}
