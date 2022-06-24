package routes

import (
	"user/handler"
	"user/prisma/db"
	"user/repository"
	"user/services"

	"github.com/labstack/echo/v4"
)

func Run(e *echo.Echo, client *db.PrismaClient) *echo.Echo {
	userRepo := repository.NewRepository(client)
	userSvc := services.NewService(userRepo)
	h:= handler.NewHandler(userSvc)

	e.GET("/api/v1/users/:email", h.GetUserByEmail)

	e.GET("/api/v1/users", h.GetUsers)

	e.Any("*", h.RouteNotFound)

	return e
}