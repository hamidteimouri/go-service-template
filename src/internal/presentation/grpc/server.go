package grpc

import (
	"github.com/hamidteimouri/htutils/htenvier"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"goservicetemplate/cmd/di"
	"goservicetemplate/internal/presentation/grpc/pbs"
	"net"
)

func StartGRPC() {
	address := htenvier.Env("GRPC_SERVER_ADDRESS")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("failed to make listener for gRPC server")
	}

	grpcServer := grpc.NewServer()

	/* register GRPC servers */
	pbs.RegisterUserServiceServer(grpcServer, di.GrpcUserServer())

	logrus.WithFields(logrus.Fields{
		"grpc_started_at": address,
	}).Debug("gRPC server started")

	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Panic("failed to serve gRPC server")
		}
	}()

}
