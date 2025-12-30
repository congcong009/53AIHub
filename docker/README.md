# 53AIHub Docker 部署说明

本文档说明如何使用 Docker 部署 53AIHub 的前后端服务。

## 目录结构要求

确保以下目录存在并包含相应的源代码：

```
/home/aihub/
├── api/                    # 后端源代码目录
│   ├── static/
│   │   ├── front/         # 前端构建产物目录
│   │   └── console/       # 管理后台构建产物目录
│   └── docs/              # Swagger 文档目录
└── web/                    # 前端源代码目录
    ├── front/             # 前台源代码
    └── console/           # 管理后台源代码
```

## 构建说明

### 方法一：完整构建（包含前端构建）

```bash
cd /home/aihub/docker
docker-compose build
```

此方法将在 Docker 构建过程中自动：
1. 安装前端依赖
2. 构建前端项目（front 和 console）
3. 将前端构建产物集成到后端应用
4. 构建完整的后端应用

> 注意：此方法需要服务器有充足的资源（内存、CPU）来构建前端项目

### 方法二：快速构建（使用预构建前端）

如果服务器资源有限，可以先在本地构建前端：

```bash
# 在本地机器上执行
cd web/front
npm install
npm run build

cd ../console
npm install
npm run build

# 将构建产物复制到 api/static 目录
cp -r front/dist ../api/static/front/
cp -r console/dist ../api/static/console/
```

然后使用简化版 Dockerfile 构建：

```bash
cd /home/aihub/api
#docker build -f Dockerfile -t aihub_web .
docker build -f Dockerfile7 -t cdqz/aihub_core .
```

## 启动服务

```bash
cd /home/aihub/docker
docker-compose up -d
```

## 服务端口

- **应用访问**: http://localhost:3040

## 环境变量配置

创建 `.env` 文件（在 docker 目录下），配置必要的环境变量：

```env
PORT=3000
REDIS_HOST=redis
MYSQL_HOST=mysql
REDIS_PASSWORD=redis@2025
LOG_LEVEL=INFO
# 其他需要的环境变量...
```

## 故障排除

### 构建失败

如果遇到构建失败，检查：

1. 服务器资源是否充足（至少 2GB 内存）
2. Node.js 依赖安装是否成功
3. 网络连接是否正常（用于下载依赖）

### 端口访问问题

- 容器内端口：3000
- 主机映射端口：3040
- 环境变量 PORT：3000（容器内监听端口）

### 前端资源问题

如果前端页面无法加载，请确认：
- static/front/dist 目录包含前端构建产物
- static/console/dist 目录包含管理后台构建产物
- embed 指令路径与实际目录结构匹配

### 查看日志
```bash
docker-compose logs -f web      # 查看应用日志
```
