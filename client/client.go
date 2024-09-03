package client

import (
	"connectrpc.com/connect"
	"github.com/core-pb/tag/tag/v1/tagconnect"
	"go.x2ox.com/sorbifolia/crpc"
)

type Client struct {
	hc   connect.HTTPClient
	addr string
	opts []connect.ClientOption
}

func New(hc connect.HTTPClient, addr string, opts ...connect.ClientOption) *Client {
	return &Client{addr: addr, hc: hc, opts: opts}
}

func (c *Client) Base() tagconnect.BaseClient {
	return crpc.Client(tagconnect.NewBaseClient, c.hc, c.addr, c.opts...)
}

func (c *Client) Internal() tagconnect.InternalClient {
	return crpc.Client(tagconnect.NewInternalClient, c.hc, c.addr, c.opts...)
}

func (c *Client) Relationship() tagconnect.RelationshipClient {
	return crpc.Client(tagconnect.NewRelationshipClient, c.hc, c.addr, c.opts...)
}
