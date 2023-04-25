package postgres

import (
	"brutalITSMbeNews/internal/models"
	"github.com/jmoiron/sqlx"
)

type NewsPostgres struct {
	db *sqlx.DB
}

func NewNewsPostgres(db *sqlx.DB) *NewsPostgres {
	return &NewsPostgres{db: db}
}

func (n NewsPostgres) CreateNews(news models.News) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsPostgres) UpdateNews(newsId string, news models.News) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsPostgres) DeleteNews(newsId string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsPostgres) GetNewsList() ([]models.NewsList, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsPostgres) GetNewsById(newsId string) (models.NewsList, error) {
	//TODO implement me
	panic("implement me")
}
