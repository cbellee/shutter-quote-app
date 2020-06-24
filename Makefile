generate:
	oapi-codegen --package=api --generate types ./api/swagger.yml > ./api/store/quotestore-types.go
	oapi-codegen --package=api --generate server,spec ./api/swagger.yml > ./api/store/quotestore-server.go
	#go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate types -o petstore-types.gen.go ../../petstore-expanded.yaml
	#go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate server,spec -o petstore-server.gen.go ../../petstore-expanded.yaml