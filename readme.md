# Gin

## Gin

安装

```cmd
go get -u github.com/gin-gonic/gin@v1.6.3
```



## 项目目录：

```shell
blog-service
├── configs
├── docs
├── global
├── internal
│   ├── dao
│   ├── middleware
│   ├── model
│   ├── routers
│   └── service
├── pkg
│   ├── app
│   ├── email
│   ├── convert
│   ├── errcode
│   ├── limiter
│   ├── logger
│   ├── setting
│   ├── tracer
│   ├── upload
│   ├── zaplogger
│   └── until
├── storage
│   ├── logs
│   └── uploads
├── scripts
└── third_party
```

- configs：配置文件。
- docs：文档集合。
- global：全局变量。
- internal：内部模块。
- dao：数据访问层（Database Access Object），所有与数据相关的操作都会在 dao 层进行，例如 MySQL、ElasticSearch 等。
- middleware：HTTP 中间件。
- model：模型层，用于存放 model 对象。
- routers：路由相关逻辑处理。
- service：项目核心业务逻辑（主要做了规定入参和校验的工作，传参给dao）。
- pkg：项目相关的模块包。
- app 相应处理、接口传参校验 分页处理
- convert 类型转换
- errcode 错误编码处理
- logger 日志处理
- setting 配置文件相关方法
- until 公共函数 MD5
- storage：项目生成的临时文件。
- scripts：各类构建，安装，分析等操作的脚本。
- third_party：第三方的资源工具，例如 Swagger UI。

## 公共组件



###### 错误码标准化

第三方开源库 viper

```cmd
go get -u github.com/spf13/viper@v1.4.0
```

###### 配置管理

###### 数据路链接

本项目中数据库相关的数据操作将使用第三方的开源库 gorm

```cmd
go get -u github.com/jinzhu/gorm@v1.9.12
```

###### 日志写入

开源库 lumberjack日志组件

```cmd
go get -u gopkg.in/natefinch/lumberjack.v2
```

###### 响应处理



## swagger

安装swagger

```cmd
go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
go get -u github.com/swaggo/gin-swagger@v1.2.0
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
```

验证：swag -v

生成: swag init

## 接口做参数校验

使用开源项目 [**go-playground/validator**](https://github.com/go-playground/validator) 作为我们的本项目的基础库，它是一个基于标签来对结构体和字段进行值验证的一个验证器

```cmd
go get -u github.com/go-playground/validator/v10
```

## 链路跟踪

 分布式链路追踪就是将一次分布式请求还原成调用链路，将一次分布式请求的调用情况集中展示，比如各个服务节点上的耗时、请求具体到达哪台机器上、每个服务节点的请求状态等等 、

Open Tracing 规范：

Trace 跟踪：

 ![img](https://pics2.baidu.com/feed/7af40ad162d9f2d3fd977039b26b8a1a6127ccf6.png?token=30fa37fc2fd2aef6c3d0274b2b3fc173) 

Span 跨度

 ![img](https://pics2.baidu.com/feed/e61190ef76c6a7ef3b10e682e17daf58f2de66f4.png?token=dc61246197557bf9d8bc0bd29a27fcde) 



SpanContext 跨度上下文：

 https://baijiahao.baidu.com/s?id=1708807913437543359&wfr=spider&for=pc 



#### Jaeger 的使用：

**安装**

``` cmd
docker pull jaegertracing/all-in-one:latest
```

**linux/amd64 启动**

 ``` cmd
docker run -d --name jaeger  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411  -p 5775:5775/udp  -p 6831:6831/udp  -p 6832:6832/udp  -p 5778:5778  -p 16686:16686  -p 14250:14250  -p 14268:14268  -p 14269:14269  -p 9411:9411  jaegertracing/all-in-one:latest
 ```

安装参照: https://zhuanlan.zhihu.com/p/524695029 

Windows操作docker 命令

``` dockerfile
docker version

docker images

docker ps -a      //命令查看所有docker容器

docker start jaeger  //启动容器 

docker restart 容器ID
```

参照： https://www.csdn.net/tags/Mtzakg4sNjY4ODAtYmxvZwO0O0OO0O0O.html 



在浏览器中访问webUI： [http://localhost:16686](http://localhost:16686/) 



安装第三方插件

``` cmd
go get github.com/opentracing/opentracing-go v1.2.0
go get github.com/uber/jaeger-client-go v2.22.1
```





### 特性

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- Cron
- Redis