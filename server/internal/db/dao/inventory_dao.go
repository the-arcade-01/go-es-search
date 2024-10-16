package dao

import (
	"log"
	"sync"

	"github.com/the-arcade-01/go-es-search/server/internal/config"
	db "github.com/the-arcade-01/go-es-search/server/internal/models/db"
	"github.com/the-arcade-01/go-es-search/server/internal/utils"
)

type InventoryDao interface {
	CreateProduct(product *db.ProductDB) error
	GetProduct(productId int) (*db.ProductDB, error)
	UpdateProduct(product *db.ProductDB) error
	DeleteProduct(productId int) error
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
