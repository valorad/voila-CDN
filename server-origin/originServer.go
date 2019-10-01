package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"./routes"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/api", routes.Index)
	// e.GET("/index", rdIndex)

	e.Static("/*", "statics")
	e.File("/", "statics/index.html")

	// Start server
	e.Logger.Fatal(e.Start(":3399"))
}

// Handler
// func apiH(c echo.Context) error {
// 	return c.String(http.StatusOK, "api works!")
// }

// func rdIndex(c echo.Context) error {
// 	return c.Redirect(http.StatusMovedPermanently, "../")
// }
