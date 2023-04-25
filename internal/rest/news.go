package rest

import (
	"github.com/gin-gonic/gin"
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
