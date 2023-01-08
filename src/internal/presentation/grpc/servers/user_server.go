package servers

import (
	"context"
	"github.com/hamidteimouri/gommon/htcolog"
	"goservicetemplate/internal/domain/controllers"
	"goservicetemplate/internal/presentation/grpc/pbs"
)

type UserServer struct {
	ctrl *controllers.UserController
}

func NewUserServer(ctrl *controllers.UserController) *UserServer {
	return &UserServer{ctrl: ctrl}
}

func (u UserServer) GetMe(ctx context.Context, request *pbs.Me) (*pbs.MeReply, error) {
	token := request.GetToken()
	htcolog.DoPurple("token is grpc : " + token)

	return &pbs.MeReply{User: nil}, nil
}

func (u UserServer) UserChangePassword(ctx context.Context, request *pbs.UserChangePasswordRequest) (*pbs.UserChangePasswordReply, error) {
	panic("implement me")
}
