package tag

import (
	_ "github.com/core-pb/dt"
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	_ "google.golang.org/protobuf/proto"
)

//go:generate buf generate pb
//go:generate buf generate pb --template buf.gen.tag.yaml
