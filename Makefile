default:
	GOOS=linux GOARCH=amd64 go build handler.go

unit-test:
	go test ./runner/... ./lib/...