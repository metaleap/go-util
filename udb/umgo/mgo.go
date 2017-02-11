package umgo

import (
	"fmt"

	"github.com/go-utils/ugo"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

//	Short-hand for `mgo.Dial` then `Session.SetSafe`.
func ConnectTo(url string, safe bool) (conn *mgo.Session, err error) {
	if conn, err = mgo.Dial(url); conn != nil {
		if err != nil {
			conn.Close()
			conn = nil
		} else if safe {
			conn.SetSafe(&mgo.Safe{})
		}
	}
	return conn, err
}

//	Returns a connection URL for `ConnectTo`.
func ConnectUrl(host string, port int, direct bool) (url string) {
	return fmt.Sprintf("%s:%d%s", host, port, ugo.Ifs(direct, "?connect=direct", ""))
}

//	Deletes all zero-value and empty-key entries from `m`, then returns `m`.
func Sparse(m bson.M) bson.M {
	for key, val := range m {
		if len(key) == 0 {
			delete(m, key)
		} else if key != "_id" {
			switch spec := val.(type) {

			case string:
				if len(spec) == 0 {
					delete(m, key)
				}
			case complex64:
				if spec == 0 {
					delete(m, key)
				}
			case complex128:
				if spec == 0 {
					delete(m, key)
				}
			case float32:
				if spec == 0 {
					delete(m, key)
				}
			case float64:
				if spec == 0 {
					delete(m, key)
				}
			case int:
				if spec == 0 {
					delete(m, key)
				}
			case int8:
				if spec == 0 {
					delete(m, key)
				}
			case int16:
				if spec == 0 {
					delete(m, key)
				}
			case int32:
				if spec == 0 {
					delete(m, key)
				}
			case int64:
				if spec == 0 {
					delete(m, key)
				}
			case uint:
				if spec == 0 {
					delete(m, key)
				}
			case uint8:
				if spec == 0 {
					delete(m, key)
				}
			case uint16:
				if spec == 0 {
					delete(m, key)
				}
			case uint32:
				if spec == 0 {
					delete(m, key)
				}
			case uint64:
				if spec == 0 {
					delete(m, key)
				}
			case bool:
				if !spec {
					delete(m, key)
				}

			case []string:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []complex64:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []complex128:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []float32:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []float64:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []int:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []int8:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []int16:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []int32:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []int64:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []uint:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []uint8:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []uint16:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []uint32:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []uint64:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []bool:
				if len(spec) == 0 {
					delete(m, key)
				}
			case []interface{}:
				if len(spec) == 0 {
					delete(m, key)
				}

			default:
				if val == nil {
					delete(m, key)
				}
			}
		}
	}
	return m
}
