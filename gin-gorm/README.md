# gin-gorm

#### 介绍
gin-gorm

#### 软件架构
基于 go 语言 gin 框架，融合了 gorm 对象关系模型




#### 使用说明

1、 本地有 go 环境，可以直接运行main.go
2、 mac 可直接运行 gin-gorm 二进制文件
3、 linux amd64 可直接运行 main 二进制文件

#### 交叉编译

`go
GOOS=linux GOARCH=amd64 go build main.go
`

| OS |	ARCH |	OS version|
|----|----  |----     |
|linux	|386 / amd64 / arm |	>= Linux 2.6 |
|darwin |386 / amd64	| OS X (Snow Leopard + Lion) |
|freebsd |	386 / amd64	| >= FreeBSD 7 |
|windows |	386 / amd64 |	>= Windows 2000 |
