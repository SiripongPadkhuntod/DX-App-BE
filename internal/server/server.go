package server

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/youruser/dexter-transport/internal/app/handler"
	"github.com/youruser/dexter-transport/internal/app/port"
	postgres_repository "github.com/youruser/dexter-transport/internal/app/repository/postgres-repository"
	"github.com/youruser/dexter-transport/internal/app/service"
	db_client "github.com/youruser/dexter-transport/internal/infrastructure/db-client"
	"github.com/youruser/dexter-transport/internal/router"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	engine := gin.Default()

	// CORS Configuration
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust in production
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return &Server{
		engine: engine,
	}
}

func (s *Server) Run() {
	// 0. Load Environment Variables
	if err := godotenv.Load("config/local.env"); err != nil {
		log.Println("Warning: .env file not found, using default environment variables")
	}

	// 1. Initial Database Connection with new infrastructure client
	db, err := db_client.NewPostgresClient(db_client.PostgresConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "user"),
		Password: getEnv("DB_PASSWORD", "password"),
		DBName:   getEnv("DB_NAME", "dexter_db"),
	})
	if err != nil {
		log.Fatalf("Failed to initialize database client: %v", err)
	}
	defer db.Close()

	// 2. Initialize Repository
	repo := port.Repository{
		Sql: postgres_repository.NewPostgresRepository(db),
	}

	// 3. Initialize Service
	svc := service.New(repo)

	// 4. Initialize Handler
	h := handler.New(svc)

	router.SetupRouter(s.engine, h)

	port := getEnv("PORT", "8080")
	log.Printf("Server is running on :%s", port)
	s.engine.Run(":" + port)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
