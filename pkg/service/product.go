package service

import (
	"E-shop-go/pkg/entity"
	"E-shop-go/pkg/repository"
)

type ProductInfoService struct {
	repo repository.ProductInfo
}

func (s ProductInfoService) GetAll(page int, size int) ([]entity.ProductInfo, error) {
	return s.repo.GetAll(page, size)

}

func NewProductInfoService(repo repository.ProductInfo) *ProductInfoService {
	return &ProductInfoService{repo: repo}
}
