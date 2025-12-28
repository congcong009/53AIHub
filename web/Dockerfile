# 使用Node.js官方镜像作为基础镜像（升级到20版本以满足依赖要求）
FROM node:20-alpine

# 为了在80端口运行，需要使用root用户（开发环境）
USER root

# 使用国内镜像源加速（阿里云镜像）
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 设置工作目录
WORKDIR /app

# 安装必要的系统依赖（用于某些npm包可能需要编译）
# build-base 包含 make, g++, gcc 等编译工具
# libx11-dev 和 libxtst-dev 是 robotjs 等原生模块需要的 X11 开发库
RUN apk add --no-cache \
    python3 \
    build-base \
    git \
    bash \
    libx11-dev \
    libxtst-dev \
    libxi-dev

# 复制package.json和package-lock.json（如果存在）到front目录
COPY front/package*.json ./front/

# 复制package.json和package-lock.json（如果存在）到console目录
COPY console/package*.json ./console/

# 安装front依赖（使用 BuildKit 缓存挂载，减少磁盘占用）
# 需要在构建时使用 DOCKER_BUILDKIT=1 环境变量
WORKDIR /app/front
RUN --mount=type=cache,target=/root/.npm \
    npm install

# 安装console依赖（使用 BuildKit 缓存挂载）
WORKDIR /app/console
RUN --mount=type=cache,target=/root/.npm \
    npm install

# 返回工作目录
WORKDIR /app

# 创建启动脚本
RUN echo '#!/bin/bash' > /app/start.sh && \
    echo 'set -e' >> /app/start.sh && \
    echo '' >> /app/start.sh && \
    echo '# 信号处理函数，确保容器可以优雅地停止' >> /app/start.sh && \
    echo 'cleanup() {' >> /app/start.sh && \
    echo '    echo ""' >> /app/start.sh && \
    echo '    echo "Shutting down servers..."' >> /app/start.sh && \
    echo '    if [ ! -z "$FRONT_PID" ]; then' >> /app/start.sh && \
    echo '        kill $FRONT_PID 2>/dev/null || true' >> /app/start.sh && \
    echo '    fi' >> /app/start.sh && \
    echo '    if [ ! -z "$CONSOLE_PID" ]; then' >> /app/start.sh && \
    echo '        kill $CONSOLE_PID 2>/dev/null || true' >> /app/start.sh && \
    echo '    fi' >> /app/start.sh && \
    echo '    exit 0' >> /app/start.sh && \
    echo '}' >> /app/start.sh && \
    echo '' >> /app/start.sh && \
    echo 'trap cleanup SIGTERM SIGINT' >> /app/start.sh && \
    echo '' >> /app/start.sh && \
    echo 'echo "Checking dependencies..."' >> /app/start.sh && \
    echo '' >> /app/start.sh && \
    echo '# 检查front依赖，如果node_modules不存在则安装' >> /app/start.sh && \
    echo 'if [ ! -d "/app/front/node_modules" ]; then' >> /app/start.sh && \
    echo '    echo "Installing front dependencies..."' >> /app/start.sh && \
    echo '    cd /app/front' >> /app/start.sh && \
    echo '    npm install' >> /app/start.sh && \
    echo 'fi' >> /app/start.sh && \
    echo '' >> /app/start.sh && \
    echo '# 检查console依赖，如果node_modules不存在则安装' >> /app/start.sh && \
    echo 'if [ ! -d "/app/console/node_modules" ]; then' >> /app/start.sh && \
    echo '    echo "Installing console dependencies..."' >> /app/start.sh && \
    echo '    cd /app/console' >> /app/start.sh && \
    echo '    npm install' >> /app/start.sh && \
    echo 'fi' >> /app/start.sh && \
    echo '' >> /app/start.sh && \
    echo 'echo "Starting front dev server..."' >> /app/start.sh && \
    echo 'cd /app/front' >> /app/start.sh && \
    echo 'npm run dev &' >> /app/start.sh && \
    echo 'FRONT_PID=$!' >> /app/start.sh && \
    echo '' >> /app/start.sh && \
    echo 'echo "Starting console dev server..."' >> /app/start.sh && \
    echo 'cd /app/console' >> /app/start.sh && \
    echo 'npm run dev &' >> /app/start.sh && \
    echo 'CONSOLE_PID=$!' >> /app/start.sh && \
    echo '' >> /app/start.sh && \
    echo 'echo "Both servers are starting..."' >> /app/start.sh && \
    echo 'echo "Front dev server PID: $FRONT_PID"' >> /app/start.sh && \
    echo 'echo "Console dev server PID: $CONSOLE_PID"' >> /app/start.sh && \
    echo 'echo ""' >> /app/start.sh && \
    echo 'echo "Front dev server: http://localhost:80 (configured in vite.common.ts)"' >> /app/start.sh && \
    echo 'echo "Console dev server: http://localhost:8002 (or 8003 if VITE_PLATFORM=km)"' >> /app/start.sh && \
    echo 'echo ""' >> /app/start.sh && \
    echo 'echo "Press Ctrl+C to stop all servers"' >> /app/start.sh && \
    echo '' >> /app/start.sh && \
    echo '# 等待所有后台进程' >> /app/start.sh && \
    echo 'wait $FRONT_PID $CONSOLE_PID' >> /app/start.sh && \
    chmod +x /app/start.sh

# 暴露端口
# front: 端口80（根据vite.common.ts中的commonServer配置）
# console: 端口8002（根据vite.config.ts配置，如果VITE_PLATFORM=km则为8003）
EXPOSE 80 8002 8003

# 设置启动命令
CMD ["/app/start.sh"]

