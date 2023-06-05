package postgres

import (
	"database/sql"
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

func (n NewsPostgres) UploadNewsFile(fileId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsPostgres) CreateNews(news models.News, userId string) (string, error) {
	var id string
	query := sq.Insert("news").
		Columns("title", "description", "user_created", "user_updated").
		Values(news.Title, news.Description, userId, userId).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar)
	sql, args, err := query.ToSql()
	if err != nil {
		logrus.Fatalf("SQL query not builde %v", err)
	}
	row := n.db.QueryRow(sql, args...)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (n NewsPostgres) UpdateNews(news models.News, userId string) (bool, error) {

	query := sq.Update("news").
		Set("title", news.Title).
		Set("description", news.Description).
		Set("state", nullString(news.State).String).
		Set("preview_image", nullString(news.PreviewImage)).
		Set("content_file", nullString(news.ContentFile)).
		Set("updated_at", sq.Expr("Now()")).
		Set("user_updated", userId).
		Where(sq.Eq{"id": news.Id}).
		PlaceholderFormat(sq.Dollar)
	sql, args, err := query.ToSql()
	if err != nil {
		logrus.Fatalf("SQL query not builde %v", err)
	}
	row := n.db.QueryRow(sql, args...)
	err = row.Err()
	if err != nil {
		return false, err
	}
	return true, nil
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

func nullString(val string) sql.NullString {
	if len(val) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: val,
		Valid:  true,
	}
}
