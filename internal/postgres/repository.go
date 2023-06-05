package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
)

type News interface {
	CreateNews(news models.News, userId string) (string, error)
	UpdateNews(newsId string, news models.News) (string, error)
	DeleteNews(newsId string) (bool, error)
	GetNewsList() ([]models.NewsList, error)
	GetNewsById(newsId string) (models.NewsList, error)
	UploadNewsFile(fileId string) (string, error)
}

type Repository struct {
	News
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		News: NewNewsPostgres(db),
	}
}
