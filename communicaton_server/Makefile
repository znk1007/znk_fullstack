
## get package
packages:
	go mod edit -replace=google.golang.org/grpc=github.com/grpc/grpc-go@latest

## 创建cockroach证书
create_cockroach_certs:
	rm -rf ${HOME}/.cockroach_certs
	mkdir ${HOME}/.cockroach_certs
	cd ${HOME}/.cockroach_certs
	mkdir ${HOME}/.cockroach_certs/certs
	mkdir ${HOME}/.cockroach_certs/safe_certs
	cockroach cert create-ca --certs-dir=${HOME}/.cockroach_certs/certs --ca-key=${HOME}/.cockroach_certs/safe_certs/ca.key
	cockroach cert create-client root --certs-dir=${HOME}/.cockroach_certs/certs --ca-key=${HOME}/.cockroach_certs/safe_certs/ca.key
	cockroach cert create-node localhost $(hostname) --certs-dir=${HOME}/.cockroach_certs/certs --ca-key=${HOME}/.cockroach_certs/safe_certs/ca.key

start_insecure_cockroach_server:
	cockroach start --insecure --host=localhost
	cockroach start --insecure --store=node2 --host=localhost --port=26258 --http-port=8081 --join=localhost:26257
	cockroach start --insecure --store=node3 --host=localhost --port=26259 --http-port=8082 --join=localhost:26257

start_secure_cockroach_server:

