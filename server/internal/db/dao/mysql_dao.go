package dao

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/the-arcade-01/go-es-search/server/internal/config"
	"github.com/the-arcade-01/go-es-search/server/internal/db/mysql"
	"github.com/the-arcade-01/go-es-search/server/internal/models/db"
	"github.com/the-arcade-01/go-es-search/server/internal/utils"
)

type MysqlDao struct {
	mysqlClient *sql.DB
}

func NewMysqlDao() *MysqlDao {
	return &MysqlDao{
		mysqlClient: mysql.NewMysqlConnection(config.FetchAndLoadConfig()),
	}
}

func (dao *MysqlDao) CreateProduct(ctx context.Context, product *db.ProductDB) (string, error) {
	result, err := dao.mysqlClient.ExecContext(ctx, utils.DB_CREATE_PRODUCT, product.Name, product.Name, product.Title, product.Description, product.Price, product.Quantity)
	if err != nil {
		return "", err
	}
	id := fmt.Sprint(result.LastInsertId())
	return id, nil
}

func (dao *MysqlDao) GetProductById(ctx context.Context, productId int) (*db.ProductDB, error) {
	product := new(db.ProductDB)
	err := dao.mysqlClient.QueryRowContext(ctx, utils.DB_GET_PRODUCT, productId).Scan(&product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (dao *MysqlDao) DeleteProduct(ctx context.Context, productId int) error {
	_, err := dao.mysqlClient.ExecContext(ctx, utils.DB_DELETE_PRODUCT, productId)
	if err != nil {
		return err
	}
	return nil
}
