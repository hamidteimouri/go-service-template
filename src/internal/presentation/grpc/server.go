package grpc

import (
	"fmt"
	"github.com/hamidteimouri/htutils/htcolog"
	"github.com/hamidteimouri/htutils/htenvier"
	"google.golang.org/grpc"
	"goservicetemplate/cmd/di"
	"goservicetemplate/internal/presentation/grpc/pbs"
	"net"
)

func StartGRPC() {
	addr := htenvier.Env("GRPC_SERVER_ADDRESS")
	port := htenvier.Env("GRPC_SERVER_PORT")
	address := fmt.Sprintf("%s:%s", addr, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(htcolog.MakeRed(err.Error()))
	}

	grpcServer := grpc.NewServer()
	log := fmt.Sprintf("⇨ grpc server started on %s", htcolog.MakeGreen(address))
	fmt.Println(log)

	/* register GRPC servers */
	pbs.RegisterUserServiceServer(grpcServer, di.GrpcUserServer())

	go func() {

		err = grpcServer.Serve(listener)
		if err != nil {
			e := fmt.Sprintf("faild to start GRPC server : %s", htcolog.MakeRed(err.Error()))
			panic(e)
		}
	}()

}
