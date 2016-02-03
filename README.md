rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go)
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/golang/protobuf) codec
- [net/http](http://godoc.org/net/http) using protobufs on the wire
- [golang.org/x/net/http2](https://godoc.org/golang.org/x/net/http2) using protobufs on the wire

results (2016-02-03)
====================
Run with `go test -benchmem -benchtime 5s -count 10 -bench . > results && benchstat results`

## go version go1.5.3 darwin/amd64

```
name              time/op
GRPC_1K-4           94.7µs ± 4%
GRPC_64K-4          1.90ms ± 6%
GobRPC_1K-4         68.5µs ±14%
GobRPC_64K-4        1.43ms ±12%
ProtoRPC_1K-4       67.6µs ± 5%
ProtoRPC_64K-4      1.57ms ±11%
ProtoHTTP1_1K-4      376µs ± 9%
ProtoHTTP1_64K-4    2.06ms ±22%
ProtoHTTP2_1K-4      167µs ±10%
ProtoHTTP2_64K-4    1.86ms ± 8%

name              speed
GRPC_1K-4         21.6MB/s ± 4%
GRPC_64K-4        69.1MB/s ± 6%
GobRPC_1K-4       30.1MB/s ±13%
GobRPC_64K-4      92.2MB/s ±11%
ProtoRPC_1K-4     30.3MB/s ± 5%
ProtoRPC_64K-4    82.0MB/s ±17%
ProtoHTTP1_1K-4   5.46MB/s ± 9%
ProtoHTTP1_64K-4  64.5MB/s ±19%
ProtoHTTP2_1K-4   12.3MB/s ±10%
ProtoHTTP2_64K-4  70.6MB/s ± 8%

name              alloc/op
GRPC_1K-4           14.8kB ± 0%
GRPC_64K-4           713kB ± 0%
GobRPC_1K-4         4.74kB ± 0%
GobRPC_64K-4         264kB ± 0%
ProtoRPC_1K-4       2.63kB ± 0%
ProtoRPC_64K-4       280kB ± 0%
ProtoHTTP1_1K-4     52.2kB ±16%
ProtoHTTP1_64K-4     832kB ± 1%
ProtoHTTP2_1K-4     89.3kB ± 0%
ProtoHTTP2_64K-4     947kB ± 0%

name              allocs/op
GRPC_1K-4             97.0 ± 0%
GRPC_64K-4             332 ± 0%
GobRPC_1K-4           33.0 ± 0%
GobRPC_64K-4          97.0 ± 0%
ProtoRPC_1K-4         27.0 ± 0%
ProtoRPC_64K-4        93.0 ± 0%
ProtoHTTP1_1K-4        392 ±17%
ProtoHTTP1_64K-4       421 ±12%
ProtoHTTP2_1K-4        118 ± 0%
ProtoHTTP2_64K-4       314 ± 0%
```

## go version go1.6rc1 darwin/amd64

These are faster than go1.5.3 thanks to cloudflare's amd64 assembly implementation of
AES-GCM: https://github.com/golang/go/commit/efeeee3.

```
name              time/op
GRPC_1K-4           66.3µs ±12%
GRPC_64K-4           874µs ±24%
GobRPC_1K-4         38.5µs ±20%
GobRPC_64K-4         294µs ± 2%
ProtoRPC_1K-4       37.1µs ±16%
ProtoRPC_64K-4       356µs ± 3%
ProtoHTTP1_1K-4      128µs ± 3%
ProtoHTTP1_64K-4     674µs ±19%
ProtoHTTP2_1K-4      125µs ± 5%
ProtoHTTP2_64K-4     640µs ±22%

name              speed
GRPC_1K-4         30.9MB/s ±11%
GRPC_64K-4         152MB/s ±20%
GobRPC_1K-4       53.6MB/s ±17%
GobRPC_64K-4       446MB/s ± 2%
ProtoRPC_1K-4     55.7MB/s ±14%
ProtoRPC_64K-4     368MB/s ± 3%
ProtoHTTP1_1K-4   16.0MB/s ± 3%
ProtoHTTP1_64K-4   196MB/s ±17%
ProtoHTTP2_1K-4   16.4MB/s ± 5%
ProtoHTTP2_64K-4   207MB/s ±19%

name              alloc/op
GRPC_1K-4           14.3kB ± 0%
GRPC_64K-4           709kB ± 0%
GobRPC_1K-4         2.41kB ± 0%
GobRPC_64K-4         131kB ± 0%
ProtoRPC_1K-4       2.38kB ± 0%
ProtoRPC_64K-4       279kB ± 0%
ProtoHTTP1_1K-4     88.8kB ± 0%
ProtoHTTP1_64K-4     973kB ± 0%
ProtoHTTP2_1K-4     88.6kB ± 0%
ProtoHTTP2_64K-4     971kB ± 0%

name              allocs/op
GRPC_1K-4             74.0 ± 0%
GRPC_64K-4             140 ± 1%
GobRPC_1K-4           13.0 ± 0%
GobRPC_64K-4          13.0 ± 0%
ProtoRPC_1K-4         11.0 ± 0%
ProtoRPC_64K-4        13.0 ± 0%
ProtoHTTP1_1K-4       85.0 ± 0%
ProtoHTTP1_64K-4       142 ± 0%
ProtoHTTP2_1K-4       84.0 ± 0%
ProtoHTTP2_64K-4       140 ± 0%
```
