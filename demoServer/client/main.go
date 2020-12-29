package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	microhttp "github.com/micro/go-plugins/client/http/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"net/http"
	"time"
)

/*
set MICRO_REGISTRY=etcd
set MICRO_REGISTRY_ADDRESS=192.168.109.131:12379
set MICRO_API_NAMESPACE=api.tutor.com
micro web
*/

func InitRouters() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.POST("/orders/", func(context *gin.Context) {
		context.String(200, "get orderinfos")
	})

	return ginRouter
}

func main() {

	//newRegistry := etcdv3.NewRegistry(
	//	registry.Addrs("111.231.255.29:12379"),
	//)
	//
	//// 定义服务，可以传入其它可选参数
	//service := micro.NewService(
	//	micro.Name("greeterer"),
	//	micro.Registry(newRegistry),
	//	micro.Selector(selector.DefaultSelector), //使用默认的负载均衡策略
	//)
	//fmt.Println("走到这里1")
	//
	//service.Init()
	//fmt.Println("走到这里2")
	//
	////set MICRO_REGISTRY=eureka
	////set MICRO_API_NAMESPACE=api.tutor.com
	//// 创建新的客户端
	//fmt.Println(service.Client().Options().Registry.Options().Addrs)
	//fmt.Println(service.Client().Options().CallOptions.Address)
	//greeter := proto.NewGreeterService("greeter", service.Client())
	//fmt.Println("走到这里3")
	//
	//// 调用greeter
	//rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	//fmt.Println("走到这里4")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// 打印响应请求
	//fmt.Println("走到这里5")
	//fmt.Println(rsp)
	//
	//microselector := selector.NewSelector(
	//	selector.Registry(newRegistry),            //传入注册
	//	selector.SetStrategy(selector.RoundRobin), //指定查询机制
	//)
	//microClient := microhttp.NewClient(
	//	client.Selector(microselector),
	//	client.ContentType("application/json"))
	//microClient.Init()
	//fmt.Println(microClient.Options().CallOptions.Address)
	//fmt.Println(microClient.Options().Registry.Options().Addrs)
	//
	//req := microClient.NewRequest("orderServer", "/orders", map[string]string{})
	//var resp map[string]interface{}
	//
	//err = microClient.Call(context.Background(), req, &resp)
	//if err == nil {
	//	fmt.Println(resp)
	//} else {
	//	fmt.Println(err)
	//}

	//初始化路由
	//新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
	consulReg := etcdv3.NewRegistry(
		registry.Addrs("111.231.255.29:12379"),
	)

	microselector := selector.NewSelector(
		selector.Registry(consulReg),              //传入consul注册
		selector.SetStrategy(selector.RoundRobin), //指定查询机制
	)
	microClient := microhttp.NewClient(
		//client.Registry(consulReg),
		client.Selector(microselector),
		client.ContentType("application/json"))

	address := GetServiceAddr("userserver", consulReg)

	if len(address) <= 0 {
		fmt.Println("hostAddress is null")
	} else {
		url := "http://" + address + "/users/"
		response, _ := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer([]byte("haha")))

		fmt.Println(response)
	}

	req := microClient.NewRequest("userserver", "/users/", map[string]string{})

	err2 := microClient.Init()
	fmt.Println(err2)
	fmt.Println(microClient.Options().Selector.Options().Registry.Options().Addrs)
	var resp map[string]interface{}

	method := req.Method()
	service := req.Service()
	fmt.Println(service)
	endpoint := req.Endpoint()
	fmt.Println(endpoint)
	contentType := req.ContentType()
	fmt.Println(contentType)
	fmt.Println(method)
	err := microClient.Call(context.Background(), req, &resp)
	if err == nil {
		fmt.Println(resp)
	} else {
		fmt.Println(err)
	}

}

func GetServiceAddr(serviceName string, consulReg registry.Registry) (address string) {
	var retryCount int
	for {
		servers, err := consulReg.GetService(serviceName)
		if err != nil {
			fmt.Println(err.Error())
		}
		var services []*registry.Service
		for _, value := range servers {
			fmt.Println(value.Name, ":", value.Version)
			services = append(services, value)
		}
		next := selector.RoundRobin(services)
		if node, err := next(); err == nil {
			address = node.Address
		}
		if len(address) > 0 {
			return
		}
		//重试次数++
		retryCount++
		time.Sleep(time.Second * 1)
		//重试5次为获取返回空
		if retryCount >= 5 {
			return
		}
	}
}
