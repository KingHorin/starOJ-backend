# starOJ-backend

## 环境要求

Go 1.21, 安装并启动Mysql和Redis服务

## 部署步骤

1.终端进入根目录, 输入go mod download, 等待依赖下载完毕

2.修改./config下的database.go和redis.go配置信息，与本地对应

3.终端输入go run main.go或goland直接运行，开始工作
