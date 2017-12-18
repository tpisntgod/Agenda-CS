# 运行go
FROM golang:latest

# 创建服务端运行环境
WORKDIR $GOPATH/src/github.com/bilibiliChangKai/Agenda-CS/cli

# 添加文件
ADD . $GOPATH/src/github.com/bilibiliChangKai/Agenda-CS

# 将8080端口暴露出来
EXPOSE 8080

# 添加GOPATH环境后，容器运行
CMD  ../service/main