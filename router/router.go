package router

import (
	"github.com/gin-gonic/gin"
	"github.com/seew0/jobscheduler/controllers"
	"github.com/seew0/jobscheduler/middlewares"
)

type Server struct{
	port string
	router *gin.Engine
}

func NewServer(port string, router *gin.Engine) (*Server){
	return &Server{
		port: port,
		router: router,
	}
}

func (serve *Server)Start(){
	serve.router.Use(middlewares.CORSmanager)

	serve.router.POST("/job/create", func(c *gin.Context){
		controllers.CreateJobHandler(c)
	})
	serve.router.POST("/job/:id/start", func(c *gin.Context){
		controllers.StartJobHandler(c)
	})
	serve.router.GET("/job/:id/status", func(c *gin.Context){
		controllers.GetJobStatusHandler(c)
	})

	serve.router.Run(serve.port)
}
