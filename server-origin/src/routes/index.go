package routes

import (
	"os"
	"io/ioutil"
	"fmt"
	"encoding/json"

	"net/http"

	"github.com/labstack/echo/v4"
)

var replicaHosts []string

type SiteConfig struct {
	ReplicaHosts []string `json:"replicaHosts" form:"replicaHosts" query:"replicaHosts"`
}

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "api works!")
}

func init() {

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

	replicaHosts = config.ReplicaHosts

}