package main

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/core-pb/tag/tag/v1"
	"github.com/core-pb/tag/tag/v1/tagconnect"
	"github.com/uptrace/bun"
)

type base struct {
	tagconnect.UnimplementedBaseHandler
}

func (base) ListModule(ctx context.Context, req *connect.Request[v1.ListModuleRequest]) (*connect.Response[v1.ListModuleResponse], error) {
	sq := db.NewSelect().Model(&Module{})
	sq = InOrEqPure(sq, `"module".id`, req.Msg.Id)
	sq = InOrEqPure(sq, `"module".key`, req.Msg.Key)
	sq = QueryFormStruct(sq, `"module".info`, req.Msg.Info)

	sq = Pagination(sq, req.Msg.Pagination)
	sq = Sorts(sq, req.Msg.Sort)

	var (
		arr        []*v1.Module
		count, err = sq.ScanAndCount(ctx, &arr)
	)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.ListModuleResponse{Data: arr, Count: int64(count)}), nil
}

func (base) SetModule(ctx context.Context, req *connect.Request[v1.SetModuleRequest]) (*connect.Response[v1.SetModuleResponse], error) {
	var val Module
	if _, err := db.NewUpdate().Model(&val).Where("id = ?", req.Msg.Id).Set("key = ?", req.Msg.Key).Returning("*").Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.SetModuleResponse{Data: val.Module}), nil
}

func (base) SetModuleInfo(ctx context.Context, req *connect.Request[v1.SetModuleInfoRequest]) (*connect.Response[v1.SetModuleInfoResponse], error) {
	var val Module
	if _, err := db.NewUpdate().Model(&val).Where("id = ?", req.Msg.Id).Set("info = ?", req.Msg.Info).Returning("*").Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.SetModuleInfoResponse{Data: val.Module}), nil
}

func (base) DeleteModule(ctx context.Context, req *connect.Request[v1.DeleteModuleRequest]) (*connect.Response[v1.DeleteModuleResponse], error) {
	if _, err := db.NewDelete().Model(&Module{}).Where("id IN (?)", bun.In(req.Msg.Id)).Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.DeleteModuleResponse{}), nil
}

func (base) UpdateModuleVisibleType(ctx context.Context, req *connect.Request[v1.UpdateModuleVisibleTypeRequest]) (*connect.Response[v1.UpdateModuleVisibleTypeResponse], error) {
	if _, err := db.NewUpdate().Model(&Module{}).Where("id IN (?)", bun.In(req.Msg.Id)).Set("visible_type = ?", req.Msg.VisibleType).Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.UpdateModuleVisibleTypeResponse{}), nil
}
