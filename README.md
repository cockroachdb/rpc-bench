rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go)
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/golang/protobuf) codec
- [net/http](http://godoc.org/net/http) using protobufs on the wire

results (2015-06-02)
====================
Run with `go test -cpu 1,2,4 -timeout 5s -bench .`

| Bench                   | Iters | Speed         | Troughput   |
|-------------------------|-------|---------------|-------------|
| BenchmarkGRPC1K         | 10000 | 149654 ns/op  | 13.68 MB/s  |
| BenchmarkGRPC1K-2       | 20000 | 84011 ns/op   | 24.38 MB/s  |
| BenchmarkGRPC1K-4       | 30000 | 61769 ns/op   | 33.16 MB/s  |
| BenchmarkGRPC64K        | 2000  | 867688 ns/op  | 151.06 MB/s |
| BenchmarkGRPC64K-2      | 3000  | 549554 ns/op  | 238.51 MB/s |
| BenchmarkGRPC64K-4      | 3000  | 530154 ns/op  | 247.23 MB/s |
| BenchmarkGobRPC1K       | 20000 | 54553 ns/op   | 37.54 MB/s  |
| BenchmarkGobRPC1K-2     | 30000 | 46111 ns/op   | 44.41 MB/s  |
| BenchmarkGobRPC1K-4     | 50000 | 27471 ns/op   | 74.55 MB/s  |
| BenchmarkGobRPC64K      | 5000  | 307457 ns/op  | 426.31 MB/s |
| BenchmarkGobRPC64K-2    | 10000 | 196209 ns/op  | 668.02 MB/s |
| BenchmarkGobRPC64K-4    | 10000 | 185178 ns/op  | 707.82 MB/s |
| BenchmarkProtoRPC1K     | 30000 | 54848 ns/op   | 37.34 MB/s  |
| BenchmarkProtoRPC1K-2   | 30000 | 44624 ns/op   | 45.89 MB/s  |
| BenchmarkProtoRPC1K-4   | 50000 | 29461 ns/op   | 69.51 MB/s  |
| BenchmarkProtoRPC64K    | 3000  | 493156 ns/op  | 265.78 MB/s |
| BenchmarkProtoRPC64K-2  | 5000  | 381305 ns/op  | 343.75 MB/s |
| BenchmarkProtoRPC64K-4  | 5000  | 361696 ns/op  | 362.38 MB/s |
| BenchmarkProtoHTTP1K    | 20000 | 84615 ns/op   | 24.20 MB/s  |
| BenchmarkProtoHTTP1K-2  | 20000 | 70180 ns/op   | 29.18 MB/s  |
| BenchmarkProtoHTTP1K-4  | 20000 | 62865 ns/op   | 32.58 MB/s  |
| BenchmarkProtoHTTP64K   | 1000  | 1110442 ns/op | 118.04 MB/s |
| BenchmarkProtoHTTP64K-2 | 2000  | 662140 ns/op  | 197.95 MB/s |
| BenchmarkProtoHTTP64K-4 | 2000  | 677619 ns/op  | 193.43 MB/s |
