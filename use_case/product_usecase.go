package usecase

import (
	"api-produtos/model"
	"api-produtos/repository"
	"fmt"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)

	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	// complementa o objeto parcial (name e price)
	// com o ID gerado pelo DB e retorna no model
	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id)

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) DeleteProductById(id int) (bool, error) {
	status, err := pu.repository.DeleteProductById(id)

	if err != nil {
		fmt.Println(err)

		return false, err
	}

	if !status {
		return false, nil
	}

	return true, nil
}

func (pu *ProductUsecase) UpdateProduct(product model.Product) (bool, error) {
	status, err := pu.repository.UpdateProduct(product)

	if err != nil {
		fmt.Println(err)

		return false, err
	}

	if !status {
		return false, nil
	}

	return true, nil
}
