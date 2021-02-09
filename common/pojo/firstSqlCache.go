package pojo

import "sync"

var sqlCacheMap map[string]interface{}

var on sync.Once

func GetSqlCacheMap() map[string]interface{} {
	return sqlCacheMap
}
