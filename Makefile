generate:
	oapi-codegen -generate types ./api/swagger-v3.yml > ./api/quoteapp.types.go
	oapi-codegen -generate server ./api/swagger-v3.yml > ./api/quoteapp.server.go
	oapi-codegen -generate spec ./api/swagger-v3.yml > ./api/quoteapp.spec.go