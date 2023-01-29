package domain

import (
	"context"
	"strconv"
)

type TransactionUsecase interface{
	Insert(c context.Context, transaction Transaction) (err error)
	FindById(c context.Context, id int) (res Response, err error)
}

type transactionUsecase struct{
	transactionRepo TransactionRepository
}

func NewTransactionUsecase(transRepo TransactionRepository) TransactionUsecase{
	return &transactionUsecase{transactionRepo: transRepo}
}

func (uc *transactionUsecase) Insert(c context.Context, transaction Transaction) (err error) {
	err = uc.transactionRepo.Save(c, transaction)

	return err
}

func (uc *transactionUsecase) FindById(c context.Context, id int) (res Response, err error) {
	transaction, err := uc.transactionRepo.GetById(c, id)
	if err != nil {
		return Response{}, err
	}

	prodId := strconv.Itoa(transaction.Product_id)
	product, err := uc.transactionRepo.GetProductById(c, prodId)

	if err != nil {
		return Response{}, err
	}

	res.Id = transaction.Id
	res.Product = product

	return res, nil
}