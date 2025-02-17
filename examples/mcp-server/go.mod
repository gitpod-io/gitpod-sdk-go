module github.com/gitpod-io/gitpod-sdk-go/examples/mcp-server

go 1.23.0

replace github.com/gitpod-io/gitpod-sdk-go => ../..

require github.com/mark3labs/mcp-go v0.8.4

require (
	github.com/gitpod-io/gitpod-sdk-go v0.0.0-00010101000000-000000000000 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
)
