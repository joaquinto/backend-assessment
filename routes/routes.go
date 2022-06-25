package routes

import (
	"user/handler"
	"user/prisma/db"
	"user/repository"
	"user/services"

	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func Run(e *echo.Echo, client *db.PrismaClient) *echo.Echo {
	userRepo := repository.NewRepository(client)
	userSvc := services.NewService(userRepo)
	h:= handler.NewHandler(userSvc)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/api/v1/users/:email", h.GetUserByEmail)

	e.GET("/api/v1/users", h.GetUsers)

	e.Any("*", h.RouteNotFound)

	return e
}