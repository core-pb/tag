module github.com/core-pb/tag/client

go 1.23.0

replace github.com/core-pb/tag => ..

require (
	connectrpc.com/connect v1.16.2
	github.com/bufbuild/httplb v0.3.0
	github.com/core-pb/dt v1.1.0
	github.com/core-pb/tag v0.0.0-00010101000000-000000000000
)

require (
	github.com/srikrsna/protoc-gen-gotag v1.0.2 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
