rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go) using (*Server).Serve
- [grpc-go](https://github.com/grpc/grpc-go) using (*Server).ServeHTTP
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/gogo/protobuf) codec
- [net/http](http://godoc.org/net/http) + [protobuf](https://github.com/gogo/protobuf) on the wire
- [golang.org/x/net/http2](https://godoc.org/golang.org/x/net/http2) + [protobuf](https://github.com/gogo/protobuf) on the wire

results (2017-03-13)
====================
Run with `go test -benchmem -benchtime 5s -count 5 -bench . -timeout 1h | tee results && benchstat results`

## go version go1.8 linux/amd64

```
name                         time/op
GRPCServe_1K-24                40.5µs ± 1%
GRPCServe_64K-24                365µs ± 1%
GRPCServe_Stream_1K-24         22.7µs ± 1%
GRPCServe_Stream_64k-24         298µs ± 1%
GRPCServeHTTP_1K-24             137µs ± 1%
GRPCServeHTTP_64K-24            585µs ± 1%
GRPCServeHTTP_Stream_1K-24     38.7µs ± 1%
GRPCServeHTTP_Stream_64k-24     455µs ± 1%
GobRPC_1K-24                   19.5µs ± 1%
GobRPC_64K-24                   159µs ± 0%
ProtoRPC_1K-24                 19.9µs ± 2%
ProtoRPC_64K-24                 208µs ± 3%
ProtoHTTP1_1K-24                153µs ±18%
ProtoHTTP1_64K-24               329µs ± 6%
ProtoHTTP2_1K-24                113µs ±32%
ProtoHTTP2_64K-24               458µs ± 3%

name                         speed
GRPCServe_1K-24              50.5MB/s ± 1%
GRPCServe_64K-24              359MB/s ± 1%
GRPCServe_Stream_1K-24       90.3MB/s ± 1%
GRPCServe_Stream_64k-24       439MB/s ± 1%
GRPCServeHTTP_1K-24          14.9MB/s ± 1%
GRPCServeHTTP_64K-24          224MB/s ± 1%
GRPCServeHTTP_Stream_1K-24   52.9MB/s ± 1%
GRPCServeHTTP_Stream_64k-24   288MB/s ± 1%
GobRPC_1K-24                  105MB/s ± 1%
GobRPC_64K-24                 822MB/s ± 0%
ProtoRPC_1K-24                103MB/s ± 2%
ProtoRPC_64K-24               631MB/s ± 3%
ProtoHTTP1_1K-24             13.5MB/s ±16%
ProtoHTTP1_64K-24             399MB/s ± 6%
ProtoHTTP2_1K-24             19.2MB/s ±29%
ProtoHTTP2_64K-24             286MB/s ± 3%

name                         alloc/op
GRPCServe_1K-24                16.3kB ± 0%
GRPCServe_64K-24                711kB ± 0%
GRPCServe_Stream_1K-24         11.6kB ± 0%
GRPCServe_Stream_64k-24         707kB ± 0%
GRPCServeHTTP_1K-24            34.2kB ± 0%
GRPCServeHTTP_64K-24            763kB ± 0%
GRPCServeHTTP_Stream_1K-24     12.4kB ± 0%
GRPCServeHTTP_Stream_64k-24     739kB ± 0%
GobRPC_1K-24                   2.41kB ± 0%
GobRPC_64K-24                   131kB ± 0%
ProtoRPC_1K-24                 2.38kB ± 0%
ProtoRPC_64K-24                 279kB ± 0%
ProtoHTTP1_1K-24               50.3kB ± 7%
ProtoHTTP1_64K-24               908kB ± 0%
ProtoHTTP2_1K-24               272kB ±104%
ProtoHTTP2_64K-24              1.09MB ±23%

name                         allocs/op
GRPCServe_1K-24                  97.0 ± 0%
GRPCServe_64K-24                  147 ± 0%
GRPCServe_Stream_1K-24           22.0 ± 0%
GRPCServe_Stream_64k-24          72.0 ± 0%
GRPCServeHTTP_1K-24               169 ± 0%
GRPCServeHTTP_64K-24              275 ± 0%
GRPCServeHTTP_Stream_1K-24       29.0 ± 0%
GRPCServeHTTP_Stream_64k-24       138 ± 0%
GobRPC_1K-24                     13.0 ± 0%
GobRPC_64K-24                    13.0 ± 0%
ProtoRPC_1K-24                   11.0 ± 0%
ProtoRPC_64K-24                  13.0 ± 0%
ProtoHTTP1_1K-24                  329 ± 5%
ProtoHTTP1_64K-24                 414 ± 1%
ProtoHTTP2_1K-24                  110 ± 3%
ProtoHTTP2_64K-24                 178 ± 1%
```
