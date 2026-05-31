package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ChernovDanila/org-api/internal/handler"
	"github.com/ChernovDanila/org-api/internal/repository"
	"github.com/ChernovDanila/org-api/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost port=5432 user=postgres password=postgres dbname=org_api sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	log.Println("Database connected!", db)

	deptRepo := repository.NewDepartmentRepository(db)
	deptService := service.NewDepartmentService(deptRepo)
	deptHandler := handler.NewDepartmentHandler(deptService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /departments/", deptHandler.CreateDepartment)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
