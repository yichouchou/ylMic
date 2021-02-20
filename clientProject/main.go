package main

import (
	"github.com/gin-gonic/gin"
	"ylMic/clientProject/autoClient"
	"ylMic/clientProject/controller"
	"ylMic/common/tool/wrappers"
)

/*replace github.com/afex/hystrix-go/hystrix v0.0.0 => /path/to/go/src/github.com/afex/hystrix-go/hystrix*/

func main() {
	//验证验证码的代码，可注释
	//tool.Main()

	//加載json配置文件_无须返回值_
	//_, err := tool.ParseConfig("./config/app.json")
	//if err != nil {
	//	panic(err.Error())
	//	tool.Logger().Error("配置文件加载异常")
	//}

	router := gin.Default()
	router.Use(wrappers.Cors())
	registerRouter(router)
	v1 := router.Group("/apis/v1/")
	{
		v1.GET("/login", wrappers.Login)
	}

	// secure v1
	sv1 := router.Group("/apis/v1/auth/")
	// 加载自定义的JWTAuth()中间件,在整个sv1的路由组中都生效
	sv1.Use(wrappers.JWTAuth())
	{
		sv1.GET("/time", wrappers.GetDataByTime)
	}
	go autoClient.StartChrome()

	router.Run(":8081")
}

//路由设置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
	new(controller.FoodCategoryController).Router(router)
	new(controller.ShopController).Router(router)
	new(controller.GoodController).Router(router)
}
