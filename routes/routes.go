package routes

import (
	"main/handlers"
	"main/middlewares"

	"github.com/gin-gonic/gin"
)

type RouteType string

const (
	GET    RouteType = "GET"
	POST   RouteType = "POST"
	PUT    RouteType = "PUT"
	DELETE RouteType = "DELETE"
)

type Route struct {
	addr       string
	routeType  RouteType
	handler    func(context *gin.Context)
	middleware gin.HandlerFunc
}

func SetupRoutes(router *gin.Engine) {
	routes := []Route{
		{
			addr:      "/saluta",
			routeType: GET,
			handler:   handlers.SalutaHandler,
		},
		{
			addr:       "/addTodo",
			routeType:  POST,
			handler:    handlers.AddTodoHandler,
			middleware: middlewares.AuthMiddleware(),
		},
		{
			addr:       "/getTodos",
			routeType:  GET,
			handler:    handlers.GetTodosHandler,
			middleware: middlewares.AuthMiddleware(),
		},
		{
			addr:       "/getTodosOfUser",
			routeType:  GET,
			handler:    handlers.GetTodosOfUserHandler,
			middleware: middlewares.AuthMiddleware(),
		},
		{
			addr:      "/register",
			routeType: POST,
			handler:   handlers.RegisterHandler,
		},
	}

	for _, route := range routes {
		if route.middleware == nil {
			switch route.routeType {
			case "GET":
				router.GET(route.addr, route.handler)
			case "POST":
				router.POST(route.addr, route.handler)
			case "PUT":
				router.PUT(route.addr, route.handler)
			case "DELETE":
				router.DELETE(route.addr, route.handler)
			}
		} else {
			switch route.routeType {
			case "GET":
				router.GET(route.addr, route.middleware, route.handler)
			case "POST":
				router.POST(route.addr, route.middleware, route.handler)
			case "PUT":
				router.PUT(route.addr, route.middleware, route.handler)
			case "DELETE":
				router.DELETE(route.addr, route.middleware, route.handler)
			}
		}
	}
}
