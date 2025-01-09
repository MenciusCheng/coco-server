# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.23 AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目文件
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o coco-server ./app/server/main.go

# 使用一个轻量级的 Alpine 镜像作为最终镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 从 builder 阶段复制构建好的二进制文件
COPY --from=builder /app/coco-server .

# 复制配置文件
COPY ./conf/config/docker/config.json ./config.json

# 暴露应用程序运行的端口
EXPOSE 9390

# 运行应用程序，并指定配置文件
CMD ["./coco-server", "-c", "config.json"]
