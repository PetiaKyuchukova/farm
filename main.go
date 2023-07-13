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
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "471219011"
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

	//repositories
	querier := db.New(mydb)
	taskRepo := gen.NewTaskRepo(querier)
	cowRepo := gen.NewCowRepo(querier)
	milkRepo := gen.NewMilkRepo(querier)
	inseminationRepo := gen.NewInseminationRepo(querier)
	pregnancyRepo := gen.NewPregnancyRepo(querier)

	//use-cases
	cowUc := usecase.NewCowUC(cowRepo, pregnancyRepo, inseminationRepo)
	taskUC := usecase.NewTaskUC(taskRepo)
	milkUC := usecase.NewMilkUC(milkRepo)

	//handlers
	cowHandler := handlers.NewCowHandler(cowUc)
	taskHandler := handlers.NewTaskHandler(taskUC)
	milkHandler := handlers.NewMilkHandler(milkUC)

	worker := worker.NewWorker(taskUC, cowUc)
	worker.Schedule(ctx, "*/5 * * * *")

	router := gin.Default()
	router.LoadHTMLGlob("src/pages/*.gohtml")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}

	router.Use(cors.New(config))
	router.GET("/done", cowHandler.LivenessHandler)
	router.GET("/template/:id", cowHandler.RenderTemplate)

	router.PUT("/upsert", cowHandler.UpsertCow)
	router.DELETE("/delete/:id", cowHandler.DeleteCow)
	router.GET("/cows", cowHandler.GetAllCows)
	router.GET("/cows/:id", cowHandler.GetCowById)

	router.GET("/tasks", taskHandler.GetTasksByDate)

	router.PUT("/milk", milkHandler.UpsertMilk)
	router.GET("/milk", milkHandler.FetchMilkSeriesInTimeframe)

	router.Run(":9030")
	fmt.Println("Server running on port 3030")

}
