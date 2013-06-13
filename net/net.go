package unet

import (
	"bytes"
	"io"
	"net/http"

	uio "github.com/metaleap/go-util/io"
)

//	Implements http.ResponseWriter with a bytes.Buffer
type ResponseBuffer struct {
	bytes.Buffer

	//	Used to return Resp.Header in Header()
	Resp http.Response
}

//	Returns me.Resp.Header
func (me *ResponseBuffer) Header() http.Header { return me.Resp.Header }

//	No-op
func (me *ResponseBuffer) WriteHeader(_ int) {}

//	Downloads a remote file to the specified local file path.
func DownloadFile(fileUrl, filePath string) (err error) {
	var rc io.ReadCloser
	if rc, err = OpenRemoteFile(fileUrl); err == nil {
		defer rc.Close()
		uio.SaveToFile(rc, filePath)
	}
	return
}

//	Opens a remote file at the specified (net/http-compatible) fileUrl and returns its io.ReadCloser.
func OpenRemoteFile(fileUrl string) (rc io.ReadCloser, err error) {
	var resp *http.Response
	if resp, err = new(http.Client).Get(fileUrl); (err == nil) && (resp != nil) {
		rc = resp.Body
	}
	return
}
