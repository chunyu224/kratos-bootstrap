module github.com/chunyu224/kratos-bootstrap/config/polaris

go 1.22.3

replace github.com/chunyu224/kratos-bootstrap/api => ../../api

require (
	github.com/chunyu224/kratos-bootstrap/api v0.0.0-00010101000000-000000000000
	github.com/go-kratos/kratos/v2 v2.7.3
)

require (
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/kr/text v0.2.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
