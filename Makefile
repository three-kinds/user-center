test:
	go test -p 1 `go list ./... | grep -v _gen.go` -cover --coverprofile=cover.out
	go tool cover -html=cover.out
init_users:
	go build -ldflags "-s -w" -o target/init_users commands/init_users/command.go
	./target/init_users
