module github.com/chunyu224/kratos-bootstrap/logger/tencent

go 1.22.3

replace github.com/chunyu224/kratos-bootstrap/api => ../../api

require (
	github.com/chunyu224/kratos-bootstrap/api v0.0.0-00010101000000-000000000000
	github.com/go-kratos/kratos/contrib/log/tencent/v2 v2.0.0-20240516020449-fbac5fa25e7a
	github.com/go-kratos/kratos/v2 v2.7.3
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/tencentcloud/tencentcloud-cls-sdk-go v1.0.9 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)
