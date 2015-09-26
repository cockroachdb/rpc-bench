rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go)
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/golang/protobuf) codec
- [net/http](http://godoc.org/net/http) using protobufs on the wire
- [golang.org/x/net/http2](https://godoc.org/golang.org/x/net/http2) using protobufs on the wire

results (2015-09-26)
====================
Run with `go test -benchmem -benchtime 5s -bench .`

## go version go1.5.1 darwin/amd64

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

## go version devel +59129c6 Fri Sep 25 22:25:52 2015 +0000 darwin/amd64

These are faster than go1.5.1 thanks to cloudflare's amd64 assembly implementation of
AES-GCM: https://github.com/golang/go/commit/efeeee3.

| Bench          | Iters  | Speed         | Throughput  | Allocations | Allocations    |
|----------------|-------:|--------------:|------------:|------------:|---------------:|
| GRPC_1K        | 100000 |   70215 ns/op |  29.17 MB/s |  16658 B/op |  114 allocs/op |
| GRPC_64K       |  10000 |  768996 ns/op | 170.45 MB/s | 714502 B/op |  357 allocs/op |
| GobRPC_1K      | 200000 |   39776 ns/op |  51.49 MB/s |   4744 B/op |   33 allocs/op |
| GobRPC_64K     |  20000 |  334447 ns/op | 391.91 MB/s | 263858 B/op |   97 allocs/op |
| ProtoRPC_1K    | 200000 |   32926 ns/op |  62.20 MB/s |   2632 B/op |   27 allocs/op |
| ProtoRPC_64K   |  20000 |  341672 ns/op | 383.62 MB/s | 280175 B/op |   93 allocs/op |
| ProtoHTTP1_1K  |  20000 |  372478 ns/op |   5.50 MB/s |  59098 B/op |  448 allocs/op |
| ProtoHTTP1_64K |  10000 | 1414341 ns/op |  92.67 MB/s | 931481 B/op | 1091 allocs/op |

golang.org/x/net/http2 currently doesn't implement sending requests with a body, so those benchmarks fail.
