# Docker Consul Handler
This is a Consul Handler that addresses Docker Networks.

It retrieves the newest network created and returns
its ID.

Running:
```
docker build -t consul-handler:latest .
GOOS=linux GOARCH=amd64 go build handler.go
docker run -d --name testing -p 8500:8500 consul-handler:latest
docker exec testing consul watch -type=keyprefix -prefix=docker/network/v1.0/endpoint /scripts/handler &
docker exec testing cat /scripts/consul_watch.log
docker exec testing cat /scripts/consul_watch.log
```
