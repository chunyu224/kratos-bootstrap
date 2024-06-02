module github.com/chunyu224/kratos-bootstrap/registry/eureka

go 1.22.3

replace github.com/chunyu224/kratos-bootstrap/api => ../../api

require (
	github.com/chunyu224/kratos-bootstrap/api v0.0.0-00010101000000-000000000000
	github.com/go-kratos/kratos/contrib/registry/eureka/v2 v2.0.0-20240516020449-fbac5fa25e7a
	github.com/go-kratos/kratos/v2 v2.7.3
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
