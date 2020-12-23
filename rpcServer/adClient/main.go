package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	proto "ylMic/common/proto/ad"
)

func main() {

	registry := etcdv3.NewRegistry(
		registry.Addrs("111.231.255.29:12379"),
	)

	// 定义服务，可以传入其它可选参数
	service := micro.NewService(
		micro.Name("Adviertisement"),
		micro.Registry(registry),
	)
	service.Init()

	// 创建新的客户端Adviertisement
	ad := proto.NewAdviertisementService("Adviertisement", service.Client())

	adver := new(proto.Adviertisement)
	adver.Id = 1
	adver.Type = 2
	adver.Tel = "18452424340"
	adver.Title = "testTital"
	adver.Name = "testName"

	p := new(proto.PageInfo)
	p.PageNum = 1
	p.PageSize = 10
	p.Total = 100
	// 调用greeter
	i := new(proto.QueryByExampleRequest)
	i.Adviertisement = adver
	i.PageInfo = p

	rsp, err := ad.QueryAdviertisement(context.TODO(), i)
	if err != nil {
		fmt.Println(err)
	}

	// 打印响应请求
	fmt.Println(rsp.Adviertisement)
	fmt.Println(rsp.PageInfo)
}
