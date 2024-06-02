#!/bin/bash

set -x
#指定起始文件夹
DIR=`pwd`

cd $DIR/config/apollo
go get all
go mod tidy

cd $DIR/config/consul
go get all
go mod tidy

cd $DIR/config/etcd
go get all
go mod tidy

cd $DIR/config/kubernetes
go get all
go mod tidy

cd $DIR/config/nacos
go get all
go mod tidy

cd $DIR/config/polaris
go get all
go mod tidy

cd $DIR/logger/aliyun
go get all
go mod tidy

cd $DIR/logger/fluent
go get all
go mod tidy

cd $DIR/logger/logrus
go get all
go mod tidy

cd $DIR/logger/tencent
go get all
go mod tidy

cd $DIR/logger/zap
go get all
go mod tidy

cd $DIR/oss/minio
go get all
go mod tidy

cd $DIR/registry/consul
go get all
go mod tidy

cd $DIR/registry/etcd
go get all
go mod tidy

cd $DIR/registry/eureka
go get all
go mod tidy

cd $DIR/registry/kubernetes
go get all
go mod tidy

cd $DIR/registry/nacos
go get all
go mod tidy

cd $DIR/registry/polaris
go get all
go mod tidy

cd $DIR/registry/servicecomb
go get all
go mod tidy

cd $DIR/registry/zookeeper
go get all
go mod tidy

cd $DIR
go get all
go mod tidy