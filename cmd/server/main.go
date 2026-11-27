package main

import (
	"fmt"

	"prob13/internal/core/entity"
	"prob13/internal/handler"
	"prob13/internal/repository/db"

	"github.com/gin-gonic/gin"
)

// forming a server

type PortalRouter struct {
	engine         *gin.Engine
	healthHandler  *handler.HealthHandlerImpl
	userHandler    *handler.UserHandlerImpl
	channelHandler *handler.ChanRelayHandlerImpl
}

func (r *PortalRouter) registerHealthRoute() {
	r.engine.GET("/health", r.healthHandler.GetStatus)
}

func (r *PortalRouter) registerUserRoute() {
	r.engine.POST("/user", r.userHandler.RegisterUser)
	r.engine.GET("/user/:id", r.userHandler.GetUser)
	r.engine.POST("/userGroup", r.userHandler.RegisterUserGroup)
	r.engine.GET("/userGroup/:id", r.userHandler.GetUserGroup)
}

func (r *PortalRouter) registerChannelHandler() {
	r.engine.POST("/send", r.channelHandler.RelayMessage)
	r.engine.GET("/consume", r.channelHandler.ConsumeMessage)
}

func main() {
	g := gin.New()
	db := db.ConnectYugabyte()
	chann := make(chan entity.Message, 100)
	healthHandler := handler.NewHandlerImpl(200, db)
	userHandler := handler.NewUserHandlerImpl(db)
	chanHandler := handler.NewChannelHandlerImpl(db, chann)
	portalRouter := PortalRouter{g, &healthHandler, &userHandler, &chanHandler}
	portalRouter.registerHealthRoute()
	portalRouter.registerUserRoute()
	portalRouter.registerChannelHandler()

	fmt.Println("Starting the server")
	g.Run(":8009")
}