module github.com/chunyu224/kratos-bootstrap/logger/logrus

go 1.22.3

replace github.com/chunyu224/kratos-bootstrap/api => ../../api

require (
	github.com/chunyu224/kratos-bootstrap/api v0.0.0-00010101000000-000000000000
	github.com/go-kratos/kratos/contrib/log/logrus/v2 v2.0.0-20240516020449-fbac5fa25e7a
	github.com/go-kratos/kratos/v2 v2.7.3
	github.com/sirupsen/logrus v1.9.3
)

require (
	golang.org/x/sys v0.20.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)
