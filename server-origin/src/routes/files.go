package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/gabriel-vasile/mimetype"
)

type ReplicaFileAddResult struct {
	OK      bool   `json:"ok" form:"ok" query:"ok"`
	Message string `json:"message" form:"message" query:"message"`
	Url     string `json:"url" form:"url" query:"url"`
}

type FileAddResults struct {
	Results []*FileAddResult
}

type FileAddResult struct {
	Message       string `json:"message" form:"message" query:"message"`
	ReplicaResult ReplicaFileAddResult
}

type FileAddRequest struct {
	Filename string `json:"filename" form:"filename" query:"filename"`
}

type GetFileListRemoteResult struct {
	Host  string                 `json:"host" form:"host" query:"host"`
	Files map[string]interface{} `json:"files" form:"files" query:"files"`
}

type GetFileListRemoteResults struct {
	Results []*GetFileListRemoteResult
}

type FileInfo struct {
	Name  string `json:"name" form:"name" query:"name"`
	Type string`json:"type" form:"type" query:"type"`
}

// file index route
func GETFilesIndex(c echo.Context) error {

	return c.JSON(http.StatusOK, "File works!")

}

// get list
func GETFilesLocal(c echo.Context) error {

	folder := "./statics/files"

	files, err := ioutil.ReadDir(folder)

	if err != nil {
		log.Fatal(err)
	}

	fileList := []*FileInfo{}

	for _, file := range files {

		if !file.IsDir() {

			mime, _, _ := mimetype.DetectFile(folder + "/" + file.Name())

			fileInfo := &FileInfo {
				Name: file.Name(),
				Type: mime,
			}
	
			fileList = append(fileList, fileInfo)
		}

	}

	return c.JSON(http.StatusOK, fileList)

}

// get list
func GETFilesRemote(c echo.Context) error {

	client := resty.New()

	getFileListRemoteResults := []*GetFileListRemoteResult{}

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

		getFileListRemoteResult := &GetFileListRemoteResult{
			Host:  replica,
			Files: replicaResponse,
		}

		getFileListRemoteResults = append(getFileListRemoteResults, getFileListRemoteResult)

	}

	results := &GetFileListRemoteResults{
		Results: getFileListRemoteResults,
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
			SetFile("document", "statics/files/"+fileName).
			Post(replica + "api/files")

		var replicaResult ReplicaFileAddResult

		if err != nil {
			replicaResult := &FileAddResult{
				Message:       `Failed to upload file ` + fileName + ` to replica ` + replica,
				ReplicaResult: replicaResult,
			}
			fmt.Println(err)
			fileAddResults = append(fileAddResults, replicaResult)

		} else {
			json.Unmarshal(resp.Body(), &replicaResult)

			replicaResult := &FileAddResult{
				Message:       `Pushed the file ` + fileName + ` to replica ` + replica,
				ReplicaResult: replicaResult,
			}
			fileAddResults = append(fileAddResults, replicaResult)
		}

	}

	results := &FileAddResults{
		Results: fileAddResults,
	}

	return c.JSON(http.StatusOK, results)
}
