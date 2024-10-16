package http

import (
	"log"
	handler "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/api/handler"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/routes"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler, otpHandler *handler.OtpHandler,
	categoryHandler *handler.CategoryHandler, inventoryHandler *handler.InventaryHandler,
	cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler,
	wishlistHandler *handler.WishlistHandler, offerHandler *handler.OfferHandler, couponHandler *handler.CouponHandler) *ServerHTTP {
	engine := gin.New()
	// logger
	engine.Use(gin.Logger())
	//LOAD HTML PATH
	engine.LoadHTMLGlob("./template/*.html")

	routes.UserRoutes(engine.Group("/users"), userHandler, otpHandler, inventoryHandler, cartHandler, orderHandler, paymentHandler, wishlistHandler, couponHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, categoryHandler, inventoryHandler, offerHandler, couponHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	err := sh.engine.Run(":8000")
	if err != nil {
		log.Fatal("gin engine couldn't Start")
	}
}
