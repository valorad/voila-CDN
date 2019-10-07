package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo"
)

type ReplicaFileAddResult struct {
	OK bool `json:"ok" form:"ok" query:"ok"`
	Message string `json:"message" form:"message" query:"message"`
	Url string `json:"url" form:"url" query:"url"`
}

type FileAddResult struct {
	Message string `json:"message" form:"message" query:"message"`
	ReplicaResult ReplicaFileAddResult
}

// get list
func FileIndex(c echo.Context) error {

	client := resty.New()

	resp, _ := client.R().
		EnableTrace().
		Get("http://localhost:8000/api/files")

	return c.JSON(http.StatusOK, resp)

}

func FileAdd(c echo.Context) error {

	client := resty.New()

	resp, _ := client.R().
		EnableTrace().
		SetHeader("Content-Type", "multipart/form-data").
		SetFile("document", "/mnt/m/workbench/_linux/voila-CDN/server-origin/statics/amd.txt").
		Post("http://localhost:8000/api/files")

	fmt.Println(resp.String())

	// res := new(fileAddResult)
	// if err := c.Bind(res); err != nil {
	// 	return c.JSON(http.StatusInternalServerError, "Failed to show result")
	// }

	// result := echo.Map{}
	// if err := c.Bind(&result); err != nil {
	// 	return err
	// }

	// replicaResult := &ReplicaFileAddResult {
	// 	OK: resp.Request.Body,
	// 	Message 
	// 	Url 
	// }

	var replicaResult ReplicaFileAddResult
	json.Unmarshal([]byte(resp.String()), &replicaResult)

  result := &FileAddResult{
    Message:  `Pushed the file "amd.txt" to replica`,
    ReplicaResult: replicaResult,
	}

	return c.JSON(http.StatusOK, result)

}
