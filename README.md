rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go)
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/golang/protobuf) codec
- [net/http](http://godoc.org/net/http) using protobufs on the wire
- [golang.org/x/net/http2](https://godoc.org/golang.org/x/net/http2) using protobufs on the wire

results (2015-09-25)
====================
Run with `go test -benchmem -benchtime 5s -bench .`

| Bench          | Iters  | Speed         | Throughput | Allocations | Allocations   |
|----------------|-------:|--------------:|-----------:|------------:|--------------:|
| GRPC_1K        | 100000 |   92573 ns/op | 22.12 MB/s |  16661 B/op | 114 allocs/op |
| GRPC_64K       |   3000 | 2044461 ns/op | 64.11 MB/s | 714326 B/op | 349 allocs/op |
| GobRPC_1K      | 100000 |   72621 ns/op | 28.20 MB/s |   4744 B/op |  33 allocs/op |
| GobRPC_64K     |   5000 | 1873094 ns/op | 69.98 MB/s | 263934 B/op |  97 allocs/op |
| ProtoRPC_1K    | 100000 |   66955 ns/op | 30.59 MB/s |   2632 B/op |  27 allocs/op |
| ProtoRPC_64K   |   5000 | 1564185 ns/op | 83.80 MB/s | 280220 B/op |  93 allocs/op |
| ProtoHTTP1_1K  |  20000 |  383342 ns/op |  5.34 MB/s |  55044 B/op | 416 allocs/op |
| ProtoHTTP1_64K |   3000 | 1932455 ns/op | 67.83 MB/s | 827150 B/op | 377 allocs/op |

golang.org/x/net/http2 currently doesn't implement sending requests with a body, so those benchmarks fail.
