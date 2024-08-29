package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/core-pb/dt/time/v1"
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
	if !isInvalidKey(req.Msg.Key) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errInvalidKey)
	}

	var val Module
	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if req.Msg.Id != 0 {
			_, err := tx.NewUpdate().Returning("*").Model(&val).Where("id = ?", req.Msg.Key).Set(`"key" = ?`, req.Msg.Key).Exec(ctx)
			return err
		}

		if err := tx.NewSelect().Model(&val).Where(`"key" = ?`, req.Msg.Key).Scan(ctx); err == nil || !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		val.Module = &v1.Module{Key: req.Msg.Key}
		_, err := tx.NewInsert().Model(&val).Exec(ctx)
		return err
	}); err != nil {
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
	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if has, err := tx.NewSelect().Model(&Relation{}).Where("module_id IN (?)", bun.In(req.Msg.Id)).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("relation has use module")
		}

		t := time.Now()
		_, err := tx.NewUpdate().Model(&Module{}).Where("id IN (?)", bun.In(req.Msg.Id)).
			Set(`deleted_at = ?, "key" = CONCAT(?, "key")`, t, fmt.Sprintf("::%d:", t.Seconds)).
			Exec(ctx)
		return err
	}); err != nil {
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
