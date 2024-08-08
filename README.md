# About
This project is a RESTful API CRUD (Create, Read, Update, Delete) application built in Go 1.22, using the Gin web framework for handling HTTP requests and PostgreSQL as the relational database for data storage. The application is containerized using Docker, and a docker-compose setup is provided to facilitate easy deployment and management of the application and its database in development enviroment.

## Main activities of this API
- Create a product
- Delete a product
- Update a product
- List one specific or all products

## Usage of the API
### Endpoint: POST /product
This operation allows the creation of a new product in the database. The product details are sent as a JSON payload, and the API responds with the newly created product or an error if the creation fails.
```bash
curl -X POST http://localhost:8080/product \
-H "Content-Type: application/json" \
-d '{
    "name": "Salada",
    "price": 15.80
}'
```

### Endpoint: GET /products
Retrieves a list of all products available in the database. The API responds with the list of products or an error if the retrieval fails.
```bash
curl -X GET http://localhost:8080/products
```

### Endpoint: GET /product/:id
Fetches a specific product by its ID. The ID is passed as a URL parameter, and the API returns the corresponding product details or an error if the product is not found.
```bash
curl -X GET http://localhost:8080/product/1
```

### Endpoint: PUT /product
Updates the details of an existing product. The updated product details are sent as a JSON payload. The API checks if the product exists and updates it accordingly, returning a success message or an error if the product is not found.
```bash
curl -X PUT http://localhost:8080/product \
-H "Content-Type: application/json" \
-d '{
    "id_product": 1,
    "name": "Salada Atualizada",
    "price": 18.00
}'
```

### Endpoint: DELETE /product/:id
Deletes a product from the database by its ID. The API returns a success message if the deletion is successful, or an error if the product is not found.
```bash
curl -X DELETE http://localhost:8080/product/1
```

## Database/Repository layer
The application uses PostgreSQL as the database to store product information. The database schema includes a "product" table with the following fields:
- id: Primary key, integer and auto-increment
- product_name: Product name, string.
- price: Product price, float.

The database connection is handled within the application, ensuring that CRUD operations are performed efficiently.

## TODOs: 
- implement authorization and authentication using JWT and OAuth 2.0
- self-document the API using Open API 3.1 specification
- implement unit and integration tests to improve quality
- implement CI/CD pipeline, to automate the build, tests run and docker image build and push to central container repository

## How to run
First, you need to build the image and start the database and (optionally) the GUI container for managing the database. You can use docker-compose (or podman-compose).

#### build the image with
```bash
docker build -t go-api-products:1.0
```

#### run docker-compose 
```bash
docker-compose up -d
```

#### to start only the database
```bash
docker-compose up -d go_db
```

To test locally, before building the docker image, you need to **download the modules** (also, start the database like above):
```bash
go mod download
```

#### start the app
```bash
go run cmd/main.go
```
