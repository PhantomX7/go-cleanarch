package main

import (
	"context"
	"fmt"
	"github.com/PhantomX7/go-cleanarch/util/validators"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PhantomX7/go-cleanarch/app/api/middleware"
	"github.com/PhantomX7/go-cleanarch/app/api/server"

	authorHTTP "github.com/PhantomX7/go-cleanarch/author/delivery/http"
	authorRepo "github.com/PhantomX7/go-cleanarch/author/repository/mysql"
	authorUsecase "github.com/PhantomX7/go-cleanarch/author/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/subosito/gotenv"
)

func main() {
	loadEnv()
	db := setupDatabase()
	// init custom validator
	validators.NewValidator(db)

	authorHandler := resolveAuthorHandler(db)
	startServer(
		authorHandler,
	)
}

func startServer(handlers ...server.Handler) {

	mConfig := middleware.Config{
	}
	m := middleware.New(mConfig)

	h := server.BuildHandler(m, handlers...)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler:      h,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func(s *http.Server) {
		log.Printf("api is available at %s\n", s.Addr)
		if serr := s.ListenAndServe(); serr != http.ErrServerClosed {
			log.Fatal(serr)
		}
	}(s)

	<-sigChan

	log.Println("\nSignal received. Waiting for readiness check...")

	//  wait 15s (kube readiness period second) before shutting down http
	sleep()

	log.Println("\nShutting down the api...")

	err := s.Shutdown(context.Background())
	if err != nil {
		log.Fatal("Something wrong when stopping server : ", err)
		return
	}

	log.Println("api gracefully stopped")
}

func loadEnv() {
	err := gotenv.Load()

	if err != nil {
		panic(err)
	}
}

func setupDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	return db
}

func sleep() {
	waitingTime := time.Duration(15)
	if os.Getenv("APP_ENV") != "production" {
		waitingTime = 1
	}
	time.Sleep(waitingTime * time.Second)
}

func resolveAuthorHandler(db *gorm.DB) server.Handler {
	authorR := authorRepo.NewAuthorRepository(db)
	authorUC := authorUsecase.NewAuthorUsecase(authorR)
	return authorHTTP.NewAuthorHandler(authorUC)
}
