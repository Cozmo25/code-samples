package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	modelId := "Your-model-id"
	imagePath := "REPLACE_IMAGE_PATH.jpg"
	url := "https://app.nanonets.com/api/v2/ObjectDetection/Model/" + modelId + "/LabelFile/"

	file, err := os.Open(imagePath)
	if err != nil {
		return
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base("REPLACE_IMAGE_PATH.jpg"))
	if err != nil {
		return
	}
	_, err = io.Copy(part, file)

	writer.WriteField("modelId", "REPLACE_MODEL_ID")

	contentType := writer.FormDataContentType()

	err = writer.Close()
	if err != nil {
		return
	}

	req, _ := http.NewRequest("POST", url, body)

	req.Header.Add("Content-Type", contentType)
	req.SetBasicAuth("REPLACE_API_KEY", "")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body2, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body2))

}
