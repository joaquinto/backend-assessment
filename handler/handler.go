package handler

import "user/services"

type handler struct {
	userServices *services.Services
}

func NewHandler(userSvc *services.Services) *handler {
	return &handler{userServices: userSvc}
}
