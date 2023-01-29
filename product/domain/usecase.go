package domain

import "context"

type ProductUsecase interface{
	Insert(c context.Context, product Product) (err error)
	FindById(c context.Context, id int) (res Product, err error)
}

type productUsecase struct{
	productRepo ProductRepository
}

func NewProductUsecase(productRepo ProductRepository) ProductUsecase{
	return &productUsecase{productRepo: productRepo}
}

func (uc *productUsecase) Insert(c context.Context, product Product) (err error) {
	err = uc.productRepo.Save(c, product)
	
	return err
}

func (uc *productUsecase) FindById(c context.Context, id int) (res Product, err error){
	res, err = uc.productRepo.GetById(c, id)

	return res, err
}