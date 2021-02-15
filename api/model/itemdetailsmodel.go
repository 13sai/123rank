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
	itemDetailsFieldNames          = builderx.RawFieldNames(&ItemDetails{})
	itemDetailsRows                = strings.Join(itemDetailsFieldNames, ",")
	itemDetailsRowsExpectAutoSet   = strings.Join(stringx.Remove(itemDetailsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	itemDetailsRowsWithPlaceHolder = strings.Join(stringx.Remove(itemDetailsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	ItemDetailsModel interface {
		Insert(data ItemDetails) (sql.Result, error)
		FindOne(id int64) (*ItemDetails, error)
		Update(data ItemDetails) error
		Delete(id int64) error
		GetAll(itemId int64) ([]ItemDetails, error)
	}

	defaultItemDetailsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ItemDetails struct {
		Id        int64     `db:"id" json:"id"`
		ItemId    int64     `db:"item_id" json:"item_id"`
		Title     string    `db:"title" json:"title"`
		CreatedAt time.Time `db:"created_at" json:"created_at"`
	}
)

func NewItemDetailsModel(conn sqlx.SqlConn) ItemDetailsModel {
	return &defaultItemDetailsModel{
		conn:  conn,
		table: "`item_details`",
	}
}

func (m *defaultItemDetailsModel) Insert(data ItemDetails) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, itemDetailsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ItemId, data.Title, data.CreatedAt)
	return ret, err
}

func (m *defaultItemDetailsModel) FindOne(id int64) (*ItemDetails, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", itemDetailsRows, m.table)
	var resp ItemDetails
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

func (m *defaultItemDetailsModel) Update(data ItemDetails) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, itemDetailsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ItemId, data.Title, data.CreatedAt, data.Id)
	return err
}

func (m *defaultItemDetailsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultItemDetailsModel) GetAll(itemId int64) ([]ItemDetails, error) {
	query := fmt.Sprintf("select %s from %s where `item_id` = %d", itemDetailsRows, m.table, itemId)
	fmt.Println(query)
	var resp []ItemDetails
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
