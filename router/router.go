package router

import (
	"iam-server/api/v1"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(r *echo.Echo) {
	r.Use(RequestLogging)
	r.Use(ResponseLogging)

	api := r.Group("/api")

	v1 := api.Group("/v1")
	v1Routes(v1)

	// v2 := api.Group("/v2")
	// v2Routes(v2)

	// v3 := api.Group("/v3")
	// v3Routes(v3)
}

func v1Routes(rg *echo.Group) {
	rg.POST("/signIn", api.Login)
	rg.POST("/signUp", api.Register)
	// rg.POST("/signUp", api.Register)
	// rg.POST("/verifyToken", api.VerifyToken, Auth())
}
