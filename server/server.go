package main

import (
	"context"
	product "demo/proto"
	"demo/tracer"
	"log"

	"demo/config"
	opentracingFn "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
)

func main() {
	// 1. 链路追踪配置
	tracer, closer, tracerErr := tracer.NewTracer(config.TracerServerName, config.TracerAddr)
	if tracerErr != nil {
		panic(tracerErr)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer) // 设置全局链路追踪

	// 2. 创建服务(这里需要实现 Service 接口)
	productService := new(ProductService)

	// 3. 创建服务和初始化
	srv := micro.NewService(
		micro.Name(config.ServiceName), // 设置服务名称
		micro.Version(config.Version),  // 设置服务版本
		micro.Address(config.Address),  // 设置服务地址
		micro.WrapHandler(opentracingFn.NewHandlerWrapper(opentracing.GlobalTracer())), // 绑定链路追踪
	)
	srv.Init()

	// 4. 注册 handler
	if err := product.RegisterProductServiceHandler(srv.Server(), productService); err != nil {
		panic(err)
	}
	// 5. 运行服务
	if runErr := srv.Run(); runErr != nil {
		log.Fatal(runErr)
	}
}

// ProductService 实现简单的 RPC 服务
type ProductService struct {
}

func (p *ProductService) GetProduct(ctx context.Context, req *product.GetProductRequest, rsp *product.GetProductResponse) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetProduct") //创建一个新的span
	defer span.Finish()
	//模拟其他操作
	outherFunc(ctx)
	//模拟查询mysql
	product, err := findProduct(ctx, req.Id)
	if err != nil {
		return err
	}
	rsp.Product = product
	return nil
}

// outherFunc 模拟其他操作
func outherFunc(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "outherFunc") //创建一个新的span
	defer span.Finish()
}

// findProduct 模拟查询数据库
func findProduct(ctx context.Context, id int64) (*product.Product, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "findProduct") //创建一个新的span
	defer span.Finish()

	return &product.Product{
		Id:   id,
		Name: "product",
	}, nil
}
