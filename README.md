# layout-go-zero

Golang服务开发架构

## 目录结构

``` lua
|-- layout-go-zero
    |-- app                     #服务中心
    |   |-- user_http           #http服务
    |       |-- controller      #控制器
    |       |-- docs            #swagger文档
    |       |-- request         #表单验证
    |       |-- route           #路由
    |   |-- user_rpc            #rpc服务
    |-- build                   #编译后的目录
    |-- cmd                     #命令
    |   |-- make                #代码生成命令
    |   |-- migrate             #数据迁移命令
    |   |-- seed                #数据填充命令
    |   |-- service             #业务层命令
    |-- config                  #配置
    |-- cron                    #定时服务
    |-- logs                    #日志文件
    |-- orm                     #数据相关
    |   |-- gen                 #gen生成
    |   |-- factory             #数据工厂
    |   |-- migration           #迁移文件
    |   |-- model               #model模型
    |   |-- seeder              #填充文件
    |-- pkg                     #自定义package
    |   |-- console             #打印
    |   |-- helper              #辅助函数
    |   |-- http                #http
    |   |-- migrate             #数据迁移
    |   |-- queue               #队列
    |   |-- redis               #redis
    |   |-- rpc                 #rpc
    |   |-- seed                #数据填充
    |-- queue                   #队列
    |-- service                 #业务抽象层

```

## 路由规则

| 请求方法   | API 示例     | 说明   |
|--------|------------|------|
| GET    | /user      | 获取列表 |
| GET    | /user/{id} | 获取详情 |
| POST   | /user      | 新增   |
| PUT    | /user/{id} | 修改   |
| DELETE | /user/{id} | 删除   |

## 所有命令

启动web服务：

```
$ go run ./app/user_http/main.go
```

启动定时服务：

```
$ go run ./cron/main.go
```

启动队列服务：

```
$ go run ./queue/main.go
```

生成web文档：

```
$ swag init --parseDependency --parseInternal -g ./app/user_http/main.go -o ./app/user_http/docs
```

cmd命令：

```
$ go run ./inlet/cmd -h
Default will run "serve" command, you can use "-h" flag to see all subcommands

Usage:
   [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  make        Generate file and code
  migrate     Run database migration
  seed        Insert fake data to the database
  service     业务层cmd

Flags:
  -h, --help          help for this command
  -p, --path string   example: --path=./config/setting.yaml (default "./config/setting.yaml")

Use " [command] --help" for more information about a command.
```

make 命令：

```
$ go run ./cmd/main.go make -h
Generate file and code

Usage:
   make [command]

Available Commands:
  cmd_service Create a command, should be snake_case, example: make cmd_service user user
  controller  Create api controller，example: make controller user_http user
  factory     Create model's factory file, example: make factory user
  gen         Generate file and code, example: make gen
  migration   Create a migration file, example: make migration user
  model       Crate model file, example: make model user
  request     Create request file, example make request user_http user
  route       Crate route file, example: make route user_http user
  seeder      Create seeder file, example: make seeder user

Flags:
  -h, --help   help for make

Use " make [command] --help" for more information about a command.
```

migrate 命令：

```
$ go run ./cmd/main.go migrate -h
Run database migration

Usage:
   migrate [command]

Available Commands:
  down        Reverse the up command
  fresh       Drop all tables and re-run all migrations
  refresh     Reset and re-run all migrations
  reset       Rollback all orm migrations
  up          Run unMigrated migrations

Flags:
  -h, --help   help for migrate

Use " migrate [command] --help" for more information about a command.
```

seed 命令：

```
$ go run ./cmd/main.go seed -h
Insert fake data to the orm

Usage:
   seed [flags]

Flags:
  -h, --help   help for seed
```
