#!/bin/bash

# 前端构建脚本
echo '开始构建前端项目...'

# 构建 front 项目
cd ../web/front
echo '正在构建 front 项目...'
npm install
npm run build

# 构建 console 项目
cd ../console
echo '正在构建 console 项目...'
npm install
npm run build

echo '前端构建完成！'