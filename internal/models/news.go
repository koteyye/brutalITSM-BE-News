package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

type News struct {
	Id           string `json:"-" db:"id"`
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description"`
	PreviewImage string `json:"previewImage"`
	ContentFile  string `json:"contentFile"`
}

type NewsList struct {
	Id           string   `json:"id" db:"id"`
	Title        string   `json:"title" db:"title"`
	Description  string   `json:"description" db:"description"`
	CreatedAt    NullTime `json:"createdAt" db:"created_at"`
	UpdatedAt    NullTime `json:"updatedAt" db:"updated_at"`
	DeletedAt    NullTime `json:"deletedAt" db:"deleted_at"`
	UserCreated  string   `json:"userCreated" db:"user_created"`
	UserUpdated  string   `json:"userUpdated" db:"user_updated"`
	UserDeleted  *string  `json:"userDeleted" db:"user_deleted"`
	State        string   `json:"state" db:"state"`
	PreviewImage string   `json:"previewImage" db:"preview_image"`
	ContentFile  string   `json:"contentFile" db:"content_file"`
}

type Files struct {
	Id         string   `json:"id" db:"id"`
	MimeType   string   `json:"mimeType" db:"mime_type"`
	BucketName string   `json:"bucketName" db:"bucket_name"`
	FileName   string   `json:"fileName" db:"file_name"`
	CreatedAt  NullTime `json:"createdAt" db:"created_at"`
	UpdatedAt  NullTime `json:"updatedAt" db:"updated_at"`
	DeletedAt  NullTime `json:"deletedAt" db:"deleted_at"`
}

func (f Files) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		err := json.Unmarshal(v, &f)
		if err != nil {
			return err
		}
		return nil
	case string:
		err := json.Unmarshal([]byte(v), &f)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

type ResponseNewsList struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	CreatedAt    NullTime `json:"createdAt"`
	UpdatedAt    NullTime `json:"updatedAt"`
	UserCreated  string   `json:"userCreated"`
	UserUpdated  string   `json:"userUpdated"`
	State        string   `json:"state" db:"state"`
	PreviewImage string   `json:"previewImage"`
	ContentFile  string   `json:"contentFile"`
}

type Avatar struct {
	MimeType   string `json:"mimeType"`
	BucketName string `json:"bucketName"`
	FileName   string `json:"fileName"`
}
