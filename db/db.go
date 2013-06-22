package udb

import (
	"database/sql"
)

//	Implemented by both `*sql.DB` and `*sql.Tx`.
type Execer interface {
	//	Exec executes a query without returning any rows. The args are for any placeholder parameters in the query.
	Exec(query string, args ...interface{}) (sql.Result, error)
}

//	Implemented by both `*sql.DB` and `*sql.Tx`.
type Querier interface {
	//	Query executes a query that returns rows. The args are for any placeholder parameters in the query.
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

//	Exec calls `execer.Exec(query, args...)` and from the returned `sql.Result`,
//	returns either its `LastInsertId()` (if `isInsert` is `true`) or its `RowsAffected()`.
func Exec(execer Execer, isInsert bool, query string, args ...interface{}) (result int64, err error) {
	var res sql.Result
	if res, err = execer.Exec(query, args...); err == nil {
		if isInsert {
			result, err = res.LastInsertId()
		} else {
			result, err = res.RowsAffected()
		}
	}
	return
}

//	Helps iterating over `*sql.Rows`. After each `rows.Next()` call, you can
//	call the `Scan()` method to return the record set as a `map[string]interface{}`.
//
//	If the columns are identical for all records (most SQL databases), call
//	`PrepareColumns()` just once prior to iteration. If the columns vary across
//	records (some NoSQL databases), call `PrepareColumns()` during each iteration
//	prior to `Scan()`.
type SqlCursor struct {
	cols       []string
	vals, ptrs []interface{}
}

//	Retrieves meta-data information about the `rows.Columns()`.
func (me *SqlCursor) PrepareColumns(rows *sql.Rows) (err error) {
	if me.cols, err = rows.Columns(); err == nil {
		me.vals, me.ptrs = make([]interface{}, len(me.cols), len(me.cols)), make([]interface{}, len(me.cols), len(me.cols))
		for i := 0; i < len(me.ptrs); i++ {
			me.ptrs[i] = &me.vals[i]
		}
	}
	return
}

//	According to the meta-data retrieved during your prior-most call to `me.PrepareColumns()`,
//	populates `rec` with all field values for the current record in the specified `rows`.
func (me *SqlCursor) Scan(rows *sql.Rows) (rec map[string]interface{}, err error) {
	if err = rows.Scan(me.ptrs...); err == nil {
		rec = make(map[string]interface{}, len(me.vals))
		var bytes []byte
		var ok bool
		for i := 0; i < len(me.vals); i++ {
			if bytes, ok = me.vals[i].([]byte); ok {
				me.vals[i] = string(bytes)
			}
			rec[me.cols[i]] = me.vals[i]
		}
	}
	return
}
