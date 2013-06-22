# unet
--
    import "github.com/metaleap/go-util/net"

Various line-savers for common networking needs.

## Usage

#### func  Addr

```go
func Addr(protocol, tcpAddr string) (fullAddr string)
```
Returns a human-readable URL representation of the specified TCP address.

Examples:

`Addr("http", ":8080")` = `http://localhost:8080`

`Addr("https", "testserver:9090")` = `https://testserver:9090`

`Addr("http", ":http")` = `http://localhost`

`Addr("https", "demomachine:https")` = `https://demomachine`

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
	//	Used to implement `http.ResponseWriter.Write()`.
	bytes.Buffer

	//	Used to implement `http.ResponseWriter.Header()`.
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
func (me *ResponseBuffer) WriteHeader(_ int)
```
No-op -- currently, headers aren't written to the underlying `bytes.Buffer`.

--
**godocdown** http://github.com/robertkrimen/godocdown