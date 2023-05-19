package service

import (
	grpcHandler "github.com/koteyye/brutalITSM-BE-News/internal/grpc"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	"github.com/koteyye/brutalITSM-BE-News/internal/postgres"
	"github.com/minio/minio-go/v7"
)

type News interface {
	CreateNews(news models.News) (string, error)
	UpdateNews(newsId string, news models.News) (string, error)
	DeleteNews(newsId string) (bool, error)
	GetNewsList() ([]models.ResponseNewsList, error)
	GetNewsById(newsId string) (models.NewsList, error)
}

type Permissions interface {
	GetMe(token string) (models.UserSingle, error)
}

type Service struct {
	News
	Permissions
}

func NewService(repos *postgres.Repository, s3 *minio.Client, gHandler *grpcHandler.GrpcHandler) *Service {
	return &Service{
		News:        NewNewsService(repos.News, s3, gHandler),
		Permissions: NewPermissionService(gHandler),
	}
}
