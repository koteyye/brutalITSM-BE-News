package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	"net/http"
)

const (
	newsContentBucket = "newsContent"
	newsImageBucket   = "newsImage"
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

//func (r *Rest) uploadNewsFile(c *gin.Context) {
//	multipartForm, err := c.MultipartForm()
//	fileHeader := multipartForm.File
//
//	file, err :=
//	defer file.Close()
//
//	if err != nil {
//		newErrorResponse(c, http.StatusBadRequest, "Cant open file")
//		return
//	}
//
//	idFile := uuid.New().String()
//	extension := filepath.Ext(fileHeader.Filename)
//
//	newFileName := idFile + extension
//
//	fileSize := fileHeader.Size
//
//	info, mimeType, uploadErr := r.services.UploadFile(c, file, newsContentBucket, newFileName, fileSize)
//
//	if uploadErr != nil {
//		newErrorResponse(c, http.StatusInternalServerError, uploadErr.Error())
//	}
//
//	entity, err := c.MultipartForm("entity")
//
//}
