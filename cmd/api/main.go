package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/auth"
	authservice "github.com/ladmakhi81/golang-ecommerce-api/internal/auth/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/category"
	categoryrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/category/repository"
	categoryservice "github.com/ladmakhi81/golang-ecommerce-api/internal/category/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	errorhandling "github.com/ladmakhi81/golang-ecommerce-api/internal/common/error_handling"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/validation"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/user"
	userrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/user/repository"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
	pkgemail "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/service"
)

func main() {
	mainConfig := config.NewMainConfig()
	mainConfig.LoadConfigs()

	storage := storage.NewStorage(mainConfig)

	validator := validation.NewInputValidator()

	port := mainConfig.GetAppPort()
	server := echo.New()

	server.Validator = validator
	server.HTTPErrorHandler = errorhandling.GlobalErrorHandling

	server.Use(middleware.Logger())
	server.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{"*"},
			},
		),
	)

	apiRoute := server.Group("/api/v1")

	// repositories
	userRepo := userrepository.NewUserRepository(storage)
	categoryRepo := categoryrepository.NewCategoryRepository(storage)

	// services
	emailService := pkgemail.NewEmailService(mainConfig)
	jwtService := authservice.NewJwtService(mainConfig)
	userService := userservice.NewUserService(userRepo, emailService)
	authService := authservice.NewAuthService(userService, jwtService, emailService)
	categoryService := categoryservice.NewCategoryService(categoryRepo)

	authRouter := auth.NewAuthRouter(apiRoute, authService)
	authRouter.SetupRouter()

	usersRouter := user.NewUserRouter(apiRoute, userService, mainConfig)
	usersRouter.SetupRouter()

	categoryRouter := category.NewCategoryRouter(apiRoute, mainConfig, categoryService)
	categoryRouter.SetupRouter()

	log.Println("the server is running")

	log.Fatalln(server.Start(port))
}
