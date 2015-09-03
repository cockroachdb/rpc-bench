rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go)
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/golang/protobuf) codec
- [net/http](http://godoc.org/net/http) using protobufs on the wire

results (2015-09-03)
====================
Run with `go test -cpu 1,2,4 -timeout 5s -bench .`

| Bench                   | Iters | Speed         | Troughput   |
|-------------------------|-------|---------------|-------------|
| BenchmarkGRPC1K        	| 10000	|  104305 ns/op	|  19.63 MB/s |
| BenchmarkGRPC1K-2      	| 20000	|   91039 ns/op	|  22.50 MB/s |
| BenchmarkGRPC1K-4      	| 20000	|   63809 ns/op	|  32.10 MB/s |
| BenchmarkGRPC64K       	|  2000	|  800945 ns/op	| 163.65 MB/s |
| BenchmarkGRPC64K-2     	|  2000	|  654298 ns/op	| 200.32 MB/s |
| BenchmarkGRPC64K-4     	|  3000	|  608590 ns/op	| 215.37 MB/s |
| BenchmarkGobRPC1K      	| 20000	|   55133 ns/op	|  37.15 MB/s |
| BenchmarkGobRPC1K-2    	| 30000	|   48343 ns/op	|  42.36 MB/s |
| BenchmarkGobRPC1K-4    	| 50000	|   28264 ns/op	|  72.46 MB/s |
| BenchmarkGobRPC64K     	|  5000	|  254183 ns/op	| 515.66 MB/s |
| BenchmarkGobRPC64K-2   	| 10000	|  215709 ns/op	| 607.63 MB/s |
| BenchmarkGobRPC64K-4   	| 10000	|  197455 ns/op	| 663.80 MB/s |
| BenchmarkProtoRPC1K    	| 20000	|   57597 ns/op	|  35.56 MB/s |
| BenchmarkProtoRPC1K-2  	| 30000	|   45918 ns/op	|  44.60 MB/s |
| BenchmarkProtoRPC1K-4  	| 50000	|   28712 ns/op	|  71.33 MB/s |
| BenchmarkProtoRPC64K   	|  3000	|  364220 ns/op	| 359.87 MB/s |
| BenchmarkProtoRPC64K-2 	|  5000	|  287173 ns/op	| 456.42 MB/s |
| BenchmarkProtoRPC64K-4 	|  5000	|  290006 ns/op	| 451.96 MB/s |
| BenchmarkProtoHTTP1K   	| 20000	|   86921 ns/op	|  23.56 MB/s |
| BenchmarkProtoHTTP1K-2 	| 20000	|   67597 ns/op	|  30.30 MB/s |
| BenchmarkProtoHTTP1K-4 	| 20000	|   61049 ns/op	|  33.55 MB/s |
| BenchmarkProtoHTTP64K  	|  2000	|  810811 ns/op	| 161.66 MB/s |
| BenchmarkProtoHTTP64K-2	|  3000	|  651075 ns/op	| 201.32 MB/s |
| BenchmarkProtoHTTP64K-4	|  2000	|  645277 ns/op	| 203.12 MB/s |
