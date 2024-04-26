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

func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler, otpHandler *handler.OtpHandler,
	categoryHandler *handler.CategoryHandler, inventoryHandler *handler.InventaryHandler,
	cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler) *ServerHTTP {
	engine := gin.New()
	// logger
	engine.Use(gin.Logger())
	//LOAD HTML PATH
	engine.LoadHTMLGlob("../template/*.html")

	routes.UserRoutes(engine.Group("/users"), userHandler, otpHandler, inventoryHandler, cartHandler, orderHandler, paymentHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, categoryHandler, inventoryHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	err := sh.engine.Run(":8000")
	if err != nil {
		log.Fatal("gin engine couldn't Start")
	}
}
