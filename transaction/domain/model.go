package domain

type Transaction struct{
	Id int `json:"id"`
	Product_id int `json:"product_id"`
}

type Response struct{
	Id int `json:"id"`
	Product Product `json:"product"`
}

type Data struct{
	Data Product `json:"data"`
	Message string `json:"message"`
}

type Product struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Stock int `json:"stock"`
}