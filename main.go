package main

import (
	"log"
	"net/http"

	"wicoady1/gojakarta-middleware/controller"
	"wicoady1/gojakarta-middleware/middleware"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	middlewares := initMiddlewares(middleware.LogMiddleware, middleware.AuthMiddleware, middleware.CustomerIDValidatorMiddleware)

	router.GET("/v1/cart/:id", middleware.MultipleMiddlewares(controller.GetCart, middlewares...))
	router.PUT("/v1/cart/", middleware.MultipleMiddlewares(controller.InitCart, middlewares...))
	router.POST("/v1/cart/items", middleware.MultipleMiddlewares(controller.AddItems, middlewares...))
	router.DELETE("/v1/cart/items/:id", middleware.MultipleMiddlewares(controller.DeleteItems, middlewares...))

	//section 1
	/*
		router.GET("/v1/cart/:id", middleware.CustomerIDValidatorMiddleware(controller.GetCart))
		router.PUT("/v1/cart/", middleware.CustomerIDValidatorMiddleware(controller.InitCart))
		router.POST("/v1/cart/items", middleware.CustomerIDValidatorMiddleware(controller.AddItems))
		router.DELETE("/v1/cart/items/:id", middleware.CustomerIDValidatorMiddleware(controller.DeleteItems))
	*/

	//section 2
	/*
		router.GET("/v1/cart/:id", middleware.LogMiddleware(middleware.AuthMiddleware(middleware.CustomerIDValidatorMiddleware(controller.GetCart))))
		router.PUT("/v1/cart/", middleware.LogMiddleware(middleware.AuthMiddleware(middleware.CustomerIDValidatorMiddleware(controller.InitCart))))
		router.POST("/v1/cart/items", middleware.LogMiddleware(middleware.AuthMiddleware(middleware.CustomerIDValidatorMiddleware(controller.AddItems))))
		router.DELETE("/v1/cart/items/:id", middleware.LogMiddleware(middleware.AuthMiddleware(middleware.CustomerIDValidatorMiddleware(controller.DeleteItems))))
	*/

	log.Fatal(http.ListenAndServe(":6969", router))
}

func initMiddlewares(m ...middleware.Middleware) []middleware.Middleware {
	middleware := []middleware.Middleware{}

	for _, v := range m {
		middleware = append(middleware, v)
	}

	return middleware
}
