dep:
	dep ensure -vendor-only

# Use this only for development
dev:
	go build -o cleanarch app/api/main.go
	./cleanarch

test:
	go test ./... -coverprofile cp.out
	go tool cover -html=cp.out