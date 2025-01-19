package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/auth"
	authservice "github.com/ladmakhi81/golang-ecommerce-api/internal/auth/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/cart"
	cartrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/repository"
	cartservice "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/category"
	categoryrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/category/repository"
	categoryservice "github.com/ladmakhi81/golang-ecommerce-api/internal/category/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	errorhandling "github.com/ladmakhi81/golang-ecommerce-api/internal/common/error_handling"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/validation"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/order"
	orderevent "github.com/ladmakhi81/golang-ecommerce-api/internal/order/event"
	orderrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/order/repository"
	orderservice "github.com/ladmakhi81/golang-ecommerce-api/internal/order/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/payment"
	paymentrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/repository"
	paymentservice "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/product"
	productrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/product/repository"
	productservice "github.com/ladmakhi81/golang-ecommerce-api/internal/product/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/transaction"
	transactionrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/repository"
	transactionservice "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/user"
	userrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/user/repository"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
	vendorincomeevent "github.com/ladmakhi81/golang-ecommerce-api/internal/vendor_income/event"
	vendorincomerepository "github.com/ladmakhi81/golang-ecommerce-api/internal/vendor_income/repository"
	vendorincomeservice "github.com/ladmakhi81/golang-ecommerce-api/internal/vendor_income/service"
	pkgemail "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/service"
	pkgzarinpalservice "github.com/ladmakhi81/golang-ecommerce-api/pkg/zarinpal/service"
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

	// event container
	eventContainer := events.NewEventsContainer()

	// repositories
	userRepo := userrepository.NewUserRepository(storage)
	categoryRepo := categoryrepository.NewCategoryRepository(storage)
	productRepo := productrepository.NewProductRepository(storage)
	productPriceRepo := productrepository.NewProductPriceRepository(storage)
	cartRepo := cartrepository.NewCartRepository(storage)
	orderRepo := orderrepository.NewOrderRepository(storage)
	paymentRepo := paymentrepository.NewPaymentRepository(storage)
	transactionRepo := transactionrepository.NewTransactionRepository(storage)
	vendorIncomeRepo := vendorincomerepository.NewVendorIncomeRepository(storage)

	// services
	zarinpalService := pkgzarinpalservice.NewZarinpalService(mainConfig)
	emailService := pkgemail.NewEmailService(mainConfig)
	jwtService := authservice.NewJwtService(mainConfig)
	userService := userservice.NewUserService(userRepo, emailService)
	authService := authservice.NewAuthService(userService, jwtService, emailService)
	categoryService := categoryservice.NewCategoryService(categoryRepo)
	transactionService := transactionservice.NewTransactionService(transactionRepo)
	productService := productservice.NewProductService(userService, categoryService, productRepo, emailService)
	productPriceService := productservice.NewProductPriceService(productService, productPriceRepo)
	cartService := cartservice.NewCartService(cartRepo, productService, productPriceService, userService)
	paymentService := paymentservice.NewPaymentService(paymentRepo, zarinpalService, transactionService, &eventContainer)
	orderService := orderservice.NewOrderService(userService, orderRepo, cartService, productService, paymentService, emailService)
	vendorIncomeService := vendorincomeservice.NewVendorIncomeService(vendorIncomeRepo, orderService, transactionService)

	// event subscribers
	vendorIncomeEventSubscriber := vendorincomeevent.NewVendorIncomeEventsSubscriber(vendorIncomeService)
	orderEventSubscriber := orderevent.NewOrderEventsSubscriber(orderService)

	// event containers
	vendorIncomeEventContainer := vendorincomeevent.NewVendorIncomeEventsContainer(&eventContainer, vendorIncomeEventSubscriber)
	vendorIncomeEventContainer.RegisterEvents()

	orderEventContainer := orderevent.NewOrderEventsContainer(&eventContainer, orderEventSubscriber)
	orderEventContainer.RegisterEvents()

	authRouter := auth.NewAuthRouter(apiRoute, authService)
	authRouter.SetupRouter()

	usersRouter := user.NewUserRouter(apiRoute, userService, mainConfig)
	usersRouter.SetupRouter()

	categoryRouter := category.NewCategoryRouter(apiRoute, mainConfig, categoryService)
	categoryRouter.SetupRouter()

	productRouter := product.NewProductRouter(apiRoute, mainConfig, productService, productPriceService)
	productRouter.SetupRouter()

	cartRouter := cart.NewCartRouter(apiRoute, mainConfig, cartService)
	cartRouter.Setup()

	orderRouter := order.NewOrderRouter(apiRoute, mainConfig, orderService)
	orderRouter.SetupRouter()

	paymentRouter := payment.NewPaymentRouter(apiRoute, mainConfig, paymentService)
	paymentRouter.SetupRouter()

	transactionRouter := transaction.NewTransactionRouter(apiRoute, mainConfig, transactionService)
	transactionRouter.Setup()

	log.Println("the server is running")

	log.Fatalln(server.Start(port))
}
