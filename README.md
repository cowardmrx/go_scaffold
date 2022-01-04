#go_scaffold

### 简介
基于gin、gorm、urfave-cli的一个快速开发手脚架

### 目录介绍
```go
go_scaffold
│   app.example.yaml                                    --本地配置示例
│   Dockerfile                                          --镜像构建文件
│   go.mod                                              --项目依赖
│   go.sum
│   main.go                                             --系统入口文件
│   README.md                                           --文档
├───app                                                 --应用核心
│   ├───cron                                            --定时任务服务
│   │   │   kernel.go                                   --定时任务核心
│   │   │
│   │   └───cron_job                                    --定时任务代码区
│   │           job_xxx.go                              --定时任务示例
│   │
│   ├───http                                            --web服务【http服务】
│   │   │   kernel.go                                   --http服务核心
│   │   │
│   │   ├───controller                                  --控制器层面，这里仅针对http服务
│   │   │   └───user_controller                         --示例
│   │   │           user_controller.go
│   │   │
│   │   ├───core                                        --核心【一些http必要包或者错误处理可以放在这里】
│   │   │   └───response                                --http 响应返回
│   │   │           http_response.go
│   │   │
│   │   ├───middleware                                  --http中间件【middleware】
│   │   ├───request                                     --http请求参数校验
│   │   │   │   request.go                              --基本校验结构体，每个校验结构体都必须嵌套该基础结构体
│   │   │   │
│   │   │   └───user_request                            --示例
│   │   │           user_add_request.go
│   │   │
│   │   ├───resources                                   --resouces返回，资源返回自定义封装
│   │   └───router                                      --http路由模块
│   │       │   router.go                               --路由的唯一入口
│   │       │
│   │       └───route                                   --[示例代码]模块之一
│   │               user.go
│   │
│   ├───model                                           --数据模型
│   │       model.go                                    --基本数据模型，主要维护created_at和updated_at字段，可根据实际情况自行更改
│   │       user_model.go
│   │
│   ├───modelfilter                                     --模型过滤器
│   ├───queue                                           --队列【未开发】
│   ├───repository                                      --数据仓库
│   │       repository.go                               --基本数据仓库
│   │       user_repository.go
│   │
│   └───service                                         --业务逻辑层
│       └───user_service                                --示例代码
│               user_service.go
│
├───cli                                                 --cli整合系统的服务分层
│   │   cli.go                                          --所有的服务均在此处设置cli命令
│   │
│   ├───cmd                                             --一些单独的cmd命令，或者公用的cmd命令
│   │       gorm_cmd.go
│   │
│   └───serve                                           --不同的服务
│           cron.go                                     --定时任务服务的入口
│           http_serve.go                               --http服务的入口
│
├───config                                              --项目配置
│       app.go                                          --项目基本配置，服务名等
│       config.go                                       --所有项目的配置聚合
│       database.go                                     --数据库的配置
│       http.go                                         --http服务的配置
│       nacos.go                                        --nacos的配置
│       redis.go                                        --nodb redis的配置
│
├───global                                              --全局变量或配置等
│       cache.go                                        --缓存初始化
│       config.go                                       --从配置中心或本地配置文件分发或加载配置到config模块中的各个配置当中去
│       database.go                                     --初始化数据库
│       global.go                                       --一些全局变量
│       redis.go                                        --初始化redis
│       validator.go                                    --校验器
│
├───initialize
│       init.go                                         --系统初始化
│
├───logs                                                --日志存储文件夹
└───pkg                                                 --第三方或者自己封装的其他包存放处
├───utils                                               --工具包
│       env.go                                          --环境变量读取
│       ips.go                                          --获取ip地址
│
└───validator_rules                                     --验证器的自定义校验规则
        rules.go
```

## 使用
```go
1.go clone git@github.com:cowardmrx/go_scaffold.git
2.根目录运行go mod download or go get
```
## 运行
```go
1.建议使用goland运行项目，方便配置必须的环境变量
2.使用本地配置文件方式  
  go run main.go --cfs=LOCAL app:serve  // 启动http服务
  go run main.go --cfs=LOCAL cron // 启动定时任务服务
  如果使用本地配置文件启动请参考app.example.yaml新建一份app.yaml文件
3.使用NACOS配置读取
  go run main.go --cfs=NACOS app:serve // 启动http服务
  go run main.go --cfs=NACOS cron // 启动定时任务服务
  如果采用NACOS启动项目，以下的环境变量必须具备：
    NACOS_HOST=192.168.0.1  // nacos的地址
    NACOS_PORT=8848         // nacos的访问端口一般为8848
    NAMESPACE_ID=43a4e90b-6d01-45f4-b195-06a052478cd0 // nacos配置所处的命名空间
    CONFIG_NAME=app.yaml    // nacos存放配置的配置名称
    APP_NAME=go_scaffold    // 项目名称
    PORT=8989               // http服务启动端口，如果启动的是cron服务可以不用填写该配置项和APP_NAME主要用于naco服务注册使用
```

## 构建&运行
```go
1.建议使用Docker进行构建，方便环境隔离以及容器隔离，可更方便管理
2.docker build -t go_scaffold:v1 .  // Dockerfile文件存在项目根目录可自行更改
3.docker-compose up -d  // docker-compose.yaml文件存在项目根目录可自行更改
```