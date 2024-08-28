package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/core-pb/dt/query/v1"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/types/known/structpb"
)

var ErrPageTooBig = errors.New("page too big")

func Pagination(tx *bun.SelectQuery, x *query.Pagination) *bun.SelectQuery {
	if x == nil {
		return nil
	}

	page, pageSize := int(x.Page), int(x.PageSize)
	if page < 1 {
		page = 1
	}

	switch {
	case pageSize == -1:
		return tx
	case pageSize < 1:
		pageSize = 1
	case pageSize > 1000:
		pageSize = 1000
	}

	switch {
	case page*pageSize > 100*10000: // query cost to big
		return tx.Err(ErrPageTooBig)
	case page == 1:
		return tx.Limit(pageSize)
	default:
		return tx.Limit(pageSize).Offset((page - 1) * pageSize)
	}
}

func Sort(tx *bun.SelectQuery, x *query.Sort) *bun.SelectQuery {
	switch v := x.GetKey().(type) {
	case *query.Sort_Desc:
		return tx.Order(v.Desc + " DESC")
	case *query.Sort_Asc:
		return tx.Order(v.Asc)
	}
	return tx
}

func Sorts(tx *bun.SelectQuery, arr []*query.Sort) *bun.SelectQuery {
	for _, v := range arr {
		tx = Sort(tx, v)
	}
	return tx
}

func InOrEq[T any](tx *bun.SelectQuery, key string, val []T) *bun.SelectQuery {
	switch len(val) {
	case 0:
		return tx
	case 1:
		return tx.Where(`? = ?`, bun.Ident(key), val[0])
	default:
		return tx.Where(`? IN (?)`, bun.Ident(key), bun.In(val))
	}
}

func InOrEqPure[T any](tx *bun.SelectQuery, key string, val []T) *bun.SelectQuery {
	switch len(val) {
	case 0:
		return tx
	case 1:
		return tx.Where(fmt.Sprintf(`%s = ?`, key), val[0])
	default:
		return tx.Where(fmt.Sprintf(`%s IN (?)`, key), bun.In(val))
	}
}

func OrLikeEq(tx *bun.SelectQuery, key string, val []string) *bun.SelectQuery {
	var precise []string

	for i := range val {
		if val[i] == "" || val[i] == "*" {
		} else if strings.ContainsAny(val[i], "*") {
			tx = tx.WhereOr(`? LIKE ?`, bun.Ident(key), strings.ReplaceAll(val[i], "*", "%"))
		} else {
			precise = append(precise, val[i])
		}
	}

	switch len(precise) {
	case 0:
		return tx
	case 1:
		return tx.WhereOr(`? = ?`, bun.Ident(key), val[0])
	default:
		return tx.WhereOr(`? IN (?)`, bun.Ident(key), bun.In(val))
	}
}

func QueryFormStruct(tx *bun.SelectQuery, qk string, val *structpb.Struct) *bun.SelectQuery {
	for k, v := range val.GetFields() {
		switch v := v.Kind.(type) {
		case nil:
			continue
		case *structpb.Value_NullValue:
			tx = tx.Where(fmt.Sprintf("%s->'%s' = 'null'", qk, k))
			continue
		case *structpb.Value_StructValue:
			tx = QueryFormStruct(tx, fmt.Sprintf("%s->'%s'", qk, k), v.StructValue)
			continue
		case *structpb.Value_ListValue:
			if len(v.ListValue.Values) == 0 {
				continue
			}
		}

		b, err := v.MarshalJSON()
		if err != nil {
			continue
		}
		tx = tx.Where(fmt.Sprintf("%s->'%s' = '%s'", qk, k, string(b)))
	}

	return tx
}