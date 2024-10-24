

run:
	go run cmd/main.go

hello:
	echo "hellloooo"
build:
	go build -o ./cmd/uniclub_project ./cmd
buildrun:
	go build -o ./cmd/uniclub_project ./cmd && ./cmd/uniclub_project
testall:
	go test -v ./...
testadmin:
	 go test -v ./pkg/usecase/admin_test.go
testuser:
	 go test -v ./pkg/usecase/cart_test.go

# run all integrated test
testallint:
	 go test -v ./pkg/tests/tests_integrated/user_db_test.go

mockall:
	 mockgen -source=./pkg/usecase/interface/admin.go -destination=./pkg/tests/mocks/admin/mock_admin_usecase.go -package=mocks
	 mockgen -source=./pkg/repository/interface/category.go -destination=./pkg/tests/mocks/admin/mock_category_repository.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/admin.go -destination=./pkg/tests/mocks/admin/mock_admin_repository.go  -package=mocks
	 mockgen -source=./pkg/helper/interface/helper.go -destination=./pkg/tests/mocks/admin/mock_admin_helper.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/order.go -destination=./pkg/tests/mocks/admin/mock_order_repository.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/user.go -destination=./pkg/tests/mocks/admin/mock_user_repository.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/order.go -destination=./pkg/tests/mocks/user/mock_order_repository.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/user.go -destination=./pkg/tests/mocks/user/mock_user_repository.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/inventory.go -destination=./pkg/tests/mocks/user/mock_inv_repo.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/cart.go -destination=./pkg/tests/mocks/user/mock_cart_repo.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/offer.go -destination=./pkg/tests/mocks/user/mock_off_repo.go  -package=mocks
