package domain

import (
	"context"

	"gorm.io/gorm"
)

type ProductRepository interface{
	Save(c context.Context, product Product) (err error)
	GetById(c context.Context, id int) (res Product, err error)
}

type productRepository struct{
	Conn *gorm.DB 
}

func NewProductRepository(Conn *gorm.DB) ProductRepository {
	return &productRepository{Conn}
}

func (repo *productRepository) Save(c context.Context, product Product) (err error){
	result := repo.Conn.Save(&product)

	return result.Error
}

func (repo *productRepository) GetById(c context.Context, id int) (res Product, err error){
	result := repo.Conn.Where("id = ?", id).First(&res)

	return res, result.Error 
}