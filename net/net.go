package net

import (
	"io"
	"net/http"

	uio "github.com/metaleap/go-util/io"
)

var (
	HttpClient = new(http.Client)
)

func DownloadFile (fileUrl, filePath string) (err error) {
	var rc io.ReadCloser
	if rc, err = OpenRemoteFile(fileUrl); rc != nil {
		defer rc.Close()
		uio.SaveToFile(rc, filePath)
	}
	return err
}

func OpenRemoteFile (fileUrl string) (rc io.ReadCloser, err error) {
	var resp *http.Response
	if resp, err = HttpClient.Get(fileUrl); resp != nil { rc = resp.Body }
	return
}
