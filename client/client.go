package client

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/bufbuild/httplb"
	"github.com/bufbuild/httplb/picker"
	"github.com/core-pb/tag/tag/v1/tagconnect"
)

var lbc = httplb.NewClient(
	httplb.WithPicker(picker.NewPowerOfTwo),
	httplb.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}, time.Second*3),
)

func Client(hc connect.HTTPClient, addr string, opts ...connect.ClientOption) tagconnect.BaseClient {
	if hc == nil {
		hc = lbc
	}
	if !strings.HasPrefix(addr, "https://") {
		addr = fmt.Sprintf("https://%s", addr)
	}

	return tagconnect.NewBaseClient(hc, addr, append(opts, connect.WithGRPC())...)
}
