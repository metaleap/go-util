# unet
--
    import "github.com/metaleap/go-util/net"

Various line-savers for common networking needs.

## Usage

#### func  Addr

```go
func Addr(protocol, tcpAddr string) (fullAddr string)
```

#### func  DownloadFile

```go
func DownloadFile(fileUrl, filePath string) (err error)
```
Downloads a remote file to the specified local file path.

#### func  OpenRemoteFile

```go
func OpenRemoteFile(fileUrl string) (rc io.ReadCloser, err error)
```
Opens a remote file at the specified (net/http-compatible) fileUrl and returns
its io.ReadCloser.

#### type ResponseBuffer

```go
type ResponseBuffer struct {
	//	Used to implement http.ResponseWriter.Write()
	bytes.Buffer

	//	Used to implement http.ResponseWriter.Header()
	Resp http.Response
}
```

Implements http.ResponseWriter with a bytes.Buffer

#### func (*ResponseBuffer) Header

```go
func (me *ResponseBuffer) Header() http.Header
```
Returns me.Resp.Header

#### func (*ResponseBuffer) WriteHeader

```go
func (me *ResponseBuffer) WriteHeader(_ int)
```
No-op

--
**godocdown** http://github.com/robertkrimen/godocdown