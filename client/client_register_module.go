package client

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/core-pb/tag/tag/v1"
)

func (c *Client) RegisterModule(ctx context.Context, name string) (*v1.Module, error) {
	if val, ok := c.registeredModule.Load(name); ok {
		return val.(*v1.Module), nil
	}

	resp, err := c.Internal().RegisterModule(ctx, connect.NewRequest(&v1.RegisterModuleRequest{Key: name}))
	if err != nil {
		return nil, err
	}

	c.registeredModule.LoadOrStore(name, resp.Msg.Data)

	return resp.Msg.Data, nil
}
