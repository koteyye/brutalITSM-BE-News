package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	"net/http"
	"path/filepath"
)

func (r *Rest) uploadFile(c *gin.Context, files models.NewsFiles) {
	filesHeader := files.MultipartForm.File["file"]
	file, err := filesHeader[0].Open()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Cant open file")
	}
	defer file.Close()

	idFile := uuid.New().String()
	extension := filepath.Ext(filesHeader[0].Filename)

	newFileName := idFile + extension

	fileSize := filesHeader[0].Size

	info, mimeType, uploadErr := r.services.UploadFile(c, file, files.BucketName, newFileName, fileSize)
	if uploadErr != nil {
		newErrorResponse(c, http.StatusBadRequest, uploadErr.Error())
	}
	c.Set("uploadFile", models.UploadedFile{
		MimeType:   mimeType,
		BucketName: files.BucketName,
		FileName:   newFileName,
	})
	c.Set("s3Info", info)
}
