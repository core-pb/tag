package client

import (
	"fmt"
	"testing"

	"connectrpc.com/connect"
	"github.com/core-pb/dt/query/v1"
	v1 "github.com/core-pb/tag/tag/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestListType(t *testing.T) {
	resp, err := client.Base().ListType(ctx, connect.NewRequest(&v1.ListTypeRequest{
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

func TestSetType(t *testing.T) {
	resp, err := client.Base().SetType(ctx, connect.NewRequest(&v1.SetTypeRequest{
		Key: "test1",
	}))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}

func TestSetTypeInfo(t *testing.T) {
	resp, err := client.Base().SetType(ctx, connect.NewRequest(&v1.SetTypeRequest{
		Key: "user:role",
	}))
	if err != nil {
		return
	}

	var set *connect.Response[v1.SetTypeInfoResponse]
	set, err = client.Base().SetTypeInfo(ctx, connect.NewRequest(&v1.SetTypeInfoRequest{
		Id: resp.Msg.Data.Id,
		Info: &structpb.Struct{Fields: map[string]*structpb.Value{
			"name": structpb.NewStringValue("is name"),
		}},
	}))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(set)
}

func TestUpdateTypeInherit(t *testing.T) {
	resp, err := client.Base().SetType(ctx, connect.NewRequest(&v1.SetTypeRequest{
		Key: "user:role",
	}))
	if err != nil {
		return
	}

	_, err = client.Base().UpdateTypeInherit(ctx, connect.NewRequest(&v1.UpdateTypeInheritRequest{
		Id:      []uint64{resp.Msg.Data.Id},
		Inherit: false,
	}))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = client.Base().UpdateTypeInherit(ctx, connect.NewRequest(&v1.UpdateTypeInheritRequest{
		Id:      []uint64{resp.Msg.Data.Id},
		Inherit: true,
	}))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestUpdateTypeExclusive(t *testing.T) {
	resp, err := client.Base().SetType(ctx, connect.NewRequest(&v1.SetTypeRequest{
		Key: "user:role",
	}))
	if err != nil {
		return
	}

	_, err = client.Base().UpdateTypeExclusive(ctx, connect.NewRequest(&v1.UpdateTypeExclusiveRequest{
		Id:        []uint64{resp.Msg.Data.Id},
		Exclusive: false,
	}))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = client.Base().UpdateTypeExclusive(ctx, connect.NewRequest(&v1.UpdateTypeExclusiveRequest{
		Id:        []uint64{resp.Msg.Data.Id},
		Exclusive: true,
	}))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestDeleteType(t *testing.T) {
	resp, err := client.Base().SetType(ctx, connect.NewRequest(&v1.SetTypeRequest{
		Key: "user:role",
	}))
	if err != nil {
		return
	}

	_, err = client.Base().DeleteType(ctx, connect.NewRequest(&v1.DeleteTypeRequest{
		Id: []uint64{resp.Msg.Data.Id},
	}))
	if err != nil {
		fmt.Println(err)
		return
	}
}
