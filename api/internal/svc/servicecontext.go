package svc

import (
	"api/internal/config"
	"api/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	ItemModel model.ItemsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	dao := model.NewItemsModel(conn)

	return &ServiceContext{
		Config:    c,
		ItemModel: dao,
	}
}
