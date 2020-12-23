package Wrappers

//protoc --proto_path=. --go_out=. --micro_out=. proto/prod/prodModel.proto  生成proto的命令
//protoc --proto_path=. --go_out=. --micro_out=. proto/adviertisement/adviertisement.proto
import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"strconv"
	"ylMic/common/proto/greeter"
	"ylMic/common/proto/prod"
)

type ProdsWrapper struct { //官方提供的例子，创建自己的struct，嵌套go-micro的client
	client.Client
}

func defaultProds(rsp interface{}) { //将rsp中传入响应值，这里响应值是我们proto定义好的返回值
	switch t := rsp.(type) {
	case *prod.ProdListResponse:
		fmt.Println("遇到该请求的降级函数执行操作")
	case *prod.ProdModel:
		fmt.Println("遇到该请求时候的降级操作")
		t.ProdID = 123
	case *greeter.HelloResponse:
		fmt.Println("greeter的降级操作")
	default:

	}

	models := make([]*prod.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		//此处看不懂什么意义，为什么是1-5遍历，
		models = append(models, newProd(uint32(20+i), "prodname"+strconv.Itoa(20+int(i))))
	}
	result := rsp.(*greeter.HelloResponse) //类型断言为我们定义好的返回值
	result.Greeting = "models"
}

func newProd(i uint32, s string) *prod.ProdModel {
	return &prod.ProdModel{ProdID: i, ProdName: s}
}

//重写Call方法
func (this *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint() //req.Service()是服务名.Endpoint是方法,这里是ProdService.GetProdsList,这个名字并不会对结果有影响，只是这里规范定义而已，其实定义hello world也可以运行
	/*
		protoc --proto_path=. --micro_out=. --go_out=. proto/prod/prodModel.proto

		protoc --proto_path=. --micro_out=. --go_out=. prodModel.proto
		protoc --proto_path=. --micro_out=ylMic/common/protp/prod/prodModel.proto=github.com/micro/go-m
		icro/v2/api/proto:. --go_out=ylMic/common/protp/prod/prodModel.proto=github.com/micro/go-micro/
		v2/api/proto:. prodModel.proto


		   service ProdService{
		       rpc GetProdsList (ProdsRequest) returns (ProdListResponse);
		   }
	*/
	configA := hystrix.CommandConfig{
		Timeout:                5000,
		RequestVolumeThreshold: 2,    //阈值：意思是有20个请求才进行错误百分比计算
		ErrorPercentThreshold:  50,   //错误百分比 20% 的错误发生之后，就直接执行降级方法
		SleepWindow:            5000, //过多少毫秒之后重新尝试后端被降级的服务

	}
	hystrix.ConfigureCommand(cmdName, configA)
	return hystrix.Do(cmdName, func() error {
		return this.Client.Call(ctx, req, rsp) //调用rpc api接口
	}, func(e error) error { //降级函数
		defaultProds(rsp)
		return nil
	})
}

func NewProdsWrapper(c client.Client) client.Client {
	return &ProdsWrapper{c}
}
