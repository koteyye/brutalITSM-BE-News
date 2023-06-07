package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	"github.com/koteyye/brutalITSM-BE-News/internal/postgres"
	"github.com/minio/minio-go/v7"
	"net/http"
)

func (r *Rest) getNews(c *gin.Context) {
	result, err := r.services.GetNewsList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, result)
}

func (r *Rest) createNews(c *gin.Context) {
	userCtx, ok := c.Get("user")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "User not found in context")
	}
	user := userCtx.(models.UserSingle)

	var input models.News

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := r.services.CreateNews(input, user.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (r *Rest) updateNews(c *gin.Context) {
	userCtx, ok := c.Get("user")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "User not found in context")
	}
	user := userCtx.(models.UserSingle)

	var input models.News

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := r.services.UpdateNews(input, user.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"updateResult": result,
	})
}

func (r *Rest) uploadNewsFile(c *gin.Context) {
	multipartForm, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "No file is received")
		return
	}
	entity := multipartForm.Value["entity"]
	if entity[0] == "" {
		newErrorResponse(c, http.StatusBadRequest, "No entity of file")
		return
	}
	newsId := multipartForm.Value["newsId"]
	if newsId[0] == "" {
		newErrorResponse(c, http.StatusBadRequest, "No entity of file")
		return
	}

	var bucketName string

	switch entity[0] {
	case postgres.NewsContent:
		bucketName = "news-content"
		break
	case postgres.NewsPreviewImage:
		bucketName = "news-images"
		break
	case postgres.NewsComment:
		bucketName = "news-comments"
		break
	}

	newsFile := models.NewsFiles{
		MultipartForm: multipartForm,
		BucketName:    bucketName,
	}

	r.uploadFile(c, newsFile)

	uploadedFile, ok := c.Get("uploadFile")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "Uploaded file not fount in context")
	}
	input := uploadedFile.(models.UploadedFile)
	s3Info, ok := c.Get("s3Info")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "s3 info not fount in context")
	}
	info := s3Info.(minio.UploadInfo)

	result, errUpdate := r.services.UpdateNewsFile(input, newsId[0], entity[0])
	if errUpdate != nil {
		newErrorResponse(c, http.StatusTeapot, errUpdate.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"updatedNewsFile": result,
		"FileS3Id":        info.Key,
	})

}
