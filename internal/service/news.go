package service

import (
	"database/sql"
	"fmt"
	grpcHandler "github.com/koteyye/brutalITSM-BE-News/internal/grpc"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	"github.com/koteyye/brutalITSM-BE-News/internal/postgres"
	"github.com/minio/minio-go/v7"
	"golang.org/x/exp/maps"
)

type NewsService struct {
	repo     postgres.News
	s3repo   *minio.Client
	gHandler *grpcHandler.GrpcHandler
}

func NewNewsService(repo postgres.News, gHandler *grpcHandler.GrpcHandler) *NewsService {
	return &NewsService{repo: repo, gHandler: gHandler}
}

func (n NewsService) CreateNews(news models.News, userId string) (string, error) {
	return n.repo.CreateNews(news, userId)
}

func (n NewsService) UpdateNews(news models.News, userId string) (bool, error) {
	return n.repo.UpdateNews(news, userId)
}

func (n NewsService) UpdateNewsFile(file models.UploadedFile, newsId string, entity string) (bool, error) {

	s3File, err := n.repo.GetNewsFile(newsId, entity)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		return n.createNewsFileAndUpdateRelation(file, entity, newsId)
	default:
		return false, err
	}

	_, err2 := n.createNewsFileAndUpdateRelation(file, entity, newsId)
	if err2 != nil {
		return false, err2
	}
	return n.repo.DeleteNewsFile(s3File.Id)
}

func (n NewsService) createNewsFileAndUpdateRelation(file models.UploadedFile, entity string, newsId string) (bool, error) {
	fileId, err := n.repo.CreateNewsFile(file, entity)
	if err != nil {
		return false, err
	}
	result, err2 := n.repo.UpdateNewsRelation(newsId, fileId, entity)
	if err2 != nil {
		return false, err2
	}
	return result, err2
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
