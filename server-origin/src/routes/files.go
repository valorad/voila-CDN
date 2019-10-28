package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

type ReplicaFileAddResult struct {
	OK bool `json:"ok" form:"ok" query:"ok"`
	Message string `json:"message" form:"message" query:"message"`
	Url string `json:"url" form:"url" query:"url"`
}

type FileAddResults struct {
	Results []*FileAddResult
}

type FileAddResult struct {
	Message string `json:"message" form:"message" query:"message"`
	ReplicaResult ReplicaFileAddResult
}

type FileAddRequest struct {
	Filename string `json:"filename" form:"filename" query:"filename"`
}

type FileGetListResult struct {
	Host string `json:"host" form:"host" query:"host"`
	Files map[string]interface{} `json:"files" form:"files" query:"files"`
}

type FileGetListResults struct {
	Results []*FileGetListResult
}

// get list
func FileIndex(c echo.Context) error {

	client := resty.New()

	fileGetListResults := []*FileGetListResult{}

	for _, replica := range replicaHosts {
		resp, err := client.R().
		EnableTrace().
		Get(replica + "api/files")
		var replicaResponse map[string]interface{}

		if err != nil {
			fmt.Println(err)
		} else {
			json.Unmarshal(resp.Body(), &replicaResponse)
		}

		fileGetListResult := &FileGetListResult{
			Host: replica,
			Files: replicaResponse,
		}
		
		fileGetListResults = append(fileGetListResults, fileGetListResult)

	}

	results := &FileGetListResults {
		Results: fileGetListResults,
	}

	return c.JSON(http.StatusOK, results)

}

func FileAdd(c echo.Context) (routeError error) {

  req := new(FileAddRequest)
  if err := c.Bind(&req); err != nil {
    return
  }

  fileName := req.Filename

  // Post the file to replica list

	client := resty.New()

	fileAddResults := []*FileAddResult{}

	for _, replica := range replicaHosts {
		resp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "multipart/form-data").
		SetFile("document", "statics/" + fileName).
		Post(replica + "api/files")

		var replicaResult ReplicaFileAddResult

		if err != nil {
			replicaResult := &FileAddResult{
				Message:  `Failed to upload file ` + fileName + ` tp replica ` + replica,
				ReplicaResult: replicaResult,
			}
			fmt.Println(err)
			fileAddResults = append(fileAddResults, replicaResult)
			
		} else {
			json.Unmarshal(resp.Body(), &replicaResult)

			replicaResult := &FileAddResult{
				Message:  `Pushed the file ` + fileName + ` to replica ` + replica,
				ReplicaResult: replicaResult,
			}
			fileAddResults = append(fileAddResults, replicaResult)
		}

	}

	results := &FileAddResults {
		Results: fileAddResults,
	}

	return c.JSON(http.StatusOK, results)
}