FROM golang:1.17.3-alpine3.15 as builder

# 安装构建的必要工具和时区
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
  apk update && \
  apk upgrade && \
  apk add ca-certificates gcc g++ && update-ca-certificates && \
  apk add --update tzdata && \
  apk add upx && \
  rm -rf /var/cache/apk/*

# 构建时需要的环境变量
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOPROXY="https://goproxy.cn,direct"
ENV TZ=Asia/Shanghai

# 创建一个目录，所有项目都将在该目录下进行构建
WORKDIR /builder

# 将本地项目复制到/builder目录中
COPY . .

# 下载项目依赖
RUN go mod download
# 开始压缩构建 && 使用 upx 再次压缩
RUN go build -ldflags "-s -w" -o app_serv && upx -9 app_serv

FROM alpine as runner

# 设置时区
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
  apk update && \
  apk upgrade && \
  apk add --update tzdata && \
  rm -rf /var/cache/apk/*

# 设置时区的环境变量
ENV TZ=Asia/Shanghai

# 运行目录
WORKDIR /golang

# 从构建容器中复制构建好的二进制文件到当前容器的/golang目录中
COPY --from=builder /builder/app_serv .
# 从构建容器中复制配置文件到当前容器的/golang目录中 【这里是为了兼容本地配置启动的模式】
COPY --from=builder /builder/app.yaml .

# 运行
CMD ["/golang/app_serv"]
