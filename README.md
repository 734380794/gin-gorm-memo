# Memo-api 备忘录

**此项目使用`Gin`+`Gorm` ，基于`RESTful API`实现的一个备忘录**。

**此项目适合`web开发`新手入门学习，掌握Gin+Gorm的使用**
# 详细教程

## 项目运行
### 手动执行
**本项目使用`Go Mod`管理依赖。**

将所有环境拉取启动

**下载依赖**

```shell
go mod tidy
```

**运行**

```shell
go run main.go
```

## 项目主要功能介绍

- 用户注册登录 ( jwt-go鉴权 )
- 新增/删除/修改/查询 备忘录
- 存储每条备忘录的浏览次数
- 分页功能


## 项目主要依赖：

**Golang V1.21**

- Gin
- Gorm
- mysql
- redis
- ini
- jwt-go
- logrus
- go-swagger

## 项目结构

```shell

├── api
├── conf
├── docs
├── middleware
├── pkg
│  ├── e
│  └── util
├── routes
├── model
├── routes
├── service
```

- api : 用于定义接口函数,也就是controller层
- cmd : 程序启动
- conf : 用于存储配置文件
- middleware : 应用中间件
- pkg/e : 封装错误码
- pkg/util : 工具函数
- model: 定义所有持久层数据库表结构的model层
- routes : 路由逻辑处理
- service : 接口函数的实现
- types : 放置所有的定义的结构体

## 配置文件
配置文件在conf/config.ini

**conf/config.ini**
```ini
# debug开发模式,release生产模式
[service]
AppMode = debug
HttpPort = :3000
# 运行端口号 3000端口

[mysql]
Db = mysql
DbHost =
# mysql ip地址
DbPort = 
# mysql 端口号
DbUser = 
# mysql 用户名
DbPassWord = 
# mysql 密码
DbName = 
# mysql 名字
```