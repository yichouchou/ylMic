package pojo

//type sqlCache struct {
//	SqlCacheMap map[string]interface{}
//}

var sqlCacheMap map[string]interface{}

func GetSqlCacheMap() map[string]interface{} {
	return sqlCacheMap
}
