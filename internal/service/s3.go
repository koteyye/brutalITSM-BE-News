package service

import (
	"context"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/minio/minio-go/v7"
	"io"
)

type S3Service struct {
	s3repo *minio.Client
}

func NewS3Service(s3repo *minio.Client) *S3Service {
	return &S3Service{s3repo: s3repo}
}

func (s S3Service) UploadFile(ctx context.Context, reader io.Reader, bucketName, fileName string, fileSize int64) (minio.UploadInfo, string, error) {
	info, err := s.s3repo.PutObject(ctx, bucketName, fileName, reader, fileSize, minio.PutObjectOptions{})

	if err != nil {
		return minio.UploadInfo{}, "", fmt.Errorf("cant upload file to s3")
	}

	mType, err := mimetype.DetectReader(reader)

	return info, mType.String(), nil
}
