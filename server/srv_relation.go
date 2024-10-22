package server

import (
	"context"
	"database/sql"
	"errors"

	"connectrpc.com/connect"
	v1 "github.com/core-pb/tag/tag/v1"
	"github.com/core-pb/tag/tag/v1/tagconnect"
	"github.com/uptrace/bun"
)

type relationship struct {
	*Server
	tagconnect.UnimplementedRelationshipHandler
}

func (x relationship) ListRelation(ctx context.Context, req *connect.Request[v1.ListRelationRequest]) (*connect.Response[v1.ListRelationResponse], error) {
	sq := x.db.NewSelect().Model(&Relation{})
	sq = InOrEqPure(sq, `"relation".module_id`, req.Msg.ModuleId)
	sq = InOrEqPure(sq, `"relation".external_id`, req.Msg.ExternalId)
	sq = InOrEqPure(sq, `"relation".tag_id`, req.Msg.TagId)
	sq = InOrEqPure(sq, `"relation".source_id`, req.Msg.SourceId)
	sq = QueryFormStruct(sq, `"relation".data`, req.Msg.Data)

	sq = Pagination(sq, req.Msg.Pagination)
	sq = Sorts(sq, req.Msg.Sort)

	var (
		arr        []*v1.Relation
		count, err = sq.ScanAndCount(ctx, &arr)
	)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.ListRelationResponse{Data: arr, Count: int64(count)}), nil
}

func (x relationship) SetRelation(ctx context.Context, req *connect.Request[v1.SetRelationRequest]) (*connect.Response[v1.SetRelationResponse], error) {
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
	) // TODO delete cache: parentTagID deleteTagID req.Msg.TagId

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

		if _, err := tx.NewInsert().Model(rel).Returning("*").Exec(ctx); err != nil {
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

			if _, err := tx.NewInsert().Model(&arr).Exec(ctx); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.SetRelationResponse{}), nil
}

func (x relationship) DeleteRelation(ctx context.Context, req *connect.Request[v1.DeleteRelationRequest]) (*connect.Response[v1.DeleteRelationResponse], error) {
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

	return connect.NewResponse(&v1.DeleteRelationResponse{}), nil
}

func (x relationship) DestroyExternalRelation(ctx context.Context, req *connect.Request[v1.DestroyExternalRelationRequest]) (*connect.Response[v1.DestroyExternalRelationResponse], error) {
	tx := x.db.NewDelete().Model(&Relation{}).Where("module_id = ?", req.Msg.ModuleId)
	if len(req.Msg.ExternalId) != 0 {
		tx = tx.Where("external_id IN (?)", bun.In(req.Msg.ExternalId))
	}

	if _, err := tx.Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.DestroyExternalRelationResponse{}), nil
}

func (x relationship) DestroyTagRelation(ctx context.Context, req *connect.Request[v1.DestroyTagRelationRequest]) (*connect.Response[v1.DestroyTagRelationResponse], error) {
	tx := x.db.NewDelete().Model(&Relation{}).Where("tag_id = ?", req.Msg.TagId)
	if len(req.Msg.ModuleId) != 0 {
		tx = tx.Where("module_id IN (?)", bun.In(req.Msg.ModuleId))
	}

	if _, err := tx.Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.DestroyTagRelationResponse{}), nil
}
