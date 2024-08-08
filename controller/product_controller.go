package controller

import (
	"api-produtos/model"
	usecase "api-produtos/use_case"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	//tenta alocar o conteúdo recebido no JSON dentro
	//de uma instância do objeto model
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProducts(ctx *gin.Context) {
	// Mock
	// products := []model.Product{
	// 	{
	// 		ID:    1,
	// 		Name:  "Batata frita",
	// 		Price: 10,
	// 	},
	// }

	products, err := p.productUseCase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) GetProducById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := model.Response{Message: "id não pode ser vazio"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	//tenta converter o id em inteiro
	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "id do produto deve ser um número inteiro",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	product, err := p.productUseCase.GetProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não encontrado pelo id informado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)

}

func (p *productController) DeleteProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := model.Response{Message: "id não pode ser vazio"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	//tenta converter o id em inteiro
	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "id do produto deve ser um número inteiro",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	status, err := p.productUseCase.DeleteProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if !status {
		response := model.Response{
			Message: "Produto não encontrado.",
		}

		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := model.Response{
		Message: "Produto deletado com sucesso.",
	}

	ctx.JSON(http.StatusAccepted, response)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {

	var product model.Product
	//tenta alocar o conteúdo recebido no JSON dentro
	//de uma instância do objeto model
	err := ctx.BindJSON(&product)

	fmt.Println(product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	status, err := p.productUseCase.UpdateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if !status {
		response := model.Response{
			Message: "Produto não encontrado.",
		}

		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := model.Response{
		Message: "Produto atualizado com sucesso.",
	}

	ctx.JSON(http.StatusAccepted, response)

}
