rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go) using (*Server).Serve
- [grpc-go](https://github.com/grpc/grpc-go) using (*Server).ServeHTTP
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/golang/protobuf) codec
- [net/http](http://godoc.org/net/http) using protobufs on the wire
- [golang.org/x/net/http2](https://godoc.org/golang.org/x/net/http2) using protobufs on the wire

results (2016-02-03)
====================
Run with `go test -benchmem -benchtime 5s -count 10 -bench . > results && benchstat results`

## go version go1.6 darwin/amd64

```
name                 time/op
GRPCServe_1K-4         71.3µs ± 7%
GRPCServe_64K-4         916µs ±17%
GRPCServeHTTP_1K-4      196µs ±12%
GRPCServeHTTP_64K-4    1.39ms ± 5%
GobRPC_1K-4            38.9µs ±11%
GobRPC_64K-4            300µs ± 5%
ProtoRPC_1K-4          36.4µs ±13%
ProtoRPC_64K-4          381µs ± 8%
ProtoHTTP1_1K-4         535µs ±17%
ProtoHTTP1_64K-4       1.21ms ± 7%
ProtoHTTP2_1K-4         147µs ±11%
ProtoHTTP2_64K-4        788µs ±12%

name                 speed
GRPCServe_1K-4       28.8MB/s ± 8%
GRPCServe_64K-4       145MB/s ±15%
GRPCServeHTTP_1K-4   10.5MB/s ±11%
GRPCServeHTTP_64K-4  94.2MB/s ± 4%
GobRPC_1K-4          52.9MB/s ±10%
GobRPC_64K-4          438MB/s ± 5%
ProtoRPC_1K-4        56.8MB/s ±14%
ProtoRPC_64K-4        345MB/s ± 8%
ProtoHTTP1_1K-4      3.85MB/s ±15%
ProtoHTTP1_64K-4      109MB/s ± 7%
ProtoHTTP2_1K-4      14.0MB/s ±10%
ProtoHTTP2_64K-4      167MB/s ±13%

name                 alloc/op
GRPCServe_1K-4         15.6kB ± 0%
GRPCServe_64K-4         711kB ± 0%
GRPCServeHTTP_1K-4     88.5kB ± 0%
GRPCServeHTTP_64K-4     801kB ± 0%
GobRPC_1K-4            2.41kB ± 0%
GobRPC_64K-4            131kB ± 0%
ProtoRPC_1K-4          2.38kB ± 0%
ProtoRPC_64K-4          279kB ± 0%
ProtoHTTP1_1K-4         111kB ± 5%
ProtoHTTP1_64K-4        939kB ± 1%
ProtoHTTP2_1K-4        89.9kB ± 0%
ProtoHTTP2_64K-4        979kB ± 0%

name                 allocs/op
GRPCServe_1K-4           96.0 ± 0%
GRPCServe_64K-4           162 ± 0%
GRPCServeHTTP_1K-4        166 ± 0%
GRPCServeHTTP_64K-4       671 ± 0%
GobRPC_1K-4              13.0 ± 0%
GobRPC_64K-4             13.0 ± 0%
ProtoRPC_1K-4            11.0 ± 0%
ProtoRPC_64K-4           13.0 ± 0%
ProtoHTTP1_1K-4           829 ± 5%
ProtoHTTP1_64K-4          966 ± 7%
ProtoHTTP2_1K-4           100 ± 0%
ProtoHTTP2_64K-4          155 ± 0%
```
