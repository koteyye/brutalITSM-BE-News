package service

import (
	grpcHandler "github.com/koteyye/brutalITSM-BE-News/internal/grpc"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
)

type PermissionService struct {
	gHandler *grpcHandler.GrpcHandler
}

func NewPermissionService(gHandler *grpcHandler.GrpcHandler) *PermissionService {
	return &PermissionService{gHandler: gHandler}
}

func (p PermissionService) GetMe(token string) (models.UserSingle, error) {
	user, err := p.gHandler.GetUserByToken(token)
	if err != nil {
		return user, err
	}

	return user, nil
}
