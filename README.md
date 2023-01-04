## Code Template
==============

Just yet another dumb code boilerplate which which doesn't do anything useful but can be copy-pasted partly.
Included: 
* grpc server, helth check endpoints
* grpc->http gateway
* buildinfo(info about build as http handler. Do `make build` to make it fully work). http://localhost:8081/build 
* cobra commands & main ctx & gracefull shutdown
* prometheus metrics http://localhost:8081/metrics
* tracing/jaeger, traceID integration
* swagger docs(!). Imagine we have many micro-services and if each will serve own docs it'll be not so good idea(additional load,difficult to find). Better have centarlized swagger and each services will serve `swagger.json`. http://localhost:8081/swagger.json
* grpc panic handling & validation middleware (which must be in library somewhere)
* makefile/docker with multistage
* mocks & simple table-tests. GRPC client mocks for usage in other services(`mockgen.sh`)
* folder structure

Not included/ToDo:
- DI (dig/ google wire)
- k8s, helm
- ginkgo tests
- other cmd options
- maybe do something at least a little useful(now just abstractions)

### Run 

go run cmd/main.go

```
Usage:
  serve [flags]

Aliases:
  serve, s

Flags:
      --config string   path to config
  -h, --help            help for serve
```

Example: `go run cmd/main.go serve --config config.yaml` or just `make serve`

#### Makefile
make
* `build` - build app
* `serve` - launch app
* `test` - run all tests
* `generate` - generate mocks

### Docs 
Swagger.json(not GUI): http://localhost:8081/swagger.json

Prometheus metrics endpoint: http://localhost:8081/metrics

Server build info: http://localhost:8081/build