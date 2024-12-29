package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anton-ag/javacode/internal/config"
	"github.com/anton-ag/javacode/internal/http"
	"github.com/anton-ag/javacode/internal/repo"
	"github.com/anton-ag/javacode/internal/server"
	"github.com/anton-ag/javacode/internal/service"
	"github.com/anton-ag/javacode/pkg/postgres"
)

func Run(configPath string) {
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := postgres.Connect(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	if err = postgres.InitTable(db); err != nil {
		log.Fatal(err)
		return
	}

	repo := repo.InitRepo(db)
	service := service.InitService(repo)
	handler := http.NewHandler(service)
	server := server.NewServer(cfg, handler.Init())

	go func() {
		if err := server.Run(); err != nil {
			log.Fatalf("Ошибка сервера: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err := server.Stop(ctx); err != nil {
		log.Fatalf("Ошибка остановки сервера: %s\n", err.Error())
	}

	return
}
