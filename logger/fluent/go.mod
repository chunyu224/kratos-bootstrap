module github.com/chunyu224/kratos-bootstrap/logger/fluent

go 1.22.3

replace github.com/chunyu224/kratos-bootstrap/api => ../../api

require (
	github.com/chunyu224/kratos-bootstrap/api v0.0.0-00010101000000-000000000000
	github.com/go-kratos/kratos/contrib/log/fluent/v2 v2.0.0-20240601080717-c0a7935bb120
	github.com/go-kratos/kratos/v2 v2.7.3
)

require (
	github.com/fluent/fluent-logger-golang v1.9.0 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/tinylib/msgp v1.1.9 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)
