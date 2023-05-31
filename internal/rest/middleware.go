package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/koteyye/brutalITSM-BE-News/internal/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	requiredPermission  = "RequiredPermission"
)

func (r *Rest) getMe(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid auth header")
		return
	}

	user, err := r.services.GetMe(headerParts[1])
	if err != nil {
		errCode, ok := status.FromError(err)
		if !ok {
			logrus.Fatalf("Error response from GRPC User Service")
		}
		if errCode.Code() == 16 {
			newErrorResponse(c, http.StatusUnauthorized, "Invalid token")
		} else if errCode.Code() == 7 {
			newErrorResponse(c, http.StatusUnauthorized, "token is expired")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
	}

	c.Set("user", user)
}

func (r *Rest) setNewsWrite(c *gin.Context) {
	c.Set(requiredPermission, "newsWrite")
}

func (r *Rest) setNewsRead(c *gin.Context) {
	c.Set(requiredPermission, "newsRead")
}

func (r *Rest) checkPermission(c *gin.Context) {
	userCtx, ok := c.Get("user")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "User not found in context")
	}
	user := userCtx.(models.UserSingle)
	godMode := slices.Contains(user.Roles, "admin")
	if godMode != true {
		reqPermissionCtx, ok := c.Get(requiredPermission)
		if !ok {
			newErrorResponse(c, http.StatusInternalServerError, "Not requiredPermission in context")
		}
		reqPermission := reqPermissionCtx.(string)

		validPermission := slices.Contains(user.Permissions, reqPermission)
		if validPermission != true {
			newErrorResponse(c, http.StatusForbidden, "Not enough rights to view")
		}
	}
}
