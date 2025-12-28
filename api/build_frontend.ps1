# PowerShell 前端构建脚本
Write-Host '开始构建前端项目...'

# 构建 front 项目
Set-Location ../web/front
Write-Host '正在构建 front 项目...'
npm install
npm run build

# 构建 console 项目
Set-Location ../console
Write-Host '正在构建 console 项目...'
npm install
npm run build

Write-Host '前端构建完成！'
Set-Location ../../api