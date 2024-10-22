package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"connectrpc.com/connect"
	"github.com/core-pb/dt/time/v1"
	v1 "github.com/core-pb/tag/tag/v1"
	"github.com/core-pb/tag/tag/v1/tagconnect"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type itn struct {
	*Server

	tagconnect.UnimplementedInternalHandler
}

func (x itn) GetTagIDTreeize(ctx context.Context, req *connect.Request[v1.GetTagIDTreeizeRequest]) (*connect.Response[v1.GetTagIDTreeizeResponse], error) {
	var (
		typ = new(v1.Type)
		arr []v1.FlatTagID
	)

	switch v := req.Msg.GetFrom().(type) {
	case *v1.GetTagIDTreeizeRequest_TagId:
		const findTagType = `SELECT type_id FROM "tag" WHERE id = ?`
		if err := x.db.NewSelect().Model(&Type{}).Join(`INNER JOIN (`+findTagType+`) AS typ ON typ.type_id = "type".id`, v.TagId).Scan(ctx, typ); err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, err)
		}

	case *v1.GetTagIDTreeizeRequest_TypeId:
		if err := x.db.NewSelect().Model(&Type{}).Where("id = ?", v.TypeId).Scan(ctx, typ); err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, err)
		}

	default:
		return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unknown from is not implemented"))
	}

	if err := x.db.NewSelect().Model(&Tag{}).Column("id", "parent_id").Where(`"type_id" = ?`, typ.Id).Scan(ctx, &arr); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	tree := new(v1.TagIDTreeize)
	tree.FromFlat(arr)

	return connect.NewResponse(&v1.GetTagIDTreeizeResponse{Data: tree, Type: typ}), nil
}

func (x itn) BindRelation(ctx context.Context, req *connect.Request[v1.BindRelationRequest]) (*connect.Response[v1.BindRelationResponse], error) {
	var (
		module                   Module
		tag                      Tag
		typ                      Type
		parentTagID, deleteTagID []uint64

		rel = &Relation{Relation: &v1.Relation{
			ModuleId:   req.Msg.ModuleId,
			ExternalId: req.Msg.ExternalId,
			TagId:      req.Msg.TagId,
			Data:       req.Msg.Data,
		}}
	)

	if err := x.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if err := tx.NewSelect().For("UPDATE").Model(&module).Where("id = ?", req.Msg.ModuleId).Scan(ctx); err != nil {
			return err
		}
		if err := tx.NewSelect().For("UPDATE").Model(&tag).Where("id = ?", req.Msg.TagId).Scan(ctx); err != nil {
			return err
		}
		if err := tx.NewSelect().For("UPDATE").Model(&typ).Where("id = ?", tag.TypeId).Scan(ctx); err != nil {
			return err
		}

		if typ.Exclusive {
			if err := tx.NewRaw(`DELETE FROM "relation" USING "tag" WHERE relation.tag_id = tag.id AND tag.type_id = ? RETURNING tag.id`, typ.Id).Scan(ctx, &deleteTagID); err != nil {
				return err
			}
		}

		if _, err := tx.NewInsert().On("CONFLICT (module_id,external_id,tag_id) DO UPDATE").
			Set("source_id = EXCLUDED.source_id, data = EXCLUDED.data").Model(rel).Returning("*").Exec(ctx); err != nil {
			return err
		}

		if typ.Inherit && tag.ParentId != 0 { // 需要查父级
			for {
				var val Tag
				if err := tx.NewSelect().For("UPDATE").Model(&val).Column("id", "parent_id").Where("id = ?", tag.ParentId).Scan(ctx); err != nil {
					return err
				}
				parentTagID = append(parentTagID, val.Id)
				if val.ParentId == 0 {
					break
				}
			}

			arr := make([]*Relation, 0, len(parentTagID))
			for _, v := range parentTagID {
				arr = append(arr, &Relation{Relation: &v1.Relation{
					ModuleId:   rel.ModuleId,
					ExternalId: rel.ExternalId,
					TagId:      v,
					SourceId:   rel.TagId,
					CreatedAt:  rel.CreatedAt,
					UpdatedAt:  rel.UpdatedAt,
				}})
			}

			if _, err := tx.NewInsert().On("CONFLICT (module_id,external_id,tag_id) DO UPDATE").
				Set("source_id = EXCLUDED.source_id, data = EXCLUDED.data").Model(&arr).Exec(ctx); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.BindRelationResponse{CleanTagId: deleteTagID, InheritTagId: parentTagID}), nil
}

func (x itn) UnbindRelation(ctx context.Context, req *connect.Request[v1.UnbindRelationRequest]) (*connect.Response[v1.UnbindRelationResponse], error) {
	var cleanID []uint64

	if err := x.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if has, err := tx.NewSelect().For("UPDATE").Model(&Relation{}).Where(
			"module_id = ? AND external_id = ? AND tag_id = ?", req.Msg.ModuleId, req.Msg.ExternalId, req.Msg.TagId,
		).Exists(ctx); err != nil {
			return err
		} else if !has {
			return nil
		}

		if err := tx.NewSelect().For("UPDATE").Model(&Relation{}).Column("tag_id").Where(
			"module_id = ? AND external_id = ? AND source_id = ?", req.Msg.ModuleId, req.Msg.ExternalId, req.Msg.TagId,
		).Scan(ctx, &cleanID); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		del := []uint64{req.Msg.TagId}
		if len(cleanID) != 0 {
			del = append(del, cleanID...)
		}

		if _, err := tx.NewDelete().Model(&Relation{}).Where(
			"module_id = ? AND external_id = ? AND tag_id IN (?)", req.Msg.ModuleId, req.Msg.ExternalId, bun.In(del),
		).Exec(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.UnbindRelationResponse{CleanTagId: cleanID}), nil
}

func (x itn) GetAllByModule(ctx context.Context, req *connect.Request[v1.GetAllByModuleRequest]) (*connect.Response[v1.GetAllByModuleResponse], error) {
	var (
		typeID []uint64
		typ    []*v1.Type
		tag    []*v1.Tag
	)
	if err := x.db.NewSelect().Model(&Module{}).ColumnExpr("DISTINCT unnest(visible_type)").Where("id = ?", req.Msg.ModuleId).Scan(ctx, &typeID); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}
	if err := x.db.NewSelect().Model(&Type{}).Where("id IN (?)", bun.In(typeID)).Scan(ctx, &typ); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}
	if err := x.db.NewSelect().Model(&Tag{}).Where("type_id IN (?)", bun.In(typeID)).Scan(ctx, &tag); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.GetAllByModuleResponse{Type: typ, Tag: tag}), nil
}

func (x itn) RegisterModule(ctx context.Context, req *connect.Request[v1.RegisterModuleRequest]) (*connect.Response[v1.RegisterModuleResponse], error) {
	var m Module
	if err := x.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if err := tx.NewSelect().Model(&m).Where(`"key" = ?`, req.Msg.Key).Scan(ctx); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				m.Module = &v1.Module{Key: req.Msg.Key}
				_, err = tx.NewInsert().Model(&m).Exec(ctx)
			}
			return err
		}
		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.RegisterModuleResponse{Data: m.Module}), nil
}

func (x itn) RegisterTag(ctx context.Context, req *connect.Request[v1.RegisterTagRequest]) (*connect.Response[v1.RegisterTagResponse], error) {
	arr := make([]*v1.Tag, 0, len(req.Msg.Data))
	if err := x.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		for i := range req.Msg.Data {
			var (
				tag                     = req.Msg.Data[i]
				moduleKey, typeKey, key string
				ok                      bool
			)

			if moduleKey, key, ok = strings.Cut(tag.GetKey(), ":"); !ok {
				return errors.New("invalid tag key: notfound module")
			}
			if typeKey, _, ok = strings.Cut(key, ":"); !ok {
				return errors.New("invalid tag key: notfound type")
			}

			var (
				typKey = fmt.Sprintf("%s:%s", moduleKey, typeKey)

				module Module
				typ    Type
				val    Tag
			)

			if err := tx.NewSelect().Model(&typ).Where("key = ?", typKey).Scan(ctx); err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					return err
				}
				typ.Type = &v1.Type{Key: typKey}

				if _, err = tx.NewInsert().Model(&typ).Exec(ctx); err != nil {
					return err
				}
			}
			if err := tx.NewSelect().Model(&module).Where("key = ?", moduleKey).Scan(ctx); err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					return err
				}
				module.Module = &v1.Module{Key: moduleKey, VisibleType: []uint64{typ.Id}}

				if _, err = tx.NewInsert().Model(&module).Exec(ctx); err != nil {
					return err
				}
			}

			if err := tx.NewSelect().Model(&val).Where("key = ?", tag.GetKey()).Scan(ctx); err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					return err
				}
				val.Tag = &v1.Tag{Key: tag.GetKey(), Data: tag.Data, TypeId: typ.Id}

				if _, err = tx.NewInsert().Model(&val).Exec(ctx); err != nil {
					return err
				}
			}
			arr = append(arr, val.Tag)
		}

		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.RegisterTagResponse{Data: arr}), nil
}

func (x itn) SetTypeWithModule(ctx context.Context, req *connect.Request[v1.SetTypeWithModuleRequest]) (*connect.Response[v1.SetTypeWithModuleResponse], error) {
	if req.Msg.TypeId == nil {
		if err := x.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			var m Module
			if err := tx.NewSelect().Model(&m).For("UPDATE").Where(`id = ?`, req.Msg.ModuleId).Scan(ctx); err != nil {
				return err
			}

			typ := &Type{Type: &v1.Type{
				Key:       defaultValue(req.Msg.Key),
				Info:      req.Msg.Info,
				Inherit:   defaultValue(req.Msg.Inherit),
				Exclusive: defaultValue(req.Msg.Exclusive),
			}}
			if typ.Key == "" {
				typ.Key = fmt.Sprintf("%s:%s", m.Key, uuid.New())
			}

			if _, err := tx.NewInsert().Model(typ).Returning("id").Exec(ctx); err != nil {
				return err
			}

			if _, err := tx.NewUpdate().Model(&Module{}).Where("id = ?", m.Id).Set("visible_type", append(m.VisibleType, typ.Id)).Exec(ctx); err != nil {
				return err
			}

			return nil
		}); err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, err)
		}
		return connect.NewResponse(&v1.SetTypeWithModuleResponse{}), nil
	}

	if req.Msg.Info != nil {
		if _, err := (base{}).SetTypeInfo(ctx, connect.NewRequest(&v1.SetTypeInfoRequest{
			Id: *req.Msg.TypeId, Info: req.Msg.Info,
		})); err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, err)
		}
	}
	if req.Msg.Inherit != nil {
		if _, err := (base{}).UpdateTypeInherit(ctx, connect.NewRequest(&v1.UpdateTypeInheritRequest{
			Id: []uint64{*req.Msg.TypeId}, Inherit: *req.Msg.Inherit,
		})); err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, err)
		}
	}
	if req.Msg.Exclusive != nil {
		if _, err := (base{}).UpdateTypeExclusive(ctx, connect.NewRequest(&v1.UpdateTypeExclusiveRequest{
			Id: []uint64{*req.Msg.TypeId}, Exclusive: *req.Msg.Exclusive,
		})); err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, err)
		}
	}

	return connect.NewResponse(&v1.SetTypeWithModuleResponse{}), nil
}

func (x itn) DeleteTypeWithModule(ctx context.Context, req *connect.Request[v1.DeleteTypeWithModuleRequest]) (*connect.Response[v1.DeleteTypeWithModuleResponse], error) {
	if err := x.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		var m Module
		if err := tx.NewSelect().Model(&m).For("UPDATE").Where(`id = ?`, req.Msg.ModuleId).Scan(ctx); err != nil {
			return err
		}

		if !arrayContains(m.VisibleType, req.Msg.TypeId) {
			return errors.New("visible type not contains")
		}

		t := time.Now()
		if _, err := tx.NewUpdate().Model(&Type{}).Where("id IN (?)", bun.In(req.Msg.TypeId)).
			Set(`deleted_at = ?, "key" = CONCAT(?, "key")`, t, fmt.Sprintf("::%d:", t.Seconds)).
			Exec(ctx); err != nil {
			return err
		}

		if _, err := tx.NewUpdate().Model(&Module{}).Where("id = ?", m.Id).Set("visible_type", arrayDifference(m.VisibleType, req.Msg.TypeId)).Exec(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.DeleteTypeWithModuleResponse{}), nil
}

func (x itn) SetTagWithModule(ctx context.Context, req *connect.Request[v1.SetTagWithModuleRequest]) (*connect.Response[v1.SetTagWithModuleResponse], error) {
	if req.Msg.TagId == nil {
		if err := x.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			var (
				m   Module
				typ Type
			)
			if err := tx.NewSelect().Model(&m).For("UPDATE").Where(`id = ?`, req.Msg.ModuleId).Scan(ctx); err != nil {
				return err
			}
			if err := tx.NewSelect().Model(&typ).Where(`id = ?`, *req.Msg.TypeId).Scan(ctx); err != nil {
				return err
			}

			if !arrayContains(m.VisibleType, []uint64{*req.Msg.TypeId}) {
				return errors.New("visible type not contains")
			}

			tag := &Tag{Tag: &v1.Tag{
				Key:      defaultValue(req.Msg.Key),
				TypeId:   typ.Id,
				ParentId: defaultValue(req.Msg.ParentId),
				Data:     req.Msg.Data,
				Info:     req.Msg.Info,
			}}
			if tag.Key == "" {
				tag.Key = fmt.Sprintf("%s:%s", typ.Key, uuid.New())
			}

			if _, err := tx.NewInsert().Model(tag).Returning("id").Exec(ctx); err != nil {
				return err
			}

			return nil
		}); err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, err)
		}
		return connect.NewResponse(&v1.SetTagWithModuleResponse{}), nil
	}

	if req.Msg.TypeId != nil {
		if _, err := (base{}).UpdateTagType(ctx, connect.NewRequest(&v1.UpdateTagTypeRequest{
			Id: []uint64{*req.Msg.TagId}, TypeId: *req.Msg.TypeId,
		})); err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, err)
		}
	}
	if req.Msg.ParentId != nil {
		if _, err := (base{}).UpdateTagParent(ctx, connect.NewRequest(&v1.UpdateTagParentRequest{
			Id: []uint64{*req.Msg.TagId}, ParentId: *req.Msg.ParentId,
		})); err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, err)
		}
	}
	if req.Msg.Data != nil {
		if _, err := (base{}).SetTag(ctx, connect.NewRequest(&v1.SetTagRequest{
			Id: *req.Msg.TagId, Data: req.Msg.Data,
		})); err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, err)
		}
	}
	if req.Msg.Info != nil {
		if _, err := (base{}).SetTagInfo(ctx, connect.NewRequest(&v1.SetTagInfoRequest{
			Id: *req.Msg.TagId, Info: req.Msg.Data,
		})); err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, err)
		}
	}

	return connect.NewResponse(&v1.SetTagWithModuleResponse{}), nil
}

func (x itn) DeleteTagWithModule(ctx context.Context, req *connect.Request[v1.DeleteTagWithModuleRequest]) (*connect.Response[v1.DeleteTagWithModuleResponse], error) {
	if err := x.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if has, err := tx.NewSelect().Model(&Module{}).For("UPDATE").Where(`id = ?`, req.Msg.ModuleId).Exists(ctx); err != nil {
			return err
		} else if !has {
			return errors.New("module not exists")
		}

		t := time.Now()
		if _, err := tx.NewUpdate().Model(&Tag{}).Where("id IN (?)", bun.In(req.Msg.TagId)).
			Set(`deleted_at = ?, "key" = CONCAT(?, "key")`, t, fmt.Sprintf("::%d:", t.Seconds)).
			Exec(ctx); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.DeleteTagWithModuleResponse{}), nil
}
