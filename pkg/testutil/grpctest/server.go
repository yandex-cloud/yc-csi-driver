/*
Copyright 2024 YANDEX LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package grpctest

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
)

func NewServer(opts ...grpc.ServerOption) *Server {
	return &Server{
		Listener: newLocalListener(),
		Server:   grpc.NewServer(opts...),
	}
}

type Server struct {
	Listener net.Listener
	Server   *grpc.Server

	serveCtx context.Context
}

// Serve registers passed registerers and starts serving.
// Example:
// instance := &computemocks.InstanceServer{}
// instance.On("Get", mock.Anything, mock.Anything).Return(
//
//	func(context.Context, *compute.GetInstanceRequest) (*compute.Instance, error) {
//	    return &compute.Instance{}, nil
//	})
//
// server := NewServer().Serve(instance)
// defer server.Close()
func (s *Server) Serve(regs ...Registrant) *Server {
	for _, reg := range regs {
		reg.GRPCTestRegister(s.Server)
	}

	s.goServe()

	return s
}

// Registrant should be implemented by generated gRPC mocks.
type Registrant interface {
	// GRPCTestRegister registers gRPC server implementation on passed server.
	GRPCTestRegister(*grpc.Server)
}

func (s *Server) Addr() string {
	return s.Listener.Addr().String()
}

func (s *Server) Close() {
	if !s.isServing() {
		return
	}

	go s.Server.GracefulStop()

	select {
	case <-time.After(10 * time.Second):
		panic("not stopped after timeout")
	case <-s.serveCtx.Done():
	}
}

func (s *Server) goServe() {
	if s.isServing() {
		panic("already serving")
	}

	ctx, cancel := context.WithCancel(context.Background())
	s.serveCtx = ctx

	go func() {
		defer cancel()
		_ = s.Server.Serve(s.Listener)
	}()
}

func (s *Server) isServing() bool {
	return s.serveCtx != nil
}

func newLocalListener() net.Listener {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		l, err = net.Listen("tcp6", "[::1]:0")
		if err != nil {
			panic(fmt.Errorf("grpctest: failed to listen on any port: %w", err))
		}
	}

	return l
}
