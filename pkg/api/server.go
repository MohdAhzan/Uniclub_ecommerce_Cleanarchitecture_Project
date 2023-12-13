package http

import (
	"log"
	handler "project/pkg/api/handler"
	"project/pkg/routes"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler) *ServerHTTP {
	engine := gin.New()
	// logger
	engine.Use(gin.Logger())

	routes.UserRoutes(engine.Group("/users"), userHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	err := sh.engine.Run(":8000")
	if err != nil {
		log.Fatal("gin engine couldn't Start")
	}
}
