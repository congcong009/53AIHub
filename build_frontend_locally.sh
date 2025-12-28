#!/bin/bash

# 在本地构建前端项目（需要Node.js环境）
echo '开始构建前端项目...'

# 构建 front 项目
cd web/front
echo '正在构建 front 项目...'
npm install
npm run build

# 构建 console 项目
cd ../console
echo '正在构建 console 项目...'
npm install
npm run build

# 将构建结果复制到 api/static 目录
cd ..
echo '复制构建结果到 api/static...'
cp -r front/dist ../api/static/front/
cp -r console/dist ../api/static/console/

echo '前端构建并复制完成！'
echo '现在可以使用普通Dockerfile进行构建，而无需在Docker构建过程中重复构建前端'