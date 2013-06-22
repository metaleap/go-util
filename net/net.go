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

//	Returns a human-readable URL representation of the specified TCP address.
//
//	Examples:
//
//	`Addr("http", ":8080")` = `http://localhost:8080`
//
//	`Addr("https", "testserver:9090")` = `https://testserver:9090`
//
//	`Addr("http", ":http")` = `http://localhost`
//
//	`Addr("https", "demomachine:https")` = `https://demomachine`
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

//	Downloads a remote file at the specified (`net/http`-compatible) `srcFileUrl` to the specified `dstFilePath`.
func DownloadFile(srcFileUrl, dstFilePath string) (err error) {
	var rc io.ReadCloser
	if rc, err = OpenRemoteFile(srcFileUrl); err == nil {
		defer rc.Close()
		uio.SaveToFile(rc, dstFilePath)
	}
	return
}

//	Opens a remote file at the specified (`net/http`-compatible) `srcFileUrl` and returns its `io.ReadCloser`.
func OpenRemoteFile(srcFileUrl string) (src io.ReadCloser, err error) {
	var resp *http.Response
	if resp, err = new(http.Client).Get(srcFileUrl); (err == nil) && (resp != nil) {
		src = resp.Body
	}
	return
}

//	Implements `http.ResponseWriter` with a `bytes.Buffer`.
type ResponseBuffer struct {
	//	Used to implement `http.ResponseWriter.Write()`.
	bytes.Buffer

	//	Used to implement `http.ResponseWriter.Header()`.
	Resp http.Response
}

//	Returns `me.Resp.Header`.
func (me *ResponseBuffer) Header() http.Header { return me.Resp.Header }

//	No-op -- currently, headers aren't written to the underlying `bytes.Buffer`.
func (me *ResponseBuffer) WriteHeader(_ int) {}
