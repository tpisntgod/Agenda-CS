# 运行go
FROM go:lastest

# 创建服务端运行环境
WORKDIR /service

# 添加文件
ADD ./service/main /service/main

# 定义GOPATH变量
#ENV GOPATH /service

# 构造第一个容器
# 1.设置go环境
# 2.安装godep包
RUN export GO15VENDOREXPERIMENT="1"
 && go get github.com/tools/godep

# 将8080端口暴露出来
EXPOSE 8080

# 容器运行
CMD ["./main"]
