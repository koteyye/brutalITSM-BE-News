package postgres

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	"github.com/sirupsen/logrus"
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
	var newsList []models.NewsList

	query := sq.Select("*").From("getNewsList()")
	sql, _, err := query.ToSql()
	if err != nil {
		logrus.Fatalf("SQL query not builde %v", err)
	}
	err1 := n.db.Select(&newsList, sql)
	return newsList, err1
}

func (n NewsPostgres) GetNewsById(newsId string) (models.NewsList, error) {
	//TODO implement me
	panic("implement me")
}
