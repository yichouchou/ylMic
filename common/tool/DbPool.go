package tool

import (
	_ "github.com/go-sql-driver/mysql"
	//"github.com/go-xorm/xorm"
	"github.com/xormplus/xorm"
	"sync"
)

type DBClient struct {
	Engine *xorm.Engine
}

//var dbClient =new(DBClient)

var dbClient *xorm.Engine = nil

//var dbClient *xorm.Engine = nil

func GetDbClient() *xorm.Engine {
	return dbClient
}

var once sync.Once

func InitDbPool() {
	once.Do(func() {
		config := GetConfig().Database
		dataSource := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DbName + "?charset=utf8"
		//dbClient, _ = xorm.NewEngine(config.Driver, "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
		dbClient, _ = xorm.NewEngine(config.Driver, dataSource)
		dbClient.SetMaxIdleConns(config.MaxIdleConns)
		dbClient.SetMaxOpenConns(config.MaxOpenConns)
		//禁用全局缓存
		dbClient.SetDisableGlobalCache(true)
		dbClient.Cascade(false)
		dbClient.ShowSQL(true)
		//dbClient.NewSession 创建一个事务
	})
}
