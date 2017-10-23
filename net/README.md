# unet

Go programming helpers for common networking needs.

## Usage

#### func  Addr

```go
func Addr(protocol, tcpAddr string) (fullAddr string)
```
Returns a human-readable URL representation of the specified TCP address.

Examples:

`unet.Addr("http", ":8080")` = `http://localhost:8080`

`unet.Addr("https", "testserver:9090")` = `https://testserver:9090`

`unet.Addr("http", ":http")` = `http://localhost`

`unet.Addr("https", "demomachine:https")` = `https://demomachine`

#### func  DownloadFile

```go
func DownloadFile(srcFileUrl, dstFilePath string) (err error)
```
Downloads a remote file at the specified (`net/http`-compatible) `srcFileUrl` to
the specified `dstFilePath`.

#### func  OpenRemoteFile

```go
func OpenRemoteFile(srcFileUrl string) (src io.ReadCloser, err error)
```
Opens a remote file at the specified (`net/http`-compatible) `srcFileUrl` and
returns its `io.ReadCloser`.

#### type ResponseBuffer

```go
type ResponseBuffer struct {
	//	Used to implement the `http.ResponseWriter.Write` method.
	bytes.Buffer

	//	Used to implement the `http.ResponseWriter.Header` method.
	Resp http.Response
}
```

Implements `http.ResponseWriter` with a `bytes.Buffer`.

#### func (*ResponseBuffer) Header

```go
func (me *ResponseBuffer) Header() http.Header
```
Returns `me.Resp.Header`.

#### func (*ResponseBuffer) WriteHeader

```go
func (_ *ResponseBuffer) WriteHeader(_ int)
```
No-op -- currently, headers aren't written to the underlying `bytes.Buffer`.
