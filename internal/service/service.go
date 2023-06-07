package service

import (
	"context"
	grpcHandler "github.com/koteyye/brutalITSM-BE-News/internal/grpc"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	"github.com/koteyye/brutalITSM-BE-News/internal/postgres"
	"github.com/minio/minio-go/v7"
	"io"
)

type News interface {
	CreateNews(news models.News, userId string) (string, error)
	UpdateNews(news models.News, userId string) (bool, error)
	DeleteNews(newsId string) (bool, error)
	GetNewsList() ([]models.ResponseNewsList, error)
	GetNewsById(newsId string) (models.NewsList, error)
	UpdateNewsFile(file models.UploadedFile, newsId string, entity string) (bool, error)
}

type Permissions interface {
	GetMe(token string) (models.UserSingle, error)
}

type S3 interface {
	UploadFile(ctx context.Context, reader io.Reader, bucketName, fileName string, fileSize int64) (minio.UploadInfo, string, error)
}

type Service struct {
	News
	Permissions
	S3
}

func NewService(repos *postgres.Repository, s3 *minio.Client, gHandler *grpcHandler.GrpcHandler) *Service {
	return &Service{
		News:        NewNewsService(repos.News, gHandler),
		Permissions: NewPermissionService(gHandler),
		S3:          NewS3Service(s3),
	}
}
