package main

import (
	"api-produtos/controller"
	"api-produtos/db"
	"api-produtos/repository"
	usecase "api-produtos/use_case"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada use case
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)
	//Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.POST("/product", ProductController.CreateProduct)
	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:id", ProductController.GetProducById)
	server.PUT("/product", ProductController.UpdateProduct)
	server.DELETE("/product/:id", ProductController.DeleteProductById)
	server.Run(":8080")
}
