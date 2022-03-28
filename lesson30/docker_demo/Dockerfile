FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/docker_demo
WORKDIR /docker_demo

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 app
RUN go build -o app .

# 分阶段构建
# 接下来创建一个小镜像
FROM scratch

# 从 builder 镜像中的把 /docker_demo/app 拷贝到当前目录
COPY --from=builder /docker_demo/app /

# 声明服务端口
EXPOSE 8888

# 需要运行的命令
ENTRYPOINT ["/app"]