package rest

import (
	"brutalITSMbeNews/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	api := router.Group("/api")
	{
		news := api.Group("/news")
		{
			news.GET("/newsList")
			news.GET("/news/:id")
		}
		newsEditor := api.Group("/newsEditor")
		{
			newsEditor.POST("/createNews")
			newsEditor.POST("/updateNews/:id")
			newsEditor.POST("/uploadNewFiles/:id")
			newsEditor.GET("/newsList")
			newsEditor.GET("/news/:id")
			newsEditor.DELETE("/deleteNews/:id")
		}
	}
	return router
}
