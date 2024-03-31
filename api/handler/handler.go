package handler

import "github.com/sosamjayo/oauth/api/service"

type Handler struct {
	AuthorizeService service.AuthorizeService
	TokenService     service.TokenService
	CallbackService  service.CallbackService
}

func NewHandler(authorizeSvc service.AuthorizeService, tokenSvc service.TokenService,
	callbackSvc service.CallbackService) *Handler {
	return &Handler{
		AuthorizeService: authorizeSvc,
		TokenService:     tokenSvc,
		CallbackService:  callbackSvc,
	}
}
