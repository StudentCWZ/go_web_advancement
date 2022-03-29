FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/bluebell
WORKDIR /bluebell

# 复制项目中的 go.mod 和 go.sum 文件并下载依赖信息
ADD go.mod .
ADD go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 app
RUN go build -o bluebell_app .

# 分阶段构建
# 接下来创建一个小镜像
FROM debian:stretch-slim

# 单个 dockerfile 构建不需要 copy wait-for.sh
COPY ./wait-for.sh /
COPY ./templates /templates
COPY ./static /static
COPY ./conf /conf

# 从 builder 镜像中的把 /bluebell/bluebell_app 拷贝到当前目录
COPY --from=builder /bluebell/bluebell_app /

# 声明服务端口
EXPOSE 9090

# 利用 docker compose + wait-for.sh 进行构建，如果只用 dockerfile，该操作可以注释
RUN set -eux; \
    apt-get update; \
    apt-get upgrade -y; \
    apt-get install -y --no-install-recommends netcat; \
    chmod 755 wait-for.sh

# 单个 dockerfile 需要运行的命令；如果利用 docker compose + wait-for.sh 该行可以注释
# ENTRYPOINT ["/bluebell_app"]

# 镜像构建命令(终端运行)
# docker build . -t docker_bluebell

# 启动容器的命令(联合启动容器命令-终端运行): 此时需要将配置文件中 mysql 和 redis 中的 host 分别改为 mysql 和 redis
# docker run --name docker_bluebell --link=mysql:mysql --link=redis:redis -p 9090:9090 docker_bluebell

# 该 dockerfile 只有在 bluebell 是一个根目录生效

