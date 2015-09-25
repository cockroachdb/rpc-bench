rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go)
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/golang/protobuf) codec
- [net/http](http://godoc.org/net/http) using protobufs on the wire

results (2015-09-25)
====================
Run with `go test -benchmem -benchtime 5s -bench .`

| Bench        | Iters  | Speed        | Throughput  | Allocations | Allocations   |
|--------------|-------:|-------------:|------------:|------------:|--------------:|
| GRPC1K       | 100000 |  88730 ns/op |  23.08 MB/s |  16229 B/op |  89 allocs/op |
| GRPC64K      |  10000 | 806472 ns/op | 162.52 MB/s | 711208 B/op | 152 allocs/op |
| GobRPC1K     | 200000 |  40529 ns/op |  50.53 MB/s |   4488 B/op |  17 allocs/op |
| GobRPC64K    |  30000 | 265331 ns/op | 493.99 MB/s | 262563 B/op |  17 allocs/op |
| ProtoRPC1K   | 300000 |  30936 ns/op |  66.20 MB/s |   2418 B/op |  11 allocs/op |
| ProtoRPC64K  |  20000 | 272812 ns/op | 480.45 MB/s | 278883 B/op |  13 allocs/op |
| ProtoHTTP1K  | 200000 |  62290 ns/op |  32.88 MB/s |  14528 B/op |  81 allocs/op |
| ProtoHTTP64K |  10000 | 765288 ns/op | 171.27 MB/s | 812307 B/op | 108 allocs/op |
