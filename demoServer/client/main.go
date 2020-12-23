package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	proto "ylMic/common/proto/greeter"
)

/*
set MICRO_REGISTRY=etcd
set MICRO_REGISTRY_ADDRESS=192.168.109.131:12379
set MICRO_API_NAMESPACE=api.tutor.com
micro web
*/

func main() {

	newRegistry := etcdv3.NewRegistry(
		registry.Addrs("111.231.255.29:12379"),
	)

	// 定义服务，可以传入其它可选参数
	service := micro.NewService(
		micro.Name("greeterer"),
		micro.Registry(newRegistry),
		micro.Selector(selector.DefaultSelector), //使用默认的负载均衡策略
	)
	fmt.Println("走到这里1")

	service.Init()
	fmt.Println("走到这里2")

	//set MICRO_REGISTRY=eureka
	//set MICRO_API_NAMESPACE=api.tutor.com
	// 创建新的客户端
	fmt.Println(service.Client().Options().Registry.Options().Addrs)
	fmt.Println(service.Client().Options().CallOptions.Address)
	greeter := proto.NewGreeterService("greeter", service.Client())
	fmt.Println("走到这里3")

	// 调用greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	fmt.Println("走到这里4")

	if err != nil {
		fmt.Println(err)
	}

	// 打印响应请求
	fmt.Println("走到这里5")
	fmt.Println(rsp.Greeting)
}
