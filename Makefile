generate-mocks:
	mockgen -source repository/quote_repository_interface.go -destination repository/mock_quote_repository.go -package repository
	mockgen -source service/quote_service_interface.go -destination service/mock_quote_service.go -package service