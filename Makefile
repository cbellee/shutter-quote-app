generate:
	oapi-codegen --package=api --generate types ./api/swagger.yml > ./api/store-types.go
	oapi-codegen --package=api --generate server,spec ./api/swagger.yml > ./api/store-server.go

addQuote:
	curl http://localhost:8080/quote -H "Content-Type: application/json" -X POST -d '{"customerId": 2000,"windows":[{"height":2100,"material":"wood","name":"","notes":"some notes...","panel":2,"price":1939.50,"width":4200}]}'
	
addCustomer:
	curl http://localhost:8080/customer -H "Content-Type: application/json" -X POST -d '{"email":"sam.tam@home.net","firstname":"Sam","lastname":"Tam","phone": "0404373373", "address": {"city":"Adelaide","phone":"0446123999","postcode":"5000","suburb":"Dulwich"}}'
	