package cmd

import (
	"github.com/joho/godotenv"
	"github.com/koteyye/brutalITSM-BE-News/internal/postgres"
	rest "github.com/koteyye/brutalITSM-BE-News/internal/rest"
	"github.com/koteyye/brutalITSM-BE-News/internal/s3"
	"github.com/koteyye/brutalITSM-BE-News/internal/service"
	brutalitsm "github.com/koteyye/brutalITSM-BE-News/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables %s", err.Error())
	}
	// Postgres Client
	db, _ := postgres.InitPostgres()
	// Minio Client
	minio, _ := s3.InitMinioClient()

	//init internal
	repos := postgres.NewRepository(db)
	services := service.NewService(repos, minio)
	handler := rest.NewRest(services)

	restSrv := new(brutalitsm.Server)
	if err := restSrv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occuped while runing http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("server")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
