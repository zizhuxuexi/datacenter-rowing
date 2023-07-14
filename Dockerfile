# 基础镜像
FROM golang:1.20

# 设置工作目录
WORKDIR /app

COPY ./config/config.yaml ./config/
COPY ./training ./

EXPOSE 10002

# 设置容器启动命令

CMD ["./training"]
