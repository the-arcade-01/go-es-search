package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/the-arcade-01/go-es-search/server/internal/db/dao"
	response "github.com/the-arcade-01/go-es-search/server/internal/models/response"
)

type APIService struct {
	inventoryDao dao.InventoryDao
}

func NewAPIService() *APIService {
	return &APIService{
		inventoryDao: dao.NewInventoryDao(),
	}
}

func (api APIService) Greet(w http.ResponseWriter, r *http.Request) {
	response.ResponseWithJson(w, http.StatusOK, "Hello, World")
}

func (api APIService) GetProducts(w http.ResponseWriter, r *http.Request) {
}

func (api APIService) PostProducts(w http.ResponseWriter, r *http.Request) {
}

func (api APIService) PutProducts(w http.ResponseWriter, r *http.Request) {
}

func (api APIService) GetProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	result, err := api.inventoryDao.GetProductById(r.Context(), productId)
	if err != nil {
		return
	}
	log.Println("[GetProductById]", result)
}

func (api APIService) Search(w http.ResponseWriter, r *http.Request) {
}
