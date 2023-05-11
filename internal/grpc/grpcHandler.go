package grpcHandler

import (
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	pb "github.com/koteyye/brutalITSM-BE-Users/proto"
)

type GrpcUsers interface {
	GetUserByToken(token string) (models.UserSingle, error)
	GetUserList(userId []string) ([]models.UserList, error)
}

type GrpcHandler struct {
	GrpcUsers
}

func NewGrpcHandler(grpcClient pb.UserServiceClient) *GrpcHandler {
	return &GrpcHandler{
		GrpcUsers: NewGrpcUserHandler(grpcClient),
	}
}
