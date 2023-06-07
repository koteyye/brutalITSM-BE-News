package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	"github.com/sirupsen/logrus"
)

const (
	NewsContent      = "content"
	NewsComment      = "comment"
	NewsPreviewImage = "previewImage"
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

func (n NewsPostgres) GetNewsFile(newsId string, entity string) (models.FileInput, error) {
	var repS3File models.FileInput
	query := sq.Select("f.id, f.mime_type, f.bucket_name, f.file_name, f.entity").
		From("news n").
		Join(foundFK(entity)).Where(sq.Eq{"n.id": newsId}).Where(sq.Eq{"f.deleted_at": nil}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		logrus.Fatalf("SQL query not builde %v", err)
	}
	err1 := n.db.Get(&repS3File, sql, args...)
	return repS3File, err1
}

func foundFK(entity string) string {
	switch entity {
	case NewsContent:
		return "files f on n.content_file = f.id"
	case NewsPreviewImage:
		return "files f on n.preview_image = f.id"
	default:
		return "files f on c.image_file = f.id"
	}

}

func (n NewsPostgres) CreateNewsFile(file models.UploadedFile, entity string) (string, error) {
	var id string
	query := sq.Insert("files").
		Columns("mime_type", "bucket_name", "file_name", "entity").
		Values(file.MimeType, file.BucketName, file.FileName, entity).
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

func (n NewsPostgres) UpdateNewsRelation(newsId string, fileId string, entity string) (bool, error) {
	query := sq.Update("news").
		Set(foundRelation(entity), fileId).
		Where(sq.Eq{"id": newsId}).
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

func foundRelation(entity string) string {
	switch entity {
	case NewsContent:
		return "content_file"
	default:
		return "preview_image"
	}
}

func (n NewsPostgres) DeleteNewsFile(fileId string) (bool, error) {
	query := sq.Update("files").
		Set("deleted_at", sq.Expr("Now()")).
		Where(sq.Eq{"id": fileId}).
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
