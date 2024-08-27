module github.com/core-pb/tag/app/tag

go 1.23.0

replace github.com/core-pb/tag => ../..

require (
	connectrpc.com/connect v1.16.2
	github.com/core-pb/dt v1.1.0
	github.com/core-pb/tag v0.0.0-00010101000000-000000000000
	github.com/redis/rueidis v1.0.44
	github.com/uptrace/bun v1.2.1
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/srikrsna/protoc-gen-gotag v1.0.2 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/uptrace/bun/driver/pgdriver v1.2.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	go.x2ox.com/sorbifolia/bunpgd v0.0.0-20240825064625-530faf278cc7 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	mellium.im/sasl v0.3.1 // indirect
)
