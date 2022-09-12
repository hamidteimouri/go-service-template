package grpc

import (
	"fmt"
	"github.com/hamidteimouri/htutils/colog"
	"github.com/hamidteimouri/htutils/envier"
	"google.golang.org/grpc"
	"net"
)

func StartGRPC() {
	addr := envier.Env("GRPC_SERVER_ADDRESS")
	port := envier.Env("GRPC_SERVER_PORT")
	address := fmt.Sprintf("%s:%s", addr, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(colog.MakeRed(err.Error()))
	}

	grpcServer := grpc.NewServer()
	log := fmt.Sprintf("⇨ grpc server started on %s", colog.MakeGreen(address))
	fmt.Println(log)

	go func() {

		err = grpcServer.Serve(listener)
		if err != nil {
			e := fmt.Sprintf("faild to start GRPC server : %s", colog.MakeRed(err.Error()))
			panic(e)
		}
	}()

}
