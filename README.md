rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go)
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/golang/protobuf) codec
- [net/http](http://godoc.org/net/http) using protobufs on the wire
- [golang.org/x/net/http2](https://godoc.org/golang.org/x/net/http2) using protobufs on the wire

results (2015-10-29)
====================
Run with `go test -benchmem -benchtime 5s -bench .`

## go version go1.5.1 darwin/amd64

| Bench            | Iters  | Speed         | Throughput | Allocations | Allocations   |
|------------------|-------:|--------------:|-----------:|------------:|--------------:|
| GRPC_1K-4        | 100000 |  103406 ns/op | 19.81 MB/s |  16134 B/op | 106 allocs/op |
| GRPC_64K-4       |   3000 | 1960238 ns/op | 66.87 MB/s | 714063 B/op | 341 allocs/op |
| GobRPC_1K-4      | 100000 |   63812 ns/op | 32.09 MB/s |   4744 B/op |  33 allocs/op |
| GobRPC_64K-4     |   5000 | 1414237 ns/op | 92.68 MB/s | 263935 B/op |  97 allocs/op |
| ProtoRPC_1K-4    | 200000 |   56618 ns/op | 36.17 MB/s |   2632 B/op |  27 allocs/op |
| ProtoRPC_64K-4   |   5000 | 1413800 ns/op | 92.71 MB/s | 280219 B/op |  93 allocs/op |
| ProtoHTTP1_1K-4  |  20000 |  388989 ns/op |  5.26 MB/s |  57736 B/op | 436 allocs/op |
| ProtoHTTP1_64K-4 |   3000 | 2001305 ns/op | 65.49 MB/s | 823248 B/op | 345 allocs/op |
| ProtoHTTP2_1K-4  |  50000 |  165243 ns/op | 12.39 MB/s |  89641 B/op | 107 allocs/op |
| ProtoHTTP2_64K-4 |   5000 | 1804282 ns/op | 72.64 MB/s | 944128 B/op | 330 allocs/op |

## go version devel +51586aa Thu Oct 29 19:27:47 2015 +0000 darwin/amd64

These are faster than go1.5.1 thanks to cloudflare's amd64 assembly implementation of
AES-GCM: https://github.com/golang/go/commit/efeeee3.

| Bench            | Iters  | Speed         | Throughput  | Allocations | Allocations    |
|------------------|-------:|--------------:|------------:|------------:|---------------:|
| GRPC_1K-4        | 100000 |   74513 ns/op |  27.48 MB/s |  16055 B/op |  106 allocs/op |
| GRPC_64K-4       |  10000 |  858855 ns/op | 152.61 MB/s | 714106 B/op |  346 allocs/op |
| GobRPC_1K-4      | 200000 |   40553 ns/op |  50.50 MB/s |   4745 B/op |   33 allocs/op |
| GobRPC_64K-4     |  20000 |  421501 ns/op | 310.96 MB/s | 263899 B/op |   97 allocs/op |
| ProtoRPC_1K-4    | 200000 |   33851 ns/op |  60.50 MB/s |   2632 B/op |   27 allocs/op |
| ProtoRPC_64K-4   |  20000 |  479764 ns/op | 273.20 MB/s | 280209 B/op |   93 allocs/op |
| ProtoHTTP1_1K-4  |  30000 |  332801 ns/op |   6.15 MB/s |  52926 B/op |  399 allocs/op |
| ProtoHTTP1_64K-4 |   5000 | 1419998 ns/op |  92.30 MB/s | 929968 B/op | 1080 allocs/op |
| ProtoHTTP2_1K-4  |  50000 |  128962 ns/op |  15.88 MB/s |  89091 B/op |  107 allocs/op |
| ProtoHTTP2_64K-4 |  10000 |  662109 ns/op | 197.96 MB/s | 971456 B/op |  316 allocs/op |
