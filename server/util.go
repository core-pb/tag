package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/core-pb/dt/query/v1"
	"github.com/uptrace/bun"
	"go.x2ox.com/sorbifolia/pyrokinesis"
	"google.golang.org/protobuf/types/known/structpb"
)

var errInvalidKey = errors.New("invalid key")

func isInvalidKey(key string) bool {
	b := pyrokinesis.String.ToBytes(key)

	switch {
	case len(b) < 3:
	case !utf8.Valid(b):
	case b[0] != ':':

	default:
		return false
	}
	return true
}

var ErrPageTooBig = errors.New("page too big")

func Pagination(tx *bun.SelectQuery, x *query.Pagination) *bun.SelectQuery {
	if x == nil {
		return tx
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

func defaultValue[T any](v *T) T {
	if v == nil {
		return *new(T)
	}
	return *v
}

func arrayContains[T comparable](arr1, arr2 []T) bool {
	for _, v2 := range arr2 {
		found := false
		for _, v1 := range arr1 {
			if v1 == v2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func arrayDifference[T comparable](arr1, arr2 []T) []T {
	var result []T
	for _, v1 := range arr1 {
		found := false
		for _, v2 := range arr2 {
			if v1 == v2 {
				found = true
				break
			}
		}
		if !found {
			result = append(result, v1)
		}
	}
	return result
}
