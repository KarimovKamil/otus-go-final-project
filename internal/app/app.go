package main

import (
	"context"
	"fmt"
	"os"

	"github.com/KarimovKamil/otus-go-final-project/internal/config"
	"github.com/KarimovKamil/otus-go-final-project/internal/controller/httpapi"
	"github.com/KarimovKamil/otus-go-final-project/internal/controller/httpapi/handler"
	"github.com/KarimovKamil/otus-go-final-project/internal/repository"
	"github.com/KarimovKamil/otus-go-final-project/internal/repository/client"
	"github.com/KarimovKamil/otus-go-final-project/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	projectConfig := config.Read("./configs/config.yaml")

	psql := client.NewPostgresSQL(projectConfig)
	err := psql.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	defer psql.Close()

	blackListRepo := repository.NewBlackListRepo(psql)
	blackList := service.NewBlackList(blackListRepo)

	whiteListRepo := repository.NewWhiteListRepo(psql)
	whiteList := service.NewWhiteList(whiteListRepo)

	listHandler := handler.NewListHandler(whiteList, blackList)

	authService := service.NewAuthorization(projectConfig, blackList, whiteList)
	authHandler := handler.NewAuthHandler(authService)

	bucketHandler := handler.NewBucketHandler(authService)

	router := httpapi.NewAPIRouter(authHandler, bucketHandler, listHandler)
	router.RegisterRoutes()

	ch := make(chan os.Signal, 1)

	httpServer := httpapi.NewServer(router.GetRouter(), &projectConfig)
	go httpServer.ShutdownService(ch)
	err = httpServer.Start()
	if err != nil {
		fmt.Println(err)
	}
}
