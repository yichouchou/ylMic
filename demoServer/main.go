package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	red "github.com/gomodule/redigo/redis"
	"ylMic/demoServer/model"

	"ylMic/common/pojo"
	//"ylMic/controller"
	"ylMic/common/tool"
)

/*replace github.com/afex/hystrix-go/hystrix v0.0.0 => /path/to/go/src/github.com/afex/hystrix-go/hystrix*/

func main() {
	//验证验证码的代码，可注释
	//tool.Main()

	//加載json配置文件_无须返回值_
	_, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
		tool.Logger().Error("配置文件加载异常")
	}

	//初始化redis連接池_从config实例对象获取属性_可插拔的，选择初始化
	tool.InitRedisPool()
	tool.Logger().Error("redis加载完成")

	//初始化数据库连接池_从config实例对象获取属性_可插拔的，选择初始化_无须返回值，实例化对象之后，返回一个client连接对象，所有操作都用这个连接对象
	tool.InitDbPool()
	tool.Logger().Error("数据库连接池配置完成")

	client := tool.GetDbClient()

	//query, _ := client.Query("select * from orders limit 20")
	//
	//for _,v:=range query{
	//	fmt.Println(string(v["ordertime"]))
	//}

	orderList := make([]model.Order, 0)

	client.SQL("select * from orders limit 20").Limit(10, pojo.PageSize).Find(&orderList)
	for _, v := range orderList {
		fmt.Println(v)
	}
	//orderResult := new(pojo.PageBean)
	//orderResult.Result = orderList
	//orderResult.PageTotal=len(orderList)
	//fmt.Println(orderResult)

	//操作redis
	tool.Exec("set", "hello", "world")
	result, err := tool.Exec("get", "hello")
	if err != nil {
		fmt.Println(err.Error())
	}
	str, _ := red.String(result, err)
	fmt.Println(str)

	////使用默认gin配置
	//app := gin.Default()
	//
	////开启日志
	//app.Use(tool.LoggerToFile())
	//
	////设置全局跨域访问
	//app.Use(tool.Cors())
	//
	////集成session
	//tool.InitSession(app)
	//
	////注册路由
	//registerRouter(app)
	//
	//app.Run(cfg.AppHost + ":" + cfg.AppPort)

	// 在主函数中定义路由规则
	router := gin.Default()
	v1 := router.Group("/apis/v1/")
	{
		v1.GET("/login", tool.Login)
	}

	// secure v1
	sv1 := router.Group("/apis/v1/auth/")
	// 加载自定义的JWTAuth()中间件,在整个sv1的路由组中都生效
	sv1.Use(tool.JWTAuth())
	{
		sv1.GET("/time", tool.GetDataByTime)
	}
	router.Run(":8081")
}

//路由设置
func registerRouter(router *gin.Engine) {
	//new(controller.HelloController).Router(router)
	//new(controller.MemberController).Router(router)
	//new(controller.FoodCategoryController).Router(router)
	//new(controller.ShopController).Router(router)
	//new(controller.GoodController).Router(router)
}
