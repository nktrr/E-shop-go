package repository

import (
	"E-shop-go/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type Cart interface {
}

type OrderMain interface {
}

type ProductCategory interface {
}

type ProductInOrder interface {
}

type ProductInfo interface {
	GetAll(int, int) ([]entity.ProductInfo, error)
}

type Users interface {
}

type Repository struct {
	Cart
	OrderMain
	ProductCategory
	ProductInOrder
	ProductInfo
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ProductInfo: NewProductInfoPostgres(db),
	}
}
