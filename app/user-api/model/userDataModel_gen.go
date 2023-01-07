// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userDataFieldNames          = builder.RawFieldNames(&UserData{})
	userDataRows                = strings.Join(userDataFieldNames, ",")
	userDataRowsExpectAutoSet   = strings.Join(stringx.Remove(userDataFieldNames, "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	userDataRowsWithPlaceHolder = strings.Join(stringx.Remove(userDataFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"

	cacheUserDataIdPrefix     = "cache:userData:id:"
	cacheUserDataUserIdPrefix = "cache:userData:userId:"
)

type (
	userDataModel interface {
		Insert(ctx context.Context, data *UserData) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *UserData) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserData, error)
		FindOneByUserId(ctx context.Context, userId int64) (*UserData, error)
		Update(ctx context.Context, data *UserData) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserDataModel struct {
		sqlc.CachedConn
		table string
	}

	UserData struct {
		Id         int64     `db:"id"`
		UserId     int64     `db:"user_id"`
		Data       string    `db:"data"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newUserDataModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserDataModel {
	return &defaultUserDataModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_data`",
	}
}

func (m *defaultUserDataModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userDataIdKey := fmt.Sprintf("%s%v", cacheUserDataIdPrefix, id)
	userDataUserIdKey := fmt.Sprintf("%s%v", cacheUserDataUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userDataIdKey, userDataUserIdKey)
	return err
}

func (m *defaultUserDataModel) FindOne(ctx context.Context, id int64) (*UserData, error) {
	userDataIdKey := fmt.Sprintf("%s%v", cacheUserDataIdPrefix, id)
	var resp UserData
	err := m.QueryRowCtx(ctx, &resp, userDataIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userDataRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserDataModel) FindOneByUserId(ctx context.Context, userId int64) (*UserData, error) {
	userDataUserIdKey := fmt.Sprintf("%s%v", cacheUserDataUserIdPrefix, userId)
	var resp UserData
	err := m.QueryRowIndexCtx(ctx, &resp, userDataUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userDataRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserDataModel) Insert(ctx context.Context, data *UserData) (sql.Result, error) {
	userDataIdKey := fmt.Sprintf("%s%v", cacheUserDataIdPrefix, data.Id)
	userDataUserIdKey := fmt.Sprintf("%s%v", cacheUserDataUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userDataRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.Data)
	}, userDataIdKey, userDataUserIdKey)
	return ret, err
}

func (m *defaultUserDataModel) TransInsert(ctx context.Context, session sqlx.Session, data *UserData) (sql.Result, error) {
	userDataIdKey := fmt.Sprintf("%s%v", cacheUserDataIdPrefix, data.Id)
	userDataUserIdKey := fmt.Sprintf("%s%v", cacheUserDataUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userDataRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.Id, data.UserId, data.Data)
	}, userDataIdKey, userDataUserIdKey)
	return ret, err
}

func (m *defaultUserDataModel) Update(ctx context.Context, newData *UserData) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	userDataIdKey := fmt.Sprintf("%s%v", cacheUserDataIdPrefix, data.Id)
	userDataUserIdKey := fmt.Sprintf("%s%v", cacheUserDataUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userDataRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.Data, newData.Id)
	}, userDataIdKey, userDataUserIdKey)
	return err
}

func (m *defaultUserDataModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserDataIdPrefix, primary)
}

func (m *defaultUserDataModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userDataRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserDataModel) tableName() string {
	return m.table
}
