rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go)
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/golang/protobuf) codec
- [net/http](http://godoc.org/net/http) using protobufs on the wire

results (2015-09-11)
====================
Run with `go test -cpu 1,2,4 -timeout 5s -bench .`

| Bench                   | Iters | Speed         | Troughput   |
|-------------------------|-------|---------------|-------------|
|BenchmarkGRPC1K          | 20000 |   95072 ns/op |  21.54 MB/s |
|BenchmarkGRPC1K-2        | 20000 |   76156 ns/op |  26.89 MB/s |
|BenchmarkGRPC1K-4        | 30000 |   59960 ns/op |  34.16 MB/s |
|BenchmarkGRPC64K         |  2000 |  752087 ns/op | 174.28 MB/s |
|BenchmarkGRPC64K-2       |  3000 |  615625 ns/op | 212.91 MB/s |
|BenchmarkGRPC64K-4       |  3000 |  589163 ns/op | 222.47 MB/s |
|BenchmarkGobRPC1K        | 30000 |   49018 ns/op |  41.78 MB/s |
|BenchmarkGobRPC1K-2      | 30000 |   44729 ns/op |  45.79 MB/s |
|BenchmarkGobRPC1K-4      | 50000 |   25894 ns/op |  79.09 MB/s |
|BenchmarkGobRPC64K       |  5000 |  245786 ns/op | 533.28 MB/s |
|BenchmarkGobRPC64K-2     | 10000 |  188841 ns/op | 694.09 MB/s |
|BenchmarkGobRPC64K-4     | 10000 |  168544 ns/op | 777.67 MB/s |
|BenchmarkProtoRPC1K      | 30000 |   46507 ns/op |  44.04 MB/s |
|BenchmarkProtoRPC1K-2    | 50000 |   38817 ns/op |  52.76 MB/s |
|BenchmarkProtoRPC1K-4    |100000 |   20293 ns/op | 100.92 MB/s |
|BenchmarkProtoRPC64K     |  5000 |  248173 ns/op | 528.15 MB/s |
|BenchmarkProtoRPC64K-2   | 10000 |  203316 ns/op | 644.67 MB/s |
|BenchmarkProtoRPC64K-4   | 10000 |  176833 ns/op | 741.22 MB/s |
|BenchmarkProtoHTTP1K     | 20000 |   67499 ns/op |  30.34 MB/s |
|BenchmarkProtoHTTP1K-2   | 20000 |   66261 ns/op |  30.91 MB/s |
|BenchmarkProtoHTTP1K-4   | 30000 |   54112 ns/op |  37.85 MB/s |
|BenchmarkProtoHTTP64K    |  2000 |  724337 ns/op | 180.95 MB/s |
|BenchmarkProtoHTTP64K-2  |  3000 |  586288 ns/op | 223.56 MB/s |
|BenchmarkProtoHTTP64K-4  |  3000 |  574101 ns/op | 228.31 MB/s |
