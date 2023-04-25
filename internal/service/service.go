package service

import (
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	"github.com/koteyye/brutalITSM-BE-News/internal/postgres"
	"github.com/minio/minio-go/v7"
)

type News interface {
	CreateNews(news models.News) (string, error)
	UpdateNews(newsId string, news models.News) (string, error)
	DeleteNews(newsId string) (bool, error)
	GetNewsList() ([]models.NewsList, error)
	GetNewsById(newsId string) (models.NewsList, error)
}

type Service struct {
	News
}

func NewService(repos *postgres.Repository, s3 *minio.Client) *Service {
	return &Service{
		News: NewNewsService(repos.News, s3),
	}
}
