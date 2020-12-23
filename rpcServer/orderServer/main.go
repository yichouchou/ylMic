package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"ylMic/common/tool"
	"ylMic/rpcServer/orderServer/handler"

	micro "github.com/micro/go-micro/v2"
	proto "ylMic/common/proto/greeter"
)

func main() {
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
	eurekaRegistry := etcdv3.NewRegistry(
		registry.Addrs("111.231.255.29:12379"),
	)

	// 创建新的服务，这里可以传入其它选项。
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(eurekaRegistry),
	)
	// 初始化方法会解析命令行标识
	service.Init()

	// 注册处理器
	proto.RegisterGreeterHandler(service.Server(), new(handler.Greeter))

	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
