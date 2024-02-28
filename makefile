migrate:
	migrate -database ${DB_POSTGRES_URL} -path ./migrations up

run-test:
	go test ./... -coverprofile=c.out
	go tool cover -html="c.out"

generate-mock:
	# product adm
	mockery --dir=internal/modules/product_adm/repository --name=IProductRepository --output=internal/modules/product_adm/mocks/repomocks  --outpkg=repomocks
	mockery --dir=internal/modules/product_adm/usecase/add_product --name=IAddProductUseCase --output=internal/modules/product_adm/mocks/usecasemocks  --outpkg=usecasemocks
	mockery --dir=internal/modules/product_adm/usecase/check_stock --name=ICheckStockUseCase --output=internal/modules/product_adm/mocks/usecasemocks  --outpkg=usecasemocks

	# store_catalog
	mockery --dir=internal/modules/store_catalog/repository --name=IProductRepository --output=internal/modules/store_catalog/mocks/repomocks  --outpkg=repomocks
	mockery --dir=internal/modules/store_catalog/usecase/find_all_products --name=IFindAllProductsUseCase --output=internal/modules/store_catalog/mocks/usecasemocks  --outpkg=usecasemocks
	mockery --dir=internal/modules/store_catalog/usecase/find_product --name=IFindProductUseCase --output=internal/modules/store_catalog/mocks/usecasemocks  --outpkg=usecasemocks
	mockery --dir=internal/modules/store_catalog/usecase/update_sales_price --name=IUpdateSalesPriceUseCase --output=internal/modules/store_catalog/mocks/usecasemocks  --outpkg=usecasemocks

	# client_adm
	mockery --dir=internal/modules/client_adm/repository --name=IClientRepository --output=internal/modules/client_adm/mocks/repomocks  --outpkg=repomocks
	mockery --dir=internal/modules/client_adm/usecase/add_client --name=IAddClientUseCase --output=internal/modules/client_adm/mocks/usecasemocks  --outpkg=usecasemocks
	mockery --dir=internal/modules/client_adm/usecase/find_client --name=IFindClientUseCase --output=internal/modules/client_adm/mocks/usecasemocks  --outpkg=usecasemocks

	# payment
	mockery --dir=internal/modules/payment/repository --name=ITransactionRepository --output=internal/modules/payment/mocks/repomocks  --outpkg=repomocks
	mockery --dir=internal/modules/payment/usecase/process_payment --name=IProcessPaymentUseCase --output=internal/modules/payment/mocks/usecasemocks  --outpkg=usecasemocks

	# checkout	
	mockery --dir=internal/modules/checkout/repository --name=ICheckoutRepository --output=internal/modules/checkout/mocks/repomocks  --outpkg=repomocks
	mockery --dir=internal/modules/checkout/usecase/place_order --name=IPlaceOrderUseCase --output=internal/modules/checkout/mocks/usecasemocks  --outpkg=usecasemocks	