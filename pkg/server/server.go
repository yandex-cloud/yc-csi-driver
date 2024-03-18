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

package server

import (
	"net"

	"google.golang.org/grpc"
	"k8s.io/klog"
)

type Server struct {
	Srv    *grpc.Server
	Scheme string
	Addr   string
}

func (s *Server) Run() error {
	listener, err := net.Listen(s.Scheme, s.Addr)
	if err != nil {
		return err
	}

	klog.Infof("Serving on %s://%s", s.Scheme, s.Addr)

	return s.Srv.Serve(listener)
}
