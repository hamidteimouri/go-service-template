package grpcd

import (
	"github.com/hamidteimouri/htutils/colog"
	"google.golang.org/grpc"
	"net"
)

func Start() {
	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		panic(colog.MakeRed(err.Error()))
	}

	grpcServer := grpc.NewServer()
	colog.DoGreen("grpc server served at ::50051")
	grpcServer.Serve(listener)

}
