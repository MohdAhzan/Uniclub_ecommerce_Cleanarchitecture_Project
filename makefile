

run:
	go run cmd/main.go

hello:
	echo "hellloooo"

testall:
	go test -v ./...
mock:
	 mockgen -source=./pkg/usecase/interface/admin.go -destination=./pkg/mocks/admin/mock_admin_usecase.go -package=mocks
	 mockgen -source=./pkg/repository/interface/admin.go -destination=./pkg/mocks/admin/mock_admin_repository.go  -package=mocks
	 mockgen -source=./pkg/helper/interface/helper.go -destination=./pkg/mocks/admin/mock_admin_helper.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/order.go -destination=./pkg/mocks/admin/mock_order_repository.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/user.go -destination=./pkg/mocks/admin/mock_user_repository.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/inventory.go -destination=./pkg/mocks/admin/mock_inv_repo.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/cart.go -destination=./pkg/mocks/admin/mock_cart_repo.go  -package=mocks
	 mockgen -source=./pkg/repository/interface/offer.go -destination=./pkg/mocks/admin/mock_off_repo.go  -package=mocks