# gboot
[![Build](https://github.com/neoean/gboot/workflows/Build/badge.svg?branch=main&event=push)](https://github.com/neoean/gboot/actions?query=workflow%3ABuild+event%3Apush+branch%3Amain)
[![Tests](https://github.com/neoean/gboot/workflows/Tests/badge.svg?branch=main&event=push)](https://github.com/neoean/gboot/actions?query=workflow%3ATests+event%3Apush+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/neoean/gboot)](https://goreportcard.com/report/github.com/neoean/gboot)

可以用一条命令，根据数据库或者SQL 直接生成完整的golang项目，包括数据模型，dao，service，api等

使用方法：
``` sh
方法一：
# 先安装命令，确保~/go/bin 在PATH里
go install github.com/neoean/gboot@latest
gboot -name=PROJECT_NAME -dbtype=mysql -username=root -password=123456 -host=127.0.0.1:3306 -dbName=test

方法二：
git clone git@github.com:neoean/gboot.git
go main.go -name=PROJECT_NAME -dbtype=mysql -username=root -password=123456 -host=127.0.0.1:3306 -dbName=test
```