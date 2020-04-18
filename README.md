rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go) using (*Server).Serve
- [grpc-go](https://github.com/grpc/grpc-go) using (*Server).ServeHTTP
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/gogo/protobuf) codec
- [net/http](http://godoc.org/net/http) + [protobuf](https://github.com/gogo/protobuf) on the wire
- [golang.org/x/net/http2](https://godoc.org/golang.org/x/net/http2) + [protobuf](https://github.com/gogo/protobuf) on the wire

results (2020-04-18)
====================
Run with `go test -benchmem -benchtime 5s -count 5 -bench . -timeout 1h | tee results && benchstat results`

## go version go1.14 linux/amd64

```
name                         time/op
GRPCServe_1K-36                62.6µs ± 7%
GRPCServe_64K-36                263µs ± 2%
GRPCServe_Stream_1K-36         31.1µs ± 3%
GRPCServe_Stream_64k-36         264µs ± 2%
GRPCServeHTTP_1K-36             262µs ± 2%
GRPCServeHTTP_64K-36            642µs ± 2%
GRPCServeHTTP_Stream_1K-36     42.3µs ± 2%
GRPCServeHTTP_Stream_64k-36     538µs ± 1%
GobRPC_1K-36                   29.7µs ± 3%
GobRPC_64K-36                   204µs ± 1%
ProtoRPC_1K-36                 27.3µs ± 2%
ProtoRPC_64K-36                 221µs ± 1%
ProtoHTTP1_1K-36               65.2µs ±15%
ProtoHTTP1_64K-36               172µs ± 9%
ProtoHTTP2_1K-36                115µs ±24%
ProtoHTTP2_64K-36               750µs ± 3%

name                         speed
GRPCServe_1K-36              32.7MB/s ± 7%
GRPCServe_64K-36              499MB/s ± 2%
GRPCServe_Stream_1K-36       65.8MB/s ± 3%
GRPCServe_Stream_64k-36       497MB/s ± 2%
GRPCServeHTTP_1K-36          7.81MB/s ± 2%
GRPCServeHTTP_64K-36          204MB/s ± 2%
GRPCServeHTTP_Stream_1K-36   48.5MB/s ± 2%
GRPCServeHTTP_Stream_64k-36   244MB/s ± 1%
GobRPC_1K-36                 69.0MB/s ± 3%
GobRPC_64K-36                 643MB/s ± 1%
ProtoRPC_1K-36               75.0MB/s ± 2%
ProtoRPC_64K-36               594MB/s ± 1%
ProtoHTTP1_1K-36             31.6MB/s ±14%
ProtoHTTP1_64K-36             763MB/s ± 9%
ProtoHTTP2_1K-36             18.0MB/s ±20%
ProtoHTTP2_64K-36             175MB/s ± 3%

name                         alloc/op
GRPCServe_1K-36                17.7kB ± 0%
GRPCServe_64K-36                489kB ± 0%
GRPCServe_Stream_1K-36         10.1kB ± 0%
GRPCServe_Stream_64k-36         484kB ± 0%
GRPCServeHTTP_1K-36            37.4kB ± 0%
GRPCServeHTTP_64K-36            563kB ± 0%
GRPCServeHTTP_Stream_1K-36     10.2kB ± 0%
GRPCServeHTTP_Stream_64k-36     546kB ± 0%
GobRPC_1K-36                   2.46kB ± 0%
GobRPC_64K-36                   132kB ± 0%
ProtoRPC_1K-36                 2.43kB ± 0%
ProtoRPC_64K-36                 279kB ± 0%
ProtoHTTP1_1K-36               56.8kB ± 1%
ProtoHTTP1_64K-36               904kB ± 0%
ProtoHTTP2_1K-36               63.6kB ± 2%
ProtoHTTP2_64K-36              1.02MB ±10%

name                         allocs/op
GRPCServe_1K-36                   154 ± 0%
GRPCServe_64K-36                  167 ± 0%
GRPCServe_Stream_1K-36           29.0 ± 0%
GRPCServe_Stream_64k-36          52.0 ± 0%
GRPCServeHTTP_1K-36               212 ± 0%
GRPCServeHTTP_64K-36              326 ± 1%
GRPCServeHTTP_Stream_1K-36       36.0 ± 0%
GRPCServeHTTP_Stream_64k-36       168 ± 1%
GobRPC_1K-36                     14.0 ± 0%
GobRPC_64K-36                    20.0 ± 0%
ProtoRPC_1K-36                   12.0 ± 0%
ProtoRPC_64K-36                  18.0 ± 0%
ProtoHTTP1_1K-36                  634 ± 1%
ProtoHTTP1_64K-36                 699 ± 1%
ProtoHTTP2_1K-36                  114 ± 0%
ProtoHTTP2_64K-36                 213 ± 0%
```
