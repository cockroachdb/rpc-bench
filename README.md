rpc-bench
=========

Currently covered:
- [grpc-go](https://github.com/grpc/grpc-go)
- [net/rpc](http://godoc.org/net/rpc)
- [net/rpc](http://godoc.org/net/rpc) + [protobuf](https://github.com/golang/protobuf) codec
- [net/http](http://godoc.org/net/http) using protobufs on the wire

results (2015-09-11)
====================
Run with `go test -benchmem -benchtime 5s -bench .`

| Bench        | Iters  | Speed        | Troughput   | Allocations | Allocations   |
|--------------|-------:|-------------:|------------:|------------:|--------------:|
| GRPC1K       | 200000 |  58742 ns/op |  34.86 MB/s |  16226 B/op |  89 allocs/op |
| GRPC64K      |  10000 | 582779 ns/op | 224.91 MB/s | 711277 B/op | 153 allocs/op |
| GobRPC1K     | 300000 |  26422 ns/op |  77.51 MB/s |   4488 B/op |  17 allocs/op |
| GobRPC64K    |  50000 | 188670 ns/op | 694.71 MB/s | 262600 B/op |  17 allocs/op |
| ProtoRPC1K   | 300000 |  22599 ns/op |  90.62 MB/s |   2376 B/op |  11 allocs/op |
| ProtoRPC64K  |  50000 | 187764 ns/op | 698.06 MB/s | 278918 B/op |  13 allocs/op |
| ProtoHTTP1K  | 200000 |  54141 ns/op |  37.83 MB/s |  14638 B/op |  81 allocs/op |
| ProtoHTTP64K |  10000 | 725727 ns/op | 180.61 MB/s | 813301 B/op | 110 allocs/op |
