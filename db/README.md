# udb
--
    import "github.com/metaleap/go-util/db"

Various line-savers for common database needs.

## Usage

#### func  Exec

```go
func Exec(execer Execer, isInsert bool, query string, args ...interface{}) (result int64, err error)
```
Exec calls `execer.Exec(query, args...)` and from the returned `sql.Result`,
returns either its `LastInsertId()` (if `isInsert` is `true`) or its
`RowsAffected()`.

#### type Execer

```go
type Execer interface {
	//	Exec executes a query without returning any rows. The args are for any placeholder parameters in the query.
	Exec(query string, args ...interface{}) (sql.Result, error)
}
```

Implemented by both `*sql.DB` and `*sql.Tx`.

#### type Querier

```go
type Querier interface {
	//	Query executes a query that returns rows. The args are for any placeholder parameters in the query.
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
```

Implemented by both `*sql.DB` and `*sql.Tx`.

#### type SqlCursor

```go
type SqlCursor struct {
}
```

Helps iterating over `*sql.Rows`. After each `rows.Next()` call, you can call
the `Scan()` method to return the record set as a `map[string]interface{}`.

If the columns are identical for all records (most SQL databases), call
`PrepareColumns()` just once prior to iteration. If the columns vary across
records (some NoSQL databases), call `PrepareColumns()` during each iteration
prior to `Scan()`.

#### func (*SqlCursor) PrepareColumns

```go
func (me *SqlCursor) PrepareColumns(rows *sql.Rows) (err error)
```
Retrieves meta-data information about the `rows.Columns()`.

#### func (*SqlCursor) Scan

```go
func (me *SqlCursor) Scan(rows *sql.Rows) (rec map[string]interface{}, err error)
```
According to the meta-data retrieved during your prior-most call to
`me.PrepareColumns()`, populates `rec` with all field values for the current
record in the specified `rows`.

--
**godocdown** http://github.com/robertkrimen/godocdown