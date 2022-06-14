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
│   ├── convert
│   ├── errcode
│   ├── logger
│   ├── setting
│   └── until
├── storage
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