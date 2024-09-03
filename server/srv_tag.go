package main

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/core-pb/dt/time/v1"
	v1 "github.com/core-pb/tag/tag/v1"
	"github.com/uptrace/bun"
)

func (base) ListTag(ctx context.Context, req *connect.Request[v1.ListTagRequest]) (*connect.Response[v1.ListTagResponse], error) {
	sq := db.NewSelect().Model(&Tag{})
	sq = InOrEqPure(sq, `"tag".id`, req.Msg.Id)
	sq = InOrEqPure(sq, `"tag".key`, req.Msg.Key)
	sq = InOrEqPure(sq, `"tag".type_id`, req.Msg.TypeId)
	sq = InOrEqPure(sq, `"tag".parent_id`, req.Msg.ParentId)
	sq = QueryFormStruct(sq, `"tag".data`, req.Msg.Data)
	sq = QueryFormStruct(sq, `"tag".info`, req.Msg.Info)

	if len(req.Msg.ModuleId) != 0 {
		const moduleVisibleType = `SELECT DISTINCT unnest(visible_type) AS type_id FROM "module" WHERE id IN (?)`
		sq = sq.Join(`INNER JOIN (`+moduleVisibleType+`) AS mvt ON mvt.type_id = "tag".id`, bun.In(req.Msg.ModuleId))
	}

	sq = Pagination(sq, req.Msg.Pagination)
	sq = Sorts(sq, req.Msg.Sort)

	var (
		arr        []*v1.Tag
		count, err = sq.ScanAndCount(ctx, &arr)
	)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.ListTagResponse{Data: arr, Count: int64(count)}), nil
}

func (base) AddTag(ctx context.Context, req *connect.Request[v1.AddTagRequest]) (*connect.Response[v1.AddTagResponse], error) {
	val := &Tag{Tag: &v1.Tag{
		Key:      req.Msg.Key,
		TypeId:   req.Msg.TypeId,
		ParentId: req.Msg.ParentId,
		Data:     req.Msg.Data,
		Info:     req.Msg.Info,
	}}

	if !isInvalidKey(req.Msg.Key) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errInvalidKey)
	}

	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if has, err := tx.NewSelect().Model(&Type{}).Where("id = ?", val.TypeId).Exists(ctx); err != nil {
			return err
		} else if !has {
			return errors.New("type not exists")
		}

		if val.ParentId != 0 {
			if has, err := tx.NewSelect().Model(&Tag{}).Where("id = ?", val.ParentId).Exists(ctx); err != nil {
				return err
			} else if !has {
				return errors.New("tag parent not exists")
			}
		}

		_, err := db.NewInsert().Model(val).Returning("*").Exec(ctx)
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.AddTagResponse{Data: val.Tag}), nil
}

func (base) SetTag(ctx context.Context, req *connect.Request[v1.SetTagRequest]) (*connect.Response[v1.SetTagResponse], error) {
	var val Tag
	if _, err := db.NewUpdate().Model(&val).Where("id = ?", req.Msg.Id).Set("data = ?", req.Msg.Data).Returning("*").Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.SetTagResponse{Data: val.Tag}), nil
}

func (base) SetTagInfo(ctx context.Context, req *connect.Request[v1.SetTagInfoRequest]) (*connect.Response[v1.SetTagInfoResponse], error) {
	var val Tag
	if _, err := db.NewUpdate().Model(&val).Where("id = ?", req.Msg.Id).Set("info = ?", req.Msg.Info).Returning("*").Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.SetTagInfoResponse{Data: val.Tag}), nil
}

