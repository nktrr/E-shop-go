package repository

import (
	"E-shop-go/pkg/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func (r *ProductPostgres) GetAll(page int, size int) ([]entity.ProductInfo, error) {
	offset := size * (page - 1)
	items := []entity.ProductInfo{}
	query := fmt.Sprintf(`SELECT * FROM product_info ORDER BY "product_id" LIMIT %d OFFSET %d`, size, offset)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}
	return items, nil
}

func NewProductInfoPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}
