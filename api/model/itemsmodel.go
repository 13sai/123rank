package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	itemsFieldNames          = builderx.RawFieldNames(&Items{})
	itemsRows                = strings.Join(itemsFieldNames, ",")
	itemsRowsExpectAutoSet   = strings.Join(stringx.Remove(itemsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	itemsRowsWithPlaceHolder = strings.Join(stringx.Remove(itemsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	ItemsModel interface {
		Insert(data Items) (sql.Result, error)
		FindOne(id int64) (*Items, error)
		Update(data Items) error
		Delete(id int64) error
		GetList(page, limit int) ([]Items, error)
	}

	defaultItemsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Items struct {
		Id        int64     `db:"id" json:"id"`
		Title     string    `db:"title" json:"title"`
		CreatedAt time.Time `db:"created_at" json:"created_at"`
		Sort      int64     `db:"sort" json:"sort"`
		Type      int64     `db:"type" json:"type"`
	}
)

func NewItemsModel(conn sqlx.SqlConn) ItemsModel {
	return &defaultItemsModel{
		conn:  conn,
		table: "`items`",
	}
}

func (m *defaultItemsModel) Insert(data Items) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, itemsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Title, data.CreatedAt, data.Sort, data.Type)
	return ret, err
}

func (m *defaultItemsModel) FindOne(id int64) (*Items, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", itemsRows, m.table)
	var resp Items
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultItemsModel) Update(data Items) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, itemsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Title, data.CreatedAt, data.Sort, data.Type, data.Id)
	return err
}

func (m *defaultItemsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultItemsModel) GetList(page, limit int) ([]Items, error) {
	query := fmt.Sprintf("select %s from %s limit %d offset %d ", itemsRows, m.table, page, limit*(page-1))
	var resp []Items
	err := m.conn.QueryRows(&resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
