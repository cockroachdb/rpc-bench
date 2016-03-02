rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go) using (*Server).Serve
- [grpc-go](https://github.com/grpc/grpc-go) using (*Server).ServeHTTP
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/golang/protobuf) codec
- [net/http](http://godoc.org/net/http) using protobufs on the wire
- [golang.org/x/net/http2](https://godoc.org/golang.org/x/net/http2) using protobufs on the wire

results (2016-03-02)
====================
Run with `go test -benchmem -benchtime 5s -count 5 -bench . | tee results && benchstat results`

## go version go1.6 darwin/amd64

```
name                 time/op
GRPCServe_1K-4         72.6µs ±10%
GRPCServe_64K-4         785µs ± 8%
GRPCServeHTTP_1K-4      178µs ± 2%
GRPCServeHTTP_64K-4    1.32ms ± 2%
GobRPC_1K-4            39.4µs ±15%
GobRPC_64K-4            354µs ± 9%
ProtoRPC_1K-4          35.1µs ± 8%
ProtoRPC_64K-4          388µs ±10%
ProtoHTTP1_1K-4         264µs ±12%
ProtoHTTP1_64K-4       1.08ms ±12%
ProtoHTTP2_1K-4         146µs ±16%
ProtoHTTP2_64K-4        702µs ±28%

name                 speed
GRPCServe_1K-4       28.3MB/s ± 9%
GRPCServe_64K-4       167MB/s ± 7%
GRPCServeHTTP_1K-4   11.5MB/s ± 2%
GRPCServeHTTP_64K-4  99.4MB/s ± 2%
GobRPC_1K-4          52.4MB/s ±13%
GobRPC_64K-4          371MB/s ± 8%
ProtoRPC_1K-4        58.6MB/s ± 8%
ProtoRPC_64K-4        339MB/s ±10%
ProtoHTTP1_1K-4      7.79MB/s ±13%
ProtoHTTP1_64K-4      122MB/s ±11%
ProtoHTTP2_1K-4      14.1MB/s ±14%
ProtoHTTP2_64K-4      190MB/s ±23%

name                 alloc/op
GRPCServe_1K-4         15.6kB ± 0%
GRPCServe_64K-4         711kB ± 0%
GRPCServeHTTP_1K-4     88.5kB ± 0%
GRPCServeHTTP_64K-4     801kB ± 0%
GobRPC_1K-4            2.41kB ± 0%
GobRPC_64K-4            131kB ± 0%
ProtoRPC_1K-4          2.38kB ± 0%
ProtoRPC_64K-4          279kB ± 0%
ProtoHTTP1_1K-4        70.5kB ± 6%
ProtoHTTP1_64K-4        935kB ± 1%
ProtoHTTP2_1K-4        89.9kB ± 0%
ProtoHTTP2_64K-4        977kB ± 1%

name                 allocs/op
GRPCServe_1K-4           96.0 ± 0%
GRPCServe_64K-4           162 ± 0%
GRPCServeHTTP_1K-4        166 ± 0%
GRPCServeHTTP_64K-4       671 ± 0%
GobRPC_1K-4              13.0 ± 0%
GobRPC_64K-4             13.0 ± 0%
ProtoRPC_1K-4            11.0 ± 0%
ProtoRPC_64K-4           13.0 ± 0%
ProtoHTTP1_1K-4           515 ± 7%
ProtoHTTP1_64K-4          936 ± 7%
ProtoHTTP2_1K-4           100 ± 0%
ProtoHTTP2_64K-4          156 ± 0%
```
