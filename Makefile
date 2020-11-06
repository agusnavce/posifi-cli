build:
	@clear
	@go run cmd/client/*.go

mock:
	@clear
	@go run cmd/build/*.go
	@./posifi-cli -M true

run:
	@clear
	@go run cmd/build/*.go
	@./posifi-cli