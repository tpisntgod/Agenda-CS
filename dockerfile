# 运行go
FROM go:lastest

# 创建服务端运行环境
WORKDIR /service

# 添加文件
ADD ./service /service

# 定义GOPATH变量
ENV GOPATH /service

# 构造第一个容器
# 先创建
RUN mkdir bin pkg src
