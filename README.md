# 本demo实现链路追踪

## docker拉取jaeger
```shell
docker pull jaegertracing/all-in-one

docker run -d --name jaeger -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one
```

## 注意
> 本demo没有将服务注册到consul，而是用go-micro默认的mdns


## 查看链路追踪
运行server后，打开输入 http://127.0.0.1:16686/