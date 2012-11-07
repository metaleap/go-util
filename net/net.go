package net

import (
	"io"
	"net/http"
	"os"
)

var (
	HttpClient = new(http.Client)
)

func DownloadFile (fileUrl, filePath string) error {
	var resp, err = HttpClient.Get(fileUrl)
	var file *os.File
	if (resp != nil) && (resp.Body != nil) {
		defer resp.Body.Close()
		if err == nil {
			file, err = os.Create(filePath)
			if file != nil {
				defer file.Close()
				if err == nil {
					_, err = io.Copy(file, resp.Body)
				}
			}
		}
	}
	return err
}
