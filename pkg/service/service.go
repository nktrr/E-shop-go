package service

import (
	"E-shop-go/pkg/entity"
	"E-shop-go/pkg/repository"
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

type Service struct {
	Cart
	OrderMain
	ProductCategory
	ProductInOrder
	ProductInfo
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		ProductInfo: NewProductInfoService(repos.ProductInfo),
	}
}
