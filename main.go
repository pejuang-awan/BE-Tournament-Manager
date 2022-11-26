package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pejuang-awan/BE-Tournament-Manager/config"
	"github.com/pejuang-awan/BE-Tournament-Manager/router"
	"github.com/pejuang-awan/BE-Tournament-Manager/utils"
)

func main() {
	config.InitializeConfig()
	utils.InitializeDatabase()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router.InitializeRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		panic("[ERROR] Failed to listen and serve")
	}
}
