package cmd

import (
	"brutalITSMbeNews/internal/postgres"
	"brutalITSMbeNews/internal/s3"
	"brutalITSMbeNews/internal/service"
	"github.com/joho/godotenv"
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
	handler := http.NewHttp(services)

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
