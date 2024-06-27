package routes

import (
	"TugasMID2/constants"
	"TugasMID2/controllers"
	"TugasMID2/middlewares"
	echoJWT "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/register", controllers.RegisterUser)
	e.POST("/login", controllers.LoginUser)

	eAuthJWT := e.Group("/auth")
	eAuthJWT.Use(echoJWT.JWT([]byte(constants.SECRET_JWT)))
	eAuthJWT.GET("/packages", controllers.GetAllPackages)
	eAuthJWT.GET("/packages/:id", controllers.GetPackageByID)
	eAuthJWT.POST("/packages", controllers.CreatePackage)
	eAuthJWT.PUT("/packages/:id", controllers.UpdatePackage)
	eAuthJWT.DELETE("/packages/:id", controllers.DeletePackage)

	middlewares.LogMiddleware(e)
	return e
}
