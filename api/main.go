package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sosamjayo/oauth/api/config"
	"github.com/sosamjayo/oauth/api/handler"
	"github.com/sosamjayo/oauth/api/service"
)

func main() {
	fmt.Println("OAuth Server is running...")

	ctx := context.Background()
	_, err := config.NewRedisClient(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// DB
	// accessTokenDB := db.NewAccessTokenRedis(client)
	// authCodeDB := db.NewAuthCodeRedis(client)
	// clientDB := db.NewClientRedis(client)

	// Service
	authorizeHandler := service.NewAuthorizeService()
	tokenHandler := service.NewTokenService()
	callbackHandler := service.NewCallbackService()

	// Handler
	h := handler.NewHandler(authorizeHandler, tokenHandler, callbackHandler)

	http.HandleFunc("/authorize", h.AuthorizeHandler)
	http.HandleFunc("/token", h.TokenHandler)
	http.HandleFunc("/callback", h.CallbackHandler)
	http.ListenAndServe(":8080", nil)
}
