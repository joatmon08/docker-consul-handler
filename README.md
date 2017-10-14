# Docker Consul Handler
This is a handler written to pass a Docker network name to an Ansible Playbook
runner application. Use this as an example to create more sophisticated
Consul handlers.

1. When you configure Consul, set a watch on the docker/network/v1.0/endpoint
keyprefix. See config.json for an example.
1. Create a Docker network on the host.
1. The handler receives a keyprefix output from Consul in the format:
    ```
    [{"Key":"docker/network/v1.0/endpoint/264cc0919c0041156f5273c8e12d1a0541663c3372506b9c5397059acf4ee10a/","CreateIndex":12,"ModifyIndex":12,"LockIndex":0,"Flags":3304740253564472344,"Value":null,"Session":""}]
    ```
1. The handler examines the ModifyIndex for the newly modified network.
1. It parses out the Docker network ID from the Key and retrieves the network's
name from Docker.
1. It passes the network name to an external API set by CONSUL_HANDLER_RUNNER.
This external API happens to be a runner that executes an Ansible Playbook.

## Run
```
CONSUL_HANDLER_RUNNER=http://<external> ./handler
```

## Build
```
GOOS=linux GOARCH=amd64 go build handler.go
```

## Test
Some of the tests depend on Docker or Consul. For the most part, if you
want to see if the modifications you made worked, you can simply start the
handler via:
```
CONSUL_HANDLER_RUNNER=http://<external> go run handler.go
```

Pass the KeyPrefix output sample above (or in nospaces.json) to test
if the behavior is as expected.
