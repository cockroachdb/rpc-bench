// Copyright 2015 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Tamir Duberstein (tamird@gmail.com)

package rpcbench

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/cockroachdb/rpc-bench/protos"
)

var echoNet = flag.String("echo-net", "tcp",
	"network to bind for the echo server used in benchmarks")
var echoAddr = flag.String("echo-addr", "127.0.0.1:0",
	"host:port to bind for the echo server used in benchmarks")
var runEchoServer = flag.Bool("start-echo-server", true,
	"start the echo server; false to connect to an already running server")
var onlyEchoServer = flag.Bool("only-echo-server", false,
	"only run the echo server; looping forever")

// To run these benchmarks between machines, on machine 1 start the
// echo server:
//
//   go test -run= -bench=BenchmarkEchoGobRPC -echoAddr :9999 -only-echo-server
//
// On machine 2:
//
//   go test -run= -bench=BenchmarkEchoGobRPC -echoAddr <machine-1-ip>:9999 -start-echo-server=false

func randString(n int) string {
	var randLetters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789#$")
	return string(bytes.Repeat(randLetters, n/len(randLetters)))
}

func newListener(tb testing.TB) net.Listener {
	listener, err := net.Listen(*echoNet, *echoAddr)
	if err != nil {
		tb.Fatal(err)
	}
	return listener
}

func benchmarkEcho(b *testing.B, size int, listenAndServe func(net.Listener) error, parallelRequest func(*testing.PB, net.Addr, string)) {
	listener := newListener(b)
	defer func() {
		if err := listener.Close(); err != nil {
			b.Fatal(err)
		}
	}()

	go func() {
		if err := listenAndServe(listener); err != nil && !strings.HasSuffix(err.Error(), "use of closed network connection") {
			b.Fatal(err)
		}
	}()

	echoMsg := randString(size)

	b.SetBytes(2 * int64(len(echoMsg)))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		parallelRequest(pb, listener.Addr(), echoMsg)
	})
}

// grpc

type echoServer struct{}

func (e *echoServer) Echo(ctx context.Context, req *protos.EchoRequest) (*protos.EchoResponse, error) {
	return &protos.EchoResponse{Msg: req.Msg}, nil
}

type benchGRPC struct{}

func listenAndServeGRPC(listener net.Listener) error {
	grpcServer := grpc.NewServer()
	protos.RegisterEchoServer(grpcServer, &echoServer{})
	return grpcServer.Serve(listener)
}

func benchmarkEchoGRPC(b *testing.B, size int) {
	benchmarkEcho(b, size, listenAndServeGRPC, func(pb *testing.PB, addr net.Addr, echoMsg string) {
		conn, err := grpc.Dial(addr.String())
		if err != nil {
			b.Fatal(err)
		}
		defer func() {
			if err = conn.Close(); err != nil {
				b.Fatal(err)
			}
		}()
		client := protos.NewEchoClient(conn)

		for pb.Next() {
			if _, err := client.Echo(context.Background(), &protos.EchoRequest{Msg: echoMsg}); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkGRPC1K(b *testing.B) {
	benchmarkEchoGRPC(b, 1<<10)
}

func BenchmarkGRPC64K(b *testing.B) {
	benchmarkEchoGRPC(b, 64<<10)
}

// gob-rpc

type Echo struct{}

func (t *Echo) Echo(args *protos.EchoRequest, reply *protos.EchoResponse) error {
	reply.Msg = args.Msg
	return nil
}

func listenAndServeGobRPC(listener net.Listener) error {
	rpcServer := rpc.NewServer()
	if err := rpcServer.Register(&Echo{}); err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go rpcServer.ServeConn(conn)
	}
}

func benchmarkEchoGobRPC(b *testing.B, size int) {
	benchmarkEcho(b, size, listenAndServeGobRPC, func(pb *testing.PB, addr net.Addr, echoMsg string) {
		client, err := rpc.Dial(addr.Network(), addr.String())
		if err != nil {
			b.Fatal(err)
		}
		defer func() {
			if err = client.Close(); err != nil {
				b.Fatal(err)
			}
		}()

		for pb.Next() {
			args := &protos.EchoRequest{Msg: echoMsg}
			reply := &protos.EchoResponse{}
			if err := client.Call("Echo.Echo", args, reply); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkGobRPC1K(b *testing.B) {
	benchmarkEchoGobRPC(b, 1<<10)
}

func BenchmarkGobRPC64K(b *testing.B) {
	benchmarkEchoGobRPC(b, 64<<10)
}

// proto-rpc

func listenAndServeProtoRPC(listener net.Listener) error {
	rpcServer := rpc.NewServer()
	if err := rpcServer.Register(&Echo{}); err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go rpcServer.ServeCodec(NewServerCodec(conn))
	}
}

func benchmarkEchoProtoRPC(b *testing.B, size int) {
	benchmarkEcho(b, size, listenAndServeProtoRPC, func(pb *testing.PB, addr net.Addr, echoMsg string) {
		conn, err := net.Dial(addr.Network(), addr.String())
		if err != nil {
			b.Fatal(err)
		}
		client := rpc.NewClientWithCodec(NewClientCodec(conn))
		defer func() {
			if err = client.Close(); err != nil {
				b.Fatal(err)
			}
		}()

		for pb.Next() {
			args := &protos.EchoRequest{Msg: echoMsg}
			reply := &protos.EchoResponse{}
			if err := client.Call("Echo.Echo", args, reply); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkProtoRPC1K(b *testing.B) {
	benchmarkEchoProtoRPC(b, 1<<10)
}

func BenchmarkProtoRPC64K(b *testing.B) {
	benchmarkEchoProtoRPC(b, 64<<10)
}

// proto-http

const (
	contentType = "Content-Type"
	xProtobuf   = "application/x-protobuf"
)

func listenAndServeProtoHTTP(listener net.Listener) error {
	return http.Serve(listener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = r.Body.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		args := &protos.EchoRequest{}
		if err = proto.Unmarshal(reqBody, args); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		reply := &protos.EchoResponse{Msg: args.Msg}
		respBody, err := proto.Marshal(reply)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set(contentType, xProtobuf)
		w.Write(respBody)
	}))
}

func benchmarkEchoProtoHTTP(b *testing.B, size int) {
	benchmarkEcho(b, size, listenAndServeProtoHTTP, func(pb *testing.PB, addr net.Addr, echoMsg string) {
		url := fmt.Sprintf("http://%s", addr)

		for pb.Next() {
			args := &protos.EchoRequest{Msg: echoMsg}
			reqBody, err := proto.Marshal(args)
			if err != nil {
				b.Fatal(err)
			}
			resp, err := http.Post(url, xProtobuf, bytes.NewReader(reqBody))
			if err != nil {
				b.Fatal(err)
			}
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				b.Fatal(err)
			}
			if err = resp.Body.Close(); err != nil {
				b.Fatal(err)
			}
			reply := &protos.EchoResponse{}
			if err := proto.Unmarshal(respBody, reply); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkProtoHTTP1K(b *testing.B) {
	benchmarkEchoProtoHTTP(b, 1<<10)
}

func BenchmarkProtoHTTP64K(b *testing.B) {
	benchmarkEchoProtoHTTP(b, 64<<10)
}
