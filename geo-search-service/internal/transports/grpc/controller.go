package grpc

import "google.golang.org/grpc"

type BaseController interface {
	Register(grpc.ServiceRegistrar)
}
