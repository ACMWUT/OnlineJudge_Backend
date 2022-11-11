# 使用golang重构WUT OnlineJudge后端业务

## 技术栈

- 语言：Golang
- 框架：gin+gorm

## 部署
1. 交叉编译到Linux amd64平台
2. 将编译出来的可执行文件上传到服务器`/home/htdocs/dev/backend/app/`目录
3. 执行`/root`目录下的`deploy_master.sh`

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
scp OnlineJudge root@172.16.2.48:/home/htdocs/dev/backend/app/
sh deploy_master.sh
```

## 版权信息

版权所有 Copyright ©  2019 by 武汉理工大学 ACM 协会

All rights reserved