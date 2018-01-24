default:
	go build handler.go

unit-test:
	go test -v ./runner/... ./lib/...

integration-test:
	docker run -d --name docker-consul-handler-int-test -p 8500:8500 consul:latest
	go test -v handler_test.go
	docker rm -f docker-consul-handler-int-test