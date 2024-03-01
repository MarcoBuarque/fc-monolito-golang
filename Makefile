#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yml


#========================#
#== DATABASE MIGRATION ==#
#========================#

# migrate-up: ## Run migrations UP
migrate-up:
	docker-compose -f ${DOCKER_COMPOSE_FILE} run --rm migrate up

migrate-down: ## Rollback migrations against non test DB
migrate-down:
	docker-compose -f ${DOCKER_COMPOSE_FILE} run --rm migrate down 1

# migrate-create: ## Create a DB migration files e.g `make migrate-create name=migration-name`
migrate-create:
	docker-compose -f ${DOCKER_COMPOSE_FILE} run --rm migrate create -ext sql -dir /migrations $(name)

run-test:
	go test ./... -coverprofile=c.out
	go tool cover -html="c.out"

deploy-docker:
	docker build -t go-docker-image .
	docker compose up	

generate-swagger:
	swag init -g  ./cmd/api/main.go -o cmd/api/swagger --parseDependency --parseInternal 

generate-mock:
	# product adm
	mockery --dir=internal/modules/product_adm/repository --name=IProductRepository --output=internal/modules/product_adm/mocks/repomocks  --outpkg=repomocks
	mockery --dir=internal/modules/product_adm/usecase/create_product --name=ICreateProductUseCase --output=internal/modules/product_adm/mocks/usecasemocks  --outpkg=usecasemocks
	mockery --dir=internal/modules/product_adm/usecase/check_stock --name=ICheckStockUseCase --output=internal/modules/product_adm/mocks/usecasemocks  --outpkg=usecasemocks
	mockery --dir=internal/modules/product_adm/usecase/update_sales_price --name=IUpdateSalesPriceUseCase --output=internal/modules/product_adm/mocks/usecasemocks  --outpkg=usecasemocks

	# store_catalog
	mockery --dir=internal/modules/store_catalog/repository --name=IProductRepository --output=internal/modules/store_catalog/mocks/repomocks  --outpkg=repomocks
	mockery --dir=internal/modules/store_catalog/usecase/list_products --name=IListProductsUseCase --output=internal/modules/store_catalog/mocks/usecasemocks  --outpkg=usecasemocks
	mockery --dir=internal/modules/store_catalog/usecase/get_product --name=IGetProductUseCase --output=internal/modules/store_catalog/mocks/usecasemocks  --outpkg=usecasemocks

	# client_adm
	mockery --dir=internal/modules/client_adm/repository --name=IClientRepository --output=internal/modules/client_adm/mocks/repomocks  --outpkg=repomocks
	mockery --dir=internal/modules/client_adm/usecase/create_client --name=ICreateClientUseCase --output=internal/modules/client_adm/mocks/usecasemocks  --outpkg=usecasemocks
	mockery --dir=internal/modules/client_adm/usecase/get_client --name=IGetClientUseCase --output=internal/modules/client_adm/mocks/usecasemocks  --outpkg=usecasemocks

	# payment
	mockery --dir=internal/modules/payment/repository --name=ITransactionRepository --output=internal/modules/payment/mocks/repomocks  --outpkg=repomocks
	mockery --dir=internal/modules/payment/usecase/process_payment --name=IProcessPaymentUseCase --output=internal/modules/payment/mocks/usecasemocks  --outpkg=usecasemocks

	# checkout	
	mockery --dir=internal/modules/checkout/repository --name=ICheckoutRepository --output=internal/modules/checkout/mocks/repomocks  --outpkg=repomocks
	mockery --dir=internal/modules/checkout/usecase/place_order --name=IPlaceOrderUseCase --output=internal/modules/checkout/mocks/usecasemocks  --outpkg=usecasemocks	