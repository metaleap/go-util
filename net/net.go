package net

import (
	"io"
	"net/http"

	uio "github.com/metaleap/go-util/io"
)

//	Downloads a remote file to the specified local file path.
func DownloadFile(fileUrl, filePath string) (err error) {
	var rc io.ReadCloser
	if rc, err = OpenRemoteFile(fileUrl); rc != nil {
		defer rc.Close()
		uio.SaveToFile(rc, filePath)
	}
	return err
}

//	Opens a remote file at the specified (net/http-compatible) fileUrl and returns its io.ReadCloser.
func OpenRemoteFile(fileUrl string) (rc io.ReadCloser, err error) {
	var resp *http.Response
	if resp, err = new(http.Client).Get(fileUrl); resp != nil {
		rc = resp.Body
	}
	return
}
