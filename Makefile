test:
	go test -p 1 `go list ./... | grep -v _gen.go` -cover --coverprofile=cover.out
	go tool cover -html=cover.out
