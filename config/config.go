package config

import "strconv"

var (
	ServiceName = "go.micro.service.product"
	ClientName  = "go.micro.client.product"
	Version     = "latest"
	host        = "127.0.0.1"
	port        = 8090
	Address     = host + ":" + strconv.Itoa(port)

	TracerServerName = "server"         //jaeger service name（服务器的tracer name）
	TracerClientName = "client"         //jaeger service name（客户端的tracer name）
	TracerAddr       = "localhost:6831" //jaeger agent
)
