package main

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/valorad/voila-CDN/server-origin/src/routes"
)

type SiteConfig struct {
	Host string `json:"host" form:"host" query:"host"`
}

var host string

func main() {
	loadConfig()
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/api", routes.Index)
	e.GET("/api/files", routes.GETFilesIndex)
	e.GET("/api/files/local", routes.GETFilesLocal)
	e.GET("/api/files/remote", routes.GETFilesRemote)
	e.POST("/api/files", routes.FileAdd)

	e.Static("/*", "statics")
	e.File("/", "statics/client/index.html")

	// Start server
	e.Logger.Fatal(e.Start(host))
}

func loadConfig() {

	configFile, err := os.Open("./configs/voila-CDN-origin.json");

	if err != nil {
		fmt.Println(err)
	}

	defer configFile.Close()

	bytes, _ := ioutil.ReadAll(configFile)

	var config SiteConfig

	err = json.Unmarshal([]byte(bytes), &config)

	if err != nil {
		fmt.Println(err)
	}

	host = config.Host

}