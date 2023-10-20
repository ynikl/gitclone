# goclo
quick clone code from git server (like github) for `git clone` with organized in local machine

## 安装

`go install github.com/ynikl/gitlcone`

确保编译后文件在系统`$PATH`内

## 使用

`gitclone github.com/go-gorm/gen`

会将目标仓库代码下载到 `@代码下载路径 + github.com/go-gorm/gen` 文件夹中

代码下载路径配置依次按以下顺序获取

1. 环境变量 `CODE`
2. 环境变量 `GOPATH`
3. GO 环境变量 `go env GOPATH`
4. Home 目录
