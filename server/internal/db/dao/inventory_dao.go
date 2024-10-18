package dao

import (
	"context"
	"log"
	"sync"

	"github.com/the-arcade-01/go-es-search/server/internal/config"
	db "github.com/the-arcade-01/go-es-search/server/internal/models/db"
	"github.com/the-arcade-01/go-es-search/server/internal/utils"
)

type InventoryDao interface {
	CreateProduct(ctx context.Context, product *db.ProductDB) (string, error)
	GetProductById(ctx context.Context, productId int) (*db.ProductDB, error)
	DeleteProduct(ctx context.Context, productId int) error
}

var (
	once         sync.Once
	inventoryDao InventoryDao
)

func NewInventoryDao() InventoryDao {
	once.Do(func() {
		switch config.FetchAndLoadConfig().DbType {
		case utils.MYSQL:
			inventoryDao = NewMysqlDao()
		}
		log.Println("[NewInventoryDao] InventoryDao initialized")
	})
	return inventoryDao
}
