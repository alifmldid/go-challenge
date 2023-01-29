package domain

import (
	"context"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type TransactionRepository interface{
	Save(c context.Context, transaction Transaction) (err error)
	GetById(c context.Context, id int) (res Transaction, err error)
	GetProductById(c context.Context, id string) (product Product, err error)
}

type transactionRepository struct{
	Conn *gorm.DB
}

func NewTransactionRepository(Conn *gorm.DB) TransactionRepository{
	return &transactionRepository{Conn}
}

func (repo *transactionRepository) Save(c context.Context, transaction Transaction) (err error){
	result := repo.Conn.Save(&transaction)

	return result.Error
}

func (repo *transactionRepository) GetById(c context.Context, id int) (res Transaction, err error) {
	result := repo.Conn.Where("id = ?", id).First(&res)

	return res, result.Error
}

func (repo *transactionRepository) GetProductById(c context.Context, prodId string) (product Product, err error){
	var data Data

	var client = &http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8000/product/"+prodId, nil)
	token := c.Value("token").(string)
	request.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return Product{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return Product{}, err
	}

	err = json.NewDecoder(response.Body).Decode(&data)

	product = data.Data

	return product, nil
}