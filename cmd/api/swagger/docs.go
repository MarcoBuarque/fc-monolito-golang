// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/client-management": {
            "post": {
                "description": "This endpoint allows you to create a new client by providing its details in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clients"
                ],
                "summary": "Create a new cleint.",
                "parameters": [
                    {
                        "description": "Client data",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_internal_modules_client_adm_repository.Client"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-github_com_MarcoBuarque_monolito_internal_modules_client_adm_repository_Client"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null"
                        }
                    }
                }
            }
        },
        "/client-management/{clientID}": {
            "get": {
                "description": "This endpoint retrieves the details of a single cleint based on the provided cleint ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clients"
                ],
                "summary": "Retrieve a specific cleint by its ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cleint ID",
                        "name": "clientID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-github_com_MarcoBuarque_monolito_internal_modules_client_adm_repository_Client"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null"
                        }
                    }
                }
            }
        },
        "/product-management": {
            "post": {
                "description": "This endpoint allows you to create a new product by providing its details in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Create a new product.",
                "parameters": [
                    {
                        "description": "Product data",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_internal_modules_product_adm_repository.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-github_com_MarcoBuarque_monolito_internal_modules_product_adm_repository_Product"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null"
                        }
                    }
                }
            }
        },
        "/product-management/{productID}": {
            "patch": {
                "description": "This API allows you to update the sales price of a specific product by providing its ID and the new price in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Update the sales price of a product.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "productID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "new price value",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_server_routes_v1.UpdateSalesPriceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-github_com_MarcoBuarque_monolito_internal_modules_product_adm_repository_Product"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null"
                        }
                    }
                }
            }
        },
        "/store-catalog/products/": {
            "get": {
                "description": "This endpoint returns a list of products with information such as name, price, description, and status.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "catalogs"
                ],
                "summary": "Retrieve a list of active products.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-array_github_com_MarcoBuarque_monolito_internal_modules_store_catalog_repository_ProductCatalog"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null"
                        }
                    }
                }
            }
        },
        "/store-catalog/products/{productID}": {
            "get": {
                "description": "This endpoint retrieves the details of a single product based on the provided product ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "catalogs"
                ],
                "summary": "Retrieve a specific product by its ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "productID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-github_com_MarcoBuarque_monolito_internal_modules_store_catalog_repository_ProductCatalog"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_MarcoBuarque_monolito_internal_modules_client_adm_repository.Client": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "complement": {
                    "type": "string"
                },
                "documentNumber": {
                    "type": "string"
                },
                "documentType": {
                    "description": "Document",
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_MarcoBuarque_monolito_internal_modules_shared_types.DocumentType"
                        }
                    ]
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "street": {
                    "description": "Address",
                    "type": "string"
                },
                "zipCode": {
                    "type": "string"
                }
            }
        },
        "github_com_MarcoBuarque_monolito_internal_modules_product_adm_repository.Product": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "purchase_price": {
                    "type": "number"
                },
                "sales_price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "github_com_MarcoBuarque_monolito_internal_modules_shared_types.DocumentType": {
            "type": "string",
            "enum": [
                "RG",
                "CPF",
                "CNPJ"
            ],
            "x-enum-varnames": [
                "RG",
                "CPF",
                "CNPJ"
            ]
        },
        "github_com_MarcoBuarque_monolito_internal_modules_store_catalog_repository.ProductCatalog": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "salesPrice": {
                    "type": "number"
                }
            }
        },
        "github_com_MarcoBuarque_monolito_pkg.ApiResponse-array_github_com_MarcoBuarque_monolito_internal_modules_store_catalog_repository_ProductCatalog": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_MarcoBuarque_monolito_internal_modules_store_catalog_repository.ProductCatalog"
                    }
                },
                "response_key": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                }
            }
        },
        "github_com_MarcoBuarque_monolito_pkg.ApiResponse-github_com_MarcoBuarque_monolito_internal_modules_client_adm_repository_Client": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/github_com_MarcoBuarque_monolito_internal_modules_client_adm_repository.Client"
                },
                "response_key": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                }
            }
        },
        "github_com_MarcoBuarque_monolito_pkg.ApiResponse-github_com_MarcoBuarque_monolito_internal_modules_product_adm_repository_Product": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/github_com_MarcoBuarque_monolito_internal_modules_product_adm_repository.Product"
                },
                "response_key": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                }
            }
        },
        "github_com_MarcoBuarque_monolito_pkg.ApiResponse-github_com_MarcoBuarque_monolito_internal_modules_store_catalog_repository_ProductCatalog": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/github_com_MarcoBuarque_monolito_internal_modules_store_catalog_repository.ProductCatalog"
                },
                "response_key": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                }
            }
        },
        "github_com_MarcoBuarque_monolito_pkg.ApiResponse-pkg_Null": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/pkg.Null"
                },
                "response_key": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                }
            }
        },
        "internal_server_routes_v1.UpdateSalesPriceRequest": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "number"
                }
            }
        },
        "pkg.Null": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "fc-monolito APIs",
	Description:      "Implementation in golang of FC monolith course.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
