package main

import (
	"context"
	"database/sql"
	"embed"
	"farm/backend/gen"
	"farm/backend/gen/db"
	"farm/backend/handlers"
	"farm/backend/usecase"
	worker "farm/backend/worker"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0041129115"
	dbname   = "farm"
)

var mydb *sql.DB

//go:embed backend/postgres/migrations/*.sql
var embedMigrations embed.FS

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	mydb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer mydb.Close()

	err = mydb.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(mydb, "backend/postgres/migrations"); err != nil {
		panic(err)
	}
	ctx := context.Background()

	querier := db.New(mydb)
	repo := gen.NewFarmRepo(querier)
	notificationRepo := gen.NewNotificationRepo(querier)
	cowUc := usecase.NewCowUC(repo)
	notificationUC := usecase.NewNotificationUC(notificationRepo)

	cowas, err := cowUc.GetAllCows(context.Background())
	fmt.Println(cowas)
	fmt.Println("err", err)

	handler := handlers.NewHandler(cowUc)
	userHandler := handlers.NewUserHandler()

	worker := worker.NewWorker(notificationUC, cowUc)
	worker.Schedule(ctx, "*/5 * * * *")

	router := gin.Default()
	//publicRoutes := router.Group("/profile")
	router.PUT("/upsert", handler.UpsertCow)
	router.PUT("/farmer", userHandler.Register)
	router.DELETE("/delete/:id", handler.DeleteCow)
	router.GET("/cows.sql", handler.GetAllCows)
	router.GET("/cows.sql/:id", handler.GetCowById)
	router.GET("/done", handler.LivenessHandler)

	router.Run(":9030")
	fmt.Println("Server running on port 3030")

}
