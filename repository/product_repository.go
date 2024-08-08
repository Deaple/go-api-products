package repository

import (
	"api-produtos/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	query, err := pr.connection.Prepare("INSERT INTO product(product_name, price)" +
		" VALUES($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var id int

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT ID, product_name, price from product"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	//retorna sem erros
	return productList, err
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT id, product_name, price" +
		" FROM product where id=$1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)

		return nil, err
	}

	query.Close()

	return &product, nil
}

func (pr *ProductRepository) DeleteProductById(id int) (bool, error) {
	rows, err := pr.connection.Exec("DELETE FROM product where id=$1", id)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	count, err := rows.RowsAffected()

	if err != nil {
		//a query executou com sucesso mas não teve linhas
		if err == sql.ErrNoRows {
			return true, nil
		}
		fmt.Println(err)

		return false, err
	}

	//verifica se há ou não linhas afetadas
	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) (bool, error) {
	var (
		rows sql.Result
		err  error
	)

	fmt.Println(product.ID)

	if product.Name != "" && product.Price > 0 {
		rows, err = pr.connection.Exec("UPDATE product SET product_name=$2, price=$3 where id=$1", product.ID, product.Name, product.Price)
	} else if product.Name != "" && product.Price == 0 {
		rows, err = pr.connection.Exec("UPDATE product SET product_name=$2 where id=$1", product.ID, product.Name)
	} else if product.Name == "" && product.Price > 0 {
		rows, err = pr.connection.Exec("UPDATE product SET price=$2 where id=$1", product.ID, product.Price)
	} else {
		return false, nil
	}

	if err != nil {
		fmt.Println(err)
		return true, err
	}

	count, err := rows.RowsAffected()

	if err != nil {
		//a query executou com sucesso mas não teve linhas
		if err == sql.ErrNoRows {
			return true, nil
		}
		fmt.Println(err)

		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}
