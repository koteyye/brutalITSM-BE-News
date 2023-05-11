package grpcHandler

import (
	"context"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	pb "github.com/koteyye/brutalITSM-BE-Users/proto"
	"github.com/sirupsen/logrus"
	"time"
)

type GrpcUserHandler struct {
	grpcUserHandler pb.UserServiceClient
}

func NewGrpcUserHandler(grpcUserHandler pb.UserServiceClient) *GrpcUserHandler {
	return &GrpcUserHandler{grpcUserHandler: grpcUserHandler}
}

func (g GrpcUserHandler) GetUserByToken(token string) (models.UserSingle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := g.grpcUserHandler.GetByToken(ctx, &pb.RequestToken{Token: token})
	if err != nil {
		logrus.Fatalf("Could not auth: %v", err)
	}
	return models.UserSingle{
		Id:          res.Id,
		Login:       res.Login,
		LastName:    res.LastName,
		FirstName:   res.FirstName,
		SurName:     res.SurName,
		Job:         res.Job,
		Org:         res.Org,
		Roles:       res.Roles,
		Permissions: res.Permissions,
	}, nil
}

func (g GrpcUserHandler) GetUserList(userId []string) ([]models.UserList, error) {
	var userList []models.UserList
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := g.grpcUserHandler.GetByUserList(ctx, &pb.RequestUsers{Id: userId})
	if err != nil {
		return nil, err
	}

	for _, users := range res.UserList {
		userList = append(userList, models.UserList{
			Id:         users.Id,
			LastName:   users.LastName,
			FirstName:  users.FirstName,
			SurName:    users.SurName,
			MimeType:   users.MimeType,
			BucketName: users.BucketName,
			FileName:   users.FileName,
		})
	}
	return userList, nil
}
