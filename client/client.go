package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/bufbuild/httplb"
	"github.com/bufbuild/httplb/picker"
	"github.com/core-pb/dt/query/v1"
	v1 "github.com/core-pb/tag/tag/v1"
	"github.com/core-pb/tag/tag/v1/tagconnect"
)

var lbc = httplb.NewClient(
	httplb.WithPicker(picker.NewPowerOfTwo),
	httplb.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}, time.Second*3),
)

func aa(hc connect.HTTPClient, addr string, opts ...connect.ClientOption) tagconnect.BaseClient {
	if hc == nil {
		hc = lbc
	}
	if !strings.HasPrefix(addr, "https://") {
		addr = fmt.Sprintf("https://%s", addr)
	}

	return tagconnect.NewBaseClient(hc, addr, append(opts, connect.WithGRPC())...)
}

func a() {
	ctx := context.Background()
	c := aa(nil, "10.9.8.8:32443")

	resp, err := c.ListType(ctx, connect.NewRequest(&v1.ListTypeRequest{
		Pagination: &query.Pagination{Page: 1, PageSize: 10},
		Sort:       nil,
		Id:         nil,
		Key:        nil,
		Info:       nil,
		Inherit:    nil,
		Exclusive:  nil,
		ModuleId:   nil,
	}))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
