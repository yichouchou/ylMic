package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"

	micro "github.com/micro/go-micro/v2"
	proto "ylMic/common/proto/greeter"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	eurekaRegistry := etcdv3.NewRegistry(
		registry.Addrs("111.231.255.29:12379"),
	)
	//credentials := eureka.OAuth2ClientCredentials("hehe", "haha", "127.0.0.1:8888")

	fmt.Println(eurekaRegistry.Options())
	// 创建新的服务，这里可以传入其它选项。
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(eurekaRegistry),
	)
	// 初始化方法会解析命令行标识
	service.Init()

	// 注册处理器
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
