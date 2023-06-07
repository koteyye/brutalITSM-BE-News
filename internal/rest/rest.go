package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/koteyye/brutalITSM-BE-News/internal/service"
	"time"
)

type Rest struct {
	services *service.Service
}

func NewRest(services *service.Service) *Rest {
	return &Rest{services: services}
}

func (r *Rest) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"x-requested-with, Content-Type, origin, authorization, accept, x-access-token"},
		ExposeHeaders:    []string{},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api", r.getMe)
	{
		news := api.Group("/news", r.setNewsRead, r.checkPermission)
		{
			news.GET("/newsList", r.getNews)
			news.GET("/news/:id")
		}
		newsEditor := api.Group("/newsEditor", r.setNewsWrite, r.checkPermission)
		{
			newsEditor.POST("/uploadNewsFile", r.uploadNewsFile)
			newsEditor.POST("/createNews", r.createNews)
			newsEditor.POST("/updateNews", r.updateNews)
			newsEditor.GET("/myNewsList")
			newsEditor.GET("/myNews/:id")
			newsEditor.DELETE("/deleteNews/:id")
		}
	}
	return router
}
