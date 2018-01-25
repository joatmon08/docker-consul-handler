default:
	go build handler.go

unit-test:
	go test -v ./runner/... ./lib/...

contract-test:
	go test -v handler_test.go

clean:
	rm -f handler

build-amd64:
	GOOS=linux GOARCH=amd64 go build handler.go