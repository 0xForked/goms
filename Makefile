proto:
	protoc --go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:.  \
        ./pkg/pb/store.proto \
        ./pkg/pb/book.proto
