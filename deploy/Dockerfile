# 多阶段构建，
FROM golang:alpine AS builder

# 设置环境变量
ENV ENV GO111MODULE=on \
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64 \
        GOPROXY="https://goproxy.cn,direct"

# 移动到工作目录 /build
WORKDIR /build

# 将代码赋值到文件中
COPY . .

# 将代码编译成二进制文件pandora
RUN go build -o pandora .

# alpine 镜像
FROM alpine
# 设置时区
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata && \
    rm -rf /var/cache/apk/* \

WORKDIR /app

COPY ./config .

RUN mkdir logs

# 从builder镜像把pandora二进制复制过来
COPY --from=builder /build/pandora .

# 暴露应用的端口
EXPOSE 5001
CMD ["./pandora"]
