package main

import (
	"context"
	"demo/config"
	product "demo/proto"
	"demo/tracer"
	"fmt"
	"log"

	opentracingWrapper "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	opentracing "github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
)

func main() {
	//1. 初始化链路追踪(client可以不用链路追踪)
	tracer, closer, tracerErr := tracer.NewTracer(config.TracerClientName, config.TracerAddr)
	if tracerErr != nil {
		panic(tracerErr)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer) // 设置全局链路追踪

	// 创建新服务
	service := micro.NewService(
		micro.Name(config.ClientName),
		micro.WrapClient(opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer())),
	)
	service.Init()

	// 创建 ProductService 客户端
	// 第一个参数为目标服务的名字，即客户端将调用这个名字的服务。
	productService := product.NewProductService(config.ServiceName, service.Client())

	// 发送请求
	req := &product.GetProductRequest{Id: 1}
	resp, err := productService.GetProduct(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}

	fmt.Printf("Product: %v\n", resp.Product)
}
