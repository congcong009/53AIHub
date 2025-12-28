# 简化版Dockerfile - 用于快速构建（假设前端已预构建）
# 如果服务器资源有限，建议先在本地构建前端，然后使用此Dockerfile

FROM golang:1.23-alpine AS builder

# 安装构建依赖
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# 设置GOPROXY为中国大陆代理
ENV GOPROXY=https://goproxy.cn,direct

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码和预构建的静态资源
COPY . .

# 安装swag并生成文档
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN if [ -f "docs/docs.go" ]; then \
        echo "docs already exist, using existing docs"; \
    else \
        echo "generating docs..."; \
        swag init -g router/main.go; \
    fi

# 构建后端应用
RUN go build -o main .

# 运行时镜像
FROM alpine:latest

# 安装ca-certificates
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制可执行文件和静态资源
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/docs ./docs

EXPOSE 3000

CMD ["./main"]