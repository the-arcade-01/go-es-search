package dao

import (
	"database/sql"

	"github.com/the-arcade-01/go-es-search/server/internal/config"
	"github.com/the-arcade-01/go-es-search/server/internal/db/mysql"
	"github.com/the-arcade-01/go-es-search/server/internal/models/db"
)

type MysqlDao struct {
	mysqlClient *sql.DB
}

func NewMysqlDao() *MysqlDao {
	return &MysqlDao{
		mysqlClient: mysql.NewMysqlConnection(config.FetchAndLoadConfig()),
	}
}

func (dao *MysqlDao) CreateProduct(product *db.ProductDB) error {
	return nil
}

func (dao *MysqlDao) GetProduct(productId int) (*db.ProductDB, error) {
	return nil, nil
}

func (dao *MysqlDao) UpdateProduct(product *db.ProductDB) error {
	return nil
}

func (dao *MysqlDao) DeleteProduct(productId int) error {
	return nil
}
