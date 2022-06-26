# 多阶段构建，
FROM golang:1.18-alpine AS builder

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

WORKDIR /data/project
RUN pwd
RUN mkdir logs

# COPY 命令是将源文件夹的里的子文件和文件夹复制到目的文件夹，因此需要复制整个文件夹时，指定目标文件夹和源文件夹一致
COPY config ./config

# 从builder镜像把pandora二进制复制过来
COPY --from=builder /build/pandora .

# 暴露应用的端口
EXPOSE 5001
CMD ["./pandora"]
