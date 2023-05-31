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

type NullString struct {
	String string
	Valid  bool
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
	UserCreated  string   `json:"userCreated" db:"user_created"`
	UserUpdated  string   `json:"userUpdated" db:"user_updated"`
	State        string   `json:"state" db:"state"`
	PreviewImage *Files   `json:"previewImage" db:"previewimage"`
	ContentFile  *Files   `json:"contentFile" db:"contentfile"`
}

type Files struct {
	MimeType   string `json:"mimeType" db:"mimeType"`
	BucketName string `json:"bucketName" db:"bucketName"`
	FileName   string `json:"fileName" db:"fileName"`
}

func (f *Files) Scan(val any) error {
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
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UserCreated  *User     `json:"userCreated"`
	UserUpdated  *User     `json:"userUpdated"`
	State        string    `json:"state" db:"state"`
	PreviewImage *Files    `json:"previewImage"`
	ContentFile  *Files    `json:"contentFile"`
}

type AvatarImg struct {
	MimeType   *string `json:"mimeType"`
	BucketName *string `json:"bucketName"`
	FileName   *string `json:"fileName"`
}

type User struct {
	Id       string     `json:"id"`
	FullName string     `json:"fullName"`
	Avatar   *AvatarImg `json:"avatar"`
}
