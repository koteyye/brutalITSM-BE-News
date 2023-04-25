package service

import (
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	"github.com/koteyye/brutalITSM-BE-News/internal/postgres"
	"github.com/minio/minio-go/v7"
)

type NewsService struct {
	repo   postgres.News
	s3repo *minio.Client
}

func NewNewsService(repo postgres.News, s3repo *minio.Client) *NewsService {
	return &NewsService{repo: repo, s3repo: s3repo}
}

func (n NewsService) CreateNews(news models.News) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsService) UpdateNews(newsId string, news models.News) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsService) DeleteNews(newsId string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsService) GetNewsList() ([]models.NewsList, error) {
	return n.repo.GetNewsList()
}

func (n NewsService) GetNewsById(newsId string) (models.NewsList, error) {
	//TODO implement me
	panic("implement me")
}
