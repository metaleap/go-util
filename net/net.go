package unet

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	ugo "github.com/metaleap/go-util"
	uio "github.com/metaleap/go-util/io"
)

func Addr(protocol, tcpAddr string) (fullAddr string) {
	localhost := ugo.HostName()
	both := strings.Split(tcpAddr, ":")
	if len(both) < 1 {
		both = []string{localhost}
	} else if len(both[0]) == 0 {
		both[0] = localhost
	}
	if len(both) > 1 {
		if both[1] == protocol {
			both[1] = ""
		}
		if len(both[1]) == 0 {
			both = both[:1]
		}
	}
	if fullAddr = strings.Join(both, ":"); len(protocol) > 0 {
		fullAddr = fmt.Sprintf("%s://%s", protocol, fullAddr)
	}
	return
}

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

//	Implements http.ResponseWriter with a bytes.Buffer
type ResponseBuffer struct {
	//	Used to implement http.ResponseWriter.Write()
	bytes.Buffer

	//	Used to implement http.ResponseWriter.Header()
	Resp http.Response
}

//	Returns me.Resp.Header
func (me *ResponseBuffer) Header() http.Header { return me.Resp.Header }

//	No-op
func (me *ResponseBuffer) WriteHeader(_ int) {}