func (base) DeleteTag(ctx context.Context, req *connect.Request[v1.DeleteTagRequest]) (*connect.Response[v1.DeleteTagResponse], error) {
	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if has, err := tx.NewSelect().Model(&Relation{}).Where("tag_id IN (?)", bun.In(req.Msg.Id)).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("tag exist relation")
		}

		if has, err := tx.NewSelect().Model(&Tag{}).Where("parent_id IN (?)", bun.In(req.Msg.Id)).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("tag exist child")
		}

		t := time.Now()
		_, err := tx.NewUpdate().Model(&Tag{}).Where("id IN (?)", bun.In(req.Msg.Id)).
			Set(`deleted_at = ?, "key" = CONCAT(?, "key")`, t, fmt.Sprintf("::%d:", t.Seconds)).
			Exec(ctx)
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.DeleteTagResponse{}), nil
}

func (base) UpdateTagType(ctx context.Context, req *connect.Request[v1.UpdateTagTypeRequest]) (*connect.Response[v1.UpdateTagTypeResponse], error) {
	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if has, err := tx.NewSelect().Model(&Type{}).Where("id = ?", req.Msg.TypeId).Exists(ctx); err != nil {
			return err
		} else if !has {
			return errors.New("type not exists")
		}

		if has, err := tx.NewSelect().Model(&Relation{}).Where("tag_id IN (?)", bun.In(req.Msg.Id)).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("tag exist relation")
		}

		if has, err := tx.NewSelect().Model(&Tag{}).Where("parent_id IN (?)", bun.In(req.Msg.Id)).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("tag exist child")
		}

		if has, err := tx.NewSelect().Model(&Tag{}).Where("id IN (?) AND parent_id <> 0", bun.In(req.Msg.Id)).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("tag exist parent")
		}

		if _, err := db.NewUpdate().Model(&Tag{}).Where("id IN (?)", bun.In(req.Msg.Id)).Set(`type_id = ?`, req.Msg.TypeId).Exec(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.UpdateTagTypeResponse{}), nil
}

func (base) UpdateTagParent(ctx context.Context, req *connect.Request[v1.UpdateTagParentRequest]) (*connect.Response[v1.UpdateTagParentResponse], error) {
	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		typ := new(Type)
		if err := tx.NewSelect().Model(typ).Where(`id = (SELECT type_id FROM "tag" WHERE id = ?)`, req.Msg.ParentId).Scan(ctx); err != nil {
			return err
		}
		if typ.Inherit {
			return errors.New("tag type is inheritable")
		}

		if has, err := tx.NewSelect().Model(&Relation{}).Where("tag_id IN (?)", bun.In(req.Msg.Id)).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("tag exist relation")
		}

		if req.Msg.ParentId != 0 {
			var pt Tag
			if err := tx.NewSelect().Model(&pt).Where("id = ?", req.Msg.ParentId).Scan(ctx); err != nil {
				return err
			}
			if err := tx.NewSelect().Model(&pt).Where("type_id = ?", pt.TypeId).Scan(ctx); err != nil {
				return err
			}
			if err := CheckTagParentRecursiveLoop(ctx, tx, pt.TypeId, req.Msg.ParentId, req.Msg.Id); err != nil {
				return err
			}
		}

		if _, err := db.NewUpdate().Model(&Tag{}).Where("id IN (?)", bun.In(req.Msg.Id)).Set(`"parent_id" = ?`, req.Msg.ParentId).Exec(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.UpdateTagParentResponse{}), nil
}

func CheckTagParentRecursiveLoop(ctx context.Context, tx bun.Tx, typeID, checkID uint64, targetID []uint64) error {
	var arr []struct {
		ID       uint64 `json:"id"`
		ParentID uint64 `json:"parent_id"`
	}
	if err := tx.NewSelect().Model(&Tag{}).Where("type_id = ?", typeID).Scan(ctx, &arr); err != nil {
		return err
	}

	root := make(map[uint64]struct{})
	for _, v := range targetID {
		root[v] = struct{}{}
	}

	m := make(map[uint64]uint64)
	for _, v := range arr {
		m[v.ID] = v.ParentID
	}

	id := checkID
	for i := 0; i < len(arr); i++ {
		id = m[id]
		if id == 0 {
			break
		}
		if _, ok := root[id]; ok {
			return errors.New("tag parent recursive loop")
		}
	}

	return nil
}
