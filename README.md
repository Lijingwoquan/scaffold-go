# Go Web 项目手脚架

这是一个基于 Go 语言的 Web 项目手脚架,集成了常用的库和功能,帮助您快速启动新项目。

## 功能特性

- 使用 Zap 进行日志记录
- 使用 Viper 进行配置管理
- 使用 SQLx 进行数据库操作
- Redis 集成
- 缓存层实现
- 定时任务 (Ticker) 功能
- JWT 认证
- 项目重命名工具 (replace.exe)

## 快速开始

1. 克隆项目:
````shell
git clone https://github.com/Lijingwoquan/scaffold.git
cd scaffold
````

2. 重命名项目:
   使用提供的 `replace.exe` 工具来快速修改项目名称:
```shell
replace.exe ./ your-new-project-name
```
这将把项目中的 "scaffold" 替换为 "your-new-project-name"。

3. 配置项目:
   编辑 `config/config.yaml` 文件,根据您的需求修改配置。


## 项目结构
   ```
├── cache/
├── config/
│   └── config.yaml
├── controller/
├── dao/
│   ├── mysql/
│   └── redis/
├── log/
├── logic/
├── middlewares/
├── models/
├── pkg/
│   ├── jwt/
│   └── snowflake/
│   └── ticker/
├── routers
│   └── router.go
├── setting
│   └── setting.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── init.sql
├── README.md
└── replace.exe
```

## 配置

项目使用 Viper 进行配置管理。主要配置文件位于 `config/config.yaml`。您可以在此文件中设置数据库连接、Redis 连接、JWT 过期时间等。

## 日志

项目使用 Zap 日志库。日志配置可以在 `config/config.yaml` 中进行调整。

## 数据库

项目使用 SQLx 进行数据库操作。数据库连接信息在 `config/config.yaml` 中配置。

## Redis

Redis 用于缓存和会话管理。Redis 连接信息同样在 `config/config.yaml` 中配置。

## 缓存层

项目实现了一个缓存层,用于优化数据访问性能。缓存策略可以在 `cache/` 目录中找到。

## 定时任务

项目包含一个定时任务 (Ticker) 系统,用于执行周期性任务。定时任务可以在 `pkg/ticker/` 目录中定义。

## JWT 认证

项目使用 JWT 进行用户认证。JWT 密钥和工具函数可以在 `pkg/jwt/` 目录中找到。

## Docker 支持

项目包含 Dockerfile 和 docker-compose.yml,支持使用 Docker 进行部署。

## 贡献

欢迎提交 Pull Requests 来改进这个项目。