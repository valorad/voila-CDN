package routes

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
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

type FileAddRequest struct {
	Filename string `json:"filename" form:"filename" query:"filename"`
}

// get list
func FileIndex(c echo.Context) error {

	client := resty.New()

	resp, _ := client.R().
		EnableTrace().
		Get("http://localhost:8000/api/files")

	return c.JSON(http.StatusOK, resp)

}

func FileAdd(c echo.Context) (routeError error) {

  req := new(FileAddRequest)
  if err := c.Bind(&req); err != nil {
    return
  }

  fileName := req.Filename

  // Post the file to replica

	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "multipart/form-data").
		SetFile("document", "statics/" + fileName).
		Post("http://localhost:8000/api/files")

	var replicaResult ReplicaFileAddResult

	if err != nil {
		result := &FileAddResult{
			Message:  `Failed to upload file `,
			ReplicaResult: replicaResult,
		}
		return c.JSON(http.StatusInternalServerError, result)
	} else {
		json.Unmarshal(resp.Body(), &replicaResult)
	
		result := &FileAddResult{
			Message:  `Pushed the file ` + fileName + ` to replica`,
			ReplicaResult: replicaResult,
		}
	
		return c.JSON(http.StatusOK, result)
	}

}
