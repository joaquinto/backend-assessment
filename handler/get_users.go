package handler

import (
	"net/http"
	"strconv"
	"time"
	"user/domain"

	"github.com/labstack/echo/v4"
)

func (h handler) GetUsers(ctx echo.Context) error {
	var to, from time.Time
	var limit, pageNumber int64
	var err error

	lq := ctx.QueryParam("limit")
	pnq := ctx.QueryParam("pageNumber")
	fq := ctx.QueryParam("from")
	tq := ctx.QueryParam("to")

	if lq != "" {
		limit, err = strconv.ParseInt(lq, 10, 64)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, domain.BaseResponse{
				Status:  "error",
				Message: "Invalid limit format",
			})
		}
	}

	if pnq != "" {
		pageNumber, err = strconv.ParseInt(pnq, 10, 64)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, domain.BaseResponse{
				Status:  "error",
				Message: "Invalid limit format",
			})
		}
	}

	if fq != "" {
		from, err = time.Parse("2006-01-02", fq)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, domain.BaseResponse{
				Status:  "error",
				Message: "Invalid date range format",
			})
		}
	}

	if tq != "" {
		to, err = time.Parse("2006-01-02", tq)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, domain.BaseResponse{
				Status:  "error",
				Message: "Invalid date range format",
			})
		}
	}

	users, totalRecord, err := h.userServices.FindUsers(ctx.Request().Context(), int(limit), int(pageNumber), from, to)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, domain.BaseResponse{
			Status:  "error",
			Message: "Unabe to retrieve users",
		})
	}

	return ctx.JSON(http.StatusOK, domain.UsersSuccessResponse{
		BaseResponse: domain.BaseResponse{
			Status:  "success",
			Message: "successfully retrieved users",
		},
		TotalRecord: totalRecord,
		Data:        users,
	})
}
