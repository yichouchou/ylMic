package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"ylMic/common/tool/wrappers"

	"github.com/gin-gonic/gin"
	//"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	limiter "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	proto "ylMic/common/proto/greeter"
	//"go-micro/Services"
)

func main() {
	const QPS = 100
	newRegistry := etcdv3.NewRegistry(
		registry.Addrs("111.231.255.29:12379"),
	)

	// 定义服务，可以传入其它可选参数
	service := micro.NewService(
		micro.Name("greeterer"),
		micro.Registry(newRegistry),
		micro.Selector(selector.DefaultSelector),          //负载均衡使用默认的
		micro.WrapClient(wrappers.NewProdsWrapper),        //使用采用熔断器的warpper，
		micro.WrapClient(wrappers.NewProdsWrapper),        //可以再添加一个来做日志搜集，比如这一行的NewProdsWrapper改成newlogwarpper
		micro.WrapHandler(limiter.NewHandlerWrapper(QPS)), //限流器
	)
	//注册服务

	engine := gin.Default()
	engine.POST("/user/", func(context *gin.Context) {
		context.String(200, "get userinfos")
	})

	microService := web.NewService(
		web.Name("userserver"),
		//web.RegisterTTL(time.Second*30),//设置注册服务的过期时间
		//web.RegisterInterval(time.Second*20),//设置间隔多久再次注册服务
		web.Address(":9999"),
		web.Handler(engine),
		web.Registry(newRegistry),
	)
	microService.Init()
	microService.Run()

	fmt.Println("走到这里1")

	service.Init()
	fmt.Println("走到这里2")
	service.Run()
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
	//consulReg := etcdv3.NewRegistry( //新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
	//	registry.Addrs("localhost:8500"),
	//)
	////下面两局代码是注册rpcserver调用客户端
	//myService := micro.NewService(
	//	micro.Name("prodservice.client"),
	//	micro.WrapClient(wrappers.NewLogWrapper),            //在注册时只需要传入方法名即可，底层会自动给这个方法传入client
	//	micro.WrapClient(wrappers.NewProdsWrapper), //在注册时只需要传入方法名即可，底层会自动给这个方法传入client
	//)
	//
	//prodService := Services.NewProdService("prodservice", myService.Client()) //生成的这个客户端绑定consul中存储的prodservice服务，只要调用了prodservice接口就会调用我们上面注册的中间件
	//
	//
	//
	////其实下面这段代码的作用就是启动webserver的同事的时候把服务注册进去
	//httpserver := web.NewService( //go-micro很灵性的实现了注册和反注册，我们启动后直接ctrl+c退出这个server，它会自动帮我们实现反注册
	//	web.Name("httpprodservice"),                   //注册进consul服务中的service名字
	//	web.Address(":8001"),                          //注册进consul服务中的端口,也是这里我们gin的server地址
	//	web.Handler(Weblib.NewGinRouter(prodService)), //web.Handler()返回一个Option，我们直接把ginRouter穿进去，就可以和gin完美的结合
	//	web.Registry(consulReg),                       //注册到哪个服务器上的consul中
	//)
	//httpserver.Init() //加了这句就可以使用命令行的形式去设置我们一些启动的配置
	//httpserver.Run()
}
