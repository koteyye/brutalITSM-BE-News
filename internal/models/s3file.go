package models

import "mime/multipart"

type NewsFiles struct {
	MultipartForm *multipart.Form
	BucketName    string
}

type UploadedFile struct {
	MimeType   string
	BucketName string
	FileName   string
}
