# umgo

Go programming helpers for common MongoDB needs.

## Usage

#### func  ConnectTo

```go
func ConnectTo(url string, safe bool) (conn *mgo.Session, err error)
```
Short-hand for `mgo.Dial` then `Session.SetSafe`.

#### func  ConnectUrl

```go
func ConnectUrl(host string, port int, direct bool) (url string)
```
Returns a connection URL for `ConnectTo`.

#### func  Sparse

```go
func Sparse(m bson.M) bson.M
```
Deletes all zero-value and empty-key entries from `m`, then returns `m`.
