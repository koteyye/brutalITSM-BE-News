package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
)

type News interface {
	CreateNews(news models.News, userId string) (string, error)
	UpdateNews(news models.News, userId string) (bool, error)
	DeleteNews(newsId string) (bool, error)
	GetNewsList() ([]models.NewsList, error)
	GetNewsById(newsId string) (models.NewsList, error)
	UploadNewsFile(fileId string) (string, error)
	GetNewsFile(newsId string, entity string) (models.FileInput, error)
	CreateNewsFile(file models.UploadedFile, entity string) (string, error)
	UpdateNewsRelation(newsId string, fileId string, entity string) (bool, error)
	DeleteNewsFile(fileId string) (bool, error)
}

type Repository struct {
	News
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		News: NewNewsPostgres(db),
	}
}
