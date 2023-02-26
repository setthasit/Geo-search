package server

import (
	controller "geo-search/internal/transports/grpc"
	"geo-search/internal/utils"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	estateCtrl *controller.EstateController
}

func WireGrpcServer(
	estateCtrl *controller.EstateController,
) *GrpcServer {
	return &GrpcServer{
		estateCtrl: estateCtrl,
	}
}

func (server *GrpcServer) Run() {
	grpcServer := grpc.NewServer()
	server.estateCtrl.Register(grpcServer)

	listener, err := net.Listen("tcp", ":7001")
	if err != nil {
		utils.LogFatalf("Failed to listen: %v", err)
	}

	utils.LogInfo("Listen to :7001")
	err = grpcServer.Serve(listener)
	if err != nil {
		utils.LogFatalf("Failed to start gRPC server: %v", err)
	}
}
