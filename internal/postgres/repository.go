package postgres

import (
	"brutalITSMbeNews/internal/models"
	"github.com/jmoiron/sqlx"
)

type News interface {
	CreateNews(news models.News) (string, error)
	UpdateNews(newsId string, news models.News) (string, error)
	DeleteNews(newsId string) (bool, error)
	GetNewsList() ([]models.NewsList, error)
	GetNewsById(newsId string) (models.NewsList, error)
}

type Repository struct {
	News
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		News: NewNewsPostgres(db),
	}
}
