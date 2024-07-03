package main

import (
	"craw-finance-report/internal/file"
	"craw-finance-report/internal/http"
	"craw-finance-report/pkg/service"
)

func main() {
	authService := &service.AuthService{}
	handler := &http.Handler{
		FileService:  &file.Service{},
		CrawlService: &service.CrawlService{AuthService: authService},
	}
	server := http.Server{
		Port:    "8080",
		Handler: handler,
	}
	server.Run()
}
