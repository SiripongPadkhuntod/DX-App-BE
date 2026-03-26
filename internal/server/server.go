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
	"github.com/youruser/dexter-transport/docs"
	"gopkg.in/yaml.v3"
)

type Server struct {
	engine *gin.Engine
}

// Config yaml struct matches dev.yaml format
type DevConfig struct {
	Env []struct {
		Name  string `yaml:"name"`
		Value string `yaml:"value"`
	} `yaml:"env"`
}

func loadConfig() {
	appEnv := os.Getenv("APP_ENV")

	switch appEnv {
	case "dev", "nonprod", "development":
		log.Printf("Loading configuration from dev.yaml (APP_ENV=%s)", appEnv)
		data, err := os.ReadFile("config/dev.yaml")
		if err != nil {
			log.Fatalf("Failed to read dev.yaml: %v", err)
		}
		var config DevConfig
		if err := yaml.Unmarshal(data, &config); err != nil {
			log.Fatalf("Failed to parse dev.yaml: %v", err)
		}
		for _, item := range config.Env {
			// Don't overwrite env vars already set by the host (e.g. Render dashboard)
			if os.Getenv(item.Name) == "" {
				os.Setenv(item.Name, item.Value)
			}
		}
	default:
		// Local development: load local.env + secret.env
		log.Println("Loading configuration from local.env")
		if err := godotenv.Load("config/local.env"); err != nil {
			log.Println("Warning: local.env file not found, using default environment variables")
		}
		if err := godotenv.Load("config/secret.env"); err != nil {
			log.Println("Warning: secret.env file not found, using default environment variables")
		}
	}
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
	// 0. Load Environment Variables dynamically
	loadConfig()

	// 1. Configure Swagger based on Environment Variables
	if host := os.Getenv("API_DOCS_HOST"); host != "" {
		docs.SwaggerInfo.Host = host
	}
	if schema := os.Getenv("API_DOCS_SCHEMA"); schema != "" {
		docs.SwaggerInfo.Schemes = []string{schema}
	} else {
		docs.SwaggerInfo.Schemes = []string{"http"} // Default to http for local
	}

	// 2. Initial Database Connection with new infrastructure client
	db, err := db_client.NewPostgresClient(db_client.PostgresConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "user"),
		Password: getEnv("DB_PASSWORD", "password"),
		DBName:   getEnv("DB_NAME", "dexter_db"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
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
