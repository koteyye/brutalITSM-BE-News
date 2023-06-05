package service

import (
	"context"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	grpcHandler "github.com/koteyye/brutalITSM-BE-News/internal/grpc"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	"github.com/koteyye/brutalITSM-BE-News/internal/postgres"
	"github.com/minio/minio-go/v7"
	"golang.org/x/exp/maps"
	"io"
)

type NewsService struct {
	repo     postgres.News
	s3repo   *minio.Client
	gHandler *grpcHandler.GrpcHandler
}

func NewNewsService(repo postgres.News, s3repo *minio.Client, gHandler *grpcHandler.GrpcHandler) *NewsService {
	return &NewsService{repo: repo, s3repo: s3repo, gHandler: gHandler}
}

func (n NewsService) UploadFile(ctx context.Context, reader io.Reader, bucketName, fileName string, fileSize int64) (minio.UploadInfo, string, error) {
	info, err := n.s3repo.PutObject(ctx, bucketName, fileName, reader, fileSize, minio.PutObjectOptions{})

	if err != nil {
		return minio.UploadInfo{}, "", err
	}

	mType, err := mimetype.DetectReader(reader)

	return info, mType.String(), nil
}

func (n NewsService) CreateNews(news models.News, userId string) (string, error) {
	return n.repo.CreateNews(news, userId)
}

func (n NewsService) UpdateNews(news models.News, userId string) (bool, error) {
	return n.repo.UpdateNews(news, userId)
}

func (n NewsService) DeleteNews(newsId string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsService) GetNewsList() ([]models.ResponseNewsList, error) {
	var ids []string
	dbData, err := n.repo.GetNewsList()
	if err != nil {
		return nil, err
	}

	for _, user := range dbData {
		ids = append(ids, user.UserCreated, user.UserUpdated)
	}
	unIds := unique(ids)
	userData, err := n.gHandler.GetUserList(unIds)
	if err != nil {
		return nil, err
	}

	return responseUserList(dbData, userData), nil
}

func (n NewsService) GetNewsById(newsId string) (models.NewsList, error) {
	//TODO implement me
	panic("implement me")
}

func unique[k comparable](arr []k) []k {
	uniqueMap := make(map[k]struct{})

	for _, arrEl := range arr {
		uniqueMap[arrEl] = struct{}{}
	}

	return maps.Keys(uniqueMap)
}

func responseUserList(dbData []models.NewsList, userData []models.UserList) []models.ResponseNewsList {
	dbDataMap := make(map[string]models.NewsList)

	for _, dbs := range dbData {
		dbDataMap[dbs.UserCreated] = dbs
		dbDataMap[dbs.UserUpdated] = dbs
	}

	userDataMap := make(map[string]models.UserList)

	for _, users := range userData {
		userDataMap[users.Id] = users
	}

	resNewsList := make([]models.ResponseNewsList, 0)

	for _, resNews := range dbData {
		resNewsList = append(resNewsList, models.ResponseNewsList{
			Id:          resNews.Id,
			Title:       resNews.Title,
			Description: resNews.Description,
			CreatedAt:   resNews.CreatedAt.Time,
			UpdatedAt:   resNews.UpdatedAt.Time,
			UserCreated: &models.User{
				Id:       fmt.Sprintf("%s", userDataMap[resNews.UserCreated].Id),
				FullName: fmt.Sprintf("%s %s %s", userDataMap[resNews.UserCreated].LastName, userDataMap[resNews.UserCreated].FirstName, userDataMap[resNews.UserCreated].SurName),
				Avatar: &models.AvatarImg{
					BucketName: userDataMap[resNews.UserCreated].BucketName,
					FileName:   userDataMap[resNews.UserCreated].FileName,
					MimeType:   userDataMap[resNews.UserCreated].MimeType,
				},
			},
			UserUpdated: &models.User{
				Id:       fmt.Sprintf("%s", userDataMap[resNews.UserUpdated].Id),
				FullName: fmt.Sprintf("%s %s %s", userDataMap[resNews.UserUpdated].LastName, userDataMap[resNews.UserUpdated].FirstName, userDataMap[resNews.UserUpdated].SurName),
				Avatar: &models.AvatarImg{
					BucketName: userDataMap[resNews.UserUpdated].BucketName,
					FileName:   userDataMap[resNews.UserUpdated].FileName,
					MimeType:   userDataMap[resNews.UserUpdated].MimeType,
				},
			},
			State: resNews.State,
			PreviewImage: &models.Files{
				BucketName: resNews.PreviewImage.BucketName,
				FileName:   resNews.PreviewImage.FileName,
				MimeType:   resNews.PreviewImage.MimeType,
			},
			ContentFile: &models.Files{
				BucketName: resNews.ContentFile.BucketName,
				FileName:   resNews.ContentFile.FileName,
				MimeType:   resNews.ContentFile.MimeType,
			},
		})
	}
	return resNewsList
}
