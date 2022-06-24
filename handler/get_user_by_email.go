package handler

import (
	"errors"
	"log"
	"net/http"
	"net/mail"
	"user/domain"
	"user/prisma/db"

	"github.com/labstack/echo/v4"
)

func (h handler) GetUserByEmail(ctx echo.Context) error {
	email := ctx.Param("email")

	if !isValidEmail(email) {
		return ctx.JSON(http.StatusBadRequest, domain.BaseResponse{
			Status: "error",
			Message: "Invalid email address",
		})
	}

	user, err := h.userServices.FindUserByEmail(ctx.Request().Context(), email)
	if err != nil {
		log.Println(err)
		if errors.Is(err, db.ErrNotFound) {
			return ctx.JSON(http.StatusNotFound, domain.BaseResponse{
				Status:  "error",
				Message: "User not found",
			})
		}

		return ctx.JSON(http.StatusInternalServerError, domain.BaseResponse{
			Status:  "error",
			Message: "Unable to retrieve user data",
		})
	}

	return ctx.JSON(http.StatusOK, domain.UserSuccessResponse{
		BaseResponse: domain.BaseResponse{
			Status:  "success",
			Message: "Successfully retrieved user",
		},
		Data: user,
	})
}

func isValidEmail(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}
