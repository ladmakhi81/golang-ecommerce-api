package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/auth"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/category"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	errorhandling "github.com/ladmakhi81/golang-ecommerce-api/internal/common/error_handling"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/validation"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/product"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/user"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
	"go.uber.org/dig"
)

func main() {

	// mainConfig := config.NewMainConfig()
	// mainConfig.LoadConfigs()

	// translation := translations.NewTranslation()

	// validator := validation.NewInputValidator()

	// port := mainConfig.GetAppPort()
	// server := echo.New()

	// server.Validator = validator
	// server.HTTPErrorHandler = errorhandling.GlobalErrorHandling

	// server.Use(middleware.Logger())
	// server.Use(
	// 	middleware.CORSWithConfig(
	// 		middleware.CORSConfig{
	// 			AllowOrigins: []string{"*"},
	// 			AllowMethods: []string{"*"},
	// 		},
	// 	),
	// )

	// apiRoute := server.Group("/api/v1")

	// diContainer := dig.New()

	// // must remove it
	// diContainer.Provide(translations.NewTranslation)
	// diContainer.Provide(userservice.NewUserService)
	// diContainer.Provide(events.NewEventsContainer2)
	// diContainer.Provide(userrepository.NewUserRepository)
	// diContainer.Provide(storage.NewStorage)
	// diContainer.Provide(config.NewMainConfig)

	// authModule := auth.NewAuthModule(diContainer, apiRoute)
	// authModule.Bootstrap()

	// // event container
	// eventContainer := events.NewEventsContainer()

	// storage := storage.NewStorage(mainConfig)

	// // repositories
	// userRepo := userrepository.NewUserRepository(storage)
	// categoryRepo := categoryrepository.NewCategoryRepository(storage)
	// productRepo := productrepository.NewProductRepository(storage)
	// productPriceRepo := productrepository.NewProductPriceRepository(storage)
	// cartRepo := cartrepository.NewCartRepository(storage)
	// orderRepo := orderrepository.NewOrderRepository(storage)
	// paymentRepo := paymentrepository.NewPaymentRepository(storage)
	// transactionRepo := transactionrepository.NewTransactionRepository(storage)
	// vendorIncomeRepo := vendorincomerepository.NewVendorIncomeRepository(storage)
	// userAddressRepo := userrepository.NewUserAddressRepository(storage)

	// // services
	// zarinpalService := pkgzarinpalservice.NewZarinpalService(mainConfig)
	// emailService := pkgemail.NewEmailService(mainConfig)
	// // jwtService := authservice.NewJwtService(mainConfig)
	// userService := userservice.NewUserService(userRepo, translation, &eventContainer)
	// // authService := authservice.NewAuthService(userService, jwtService, translation, &eventContainer)
	// categoryService := categoryservice.NewCategoryService(categoryRepo, mainConfig, translation)
	// transactionService := transactionservice.NewTransactionService(transactionRepo)
	// productService := productservice.NewProductService(
	// 	userService,
	// 	categoryService,
	// 	productRepo,
	// 	translation,
	// 	&eventContainer,
	// )
	// productPriceService := productservice.NewProductPriceService(
	// 	productService,
	// 	productPriceRepo,
	// 	translation,
	// )
	// cartService := cartservice.NewCartService(cartRepo, productService, productPriceService, userService, translation)
	// paymentService := paymentservice.NewPaymentService(paymentRepo, zarinpalService, transactionService, &eventContainer, translation)
	// userAddressService := userservice.NewUserAddressService(userAddressRepo, userService, translation)
	// orderService := orderservice.NewOrderService(userService, orderRepo, cartService, productService, paymentService, userAddressService, translation)
	// vendorIncomeService := vendorincomeservice.NewVendorIncomeService(vendorIncomeRepo, orderService, transactionService)

	// // event subscribers
	// vendorIncomeEventSubscriber := vendorincomeevent.NewVendorIncomeEventsSubscriber(vendorIncomeService)
	// orderEventSubscriber := orderevent.NewOrderEventsSubscriber(orderService, emailService, translation)
	// userEventSubscriber := userevent.NewUserEventsSubscriber(emailService, translation)
	// productEventSubscriber := productevent.NewProductEventsSubscriber(emailService, translation)

	// // event containers
	// vendorIncomeEventContainer := vendorincomeevent.NewVendorIncomeEventsContainer(&eventContainer, vendorIncomeEventSubscriber)
	// vendorIncomeEventContainer.RegisterEvents()

	// productEventContainer := productevent.NewProductEventsContainer(&eventContainer, productEventSubscriber)
	// productEventContainer.RegisterEvents()

	// userEventContainer := userevent.NewUserEventsContainer(&eventContainer, userEventSubscriber)
	// userEventContainer.RegisterEvents()

	// orderEventContainer := orderevent.NewOrderEventsContainer(&eventContainer, orderEventSubscriber)
	// orderEventContainer.RegisterEvents()

	// // authRouter := auth.NewAuthRouter(apiRoute, authService, translation)
	// // authRouter.SetupRouter()

	// usersRouter := user.NewUserRouter(apiRoute, userService, userAddressService, mainConfig, translation)
	// usersRouter.SetupRouter()

	// categoryRouter := category.NewCategoryRouter(apiRoute, mainConfig, categoryService, translation)
	// categoryRouter.SetupRouter()

	// productRouter := product.NewProductRouter(apiRoute, mainConfig, productService, productPriceService, translation)
	// productRouter.SetupRouter()

	// cartRouter := cart.NewCartRouter(apiRoute, mainConfig, cartService, translation)
	// cartRouter.Setup()

	// orderRouter := order.NewOrderRouter(apiRoute, mainConfig, orderService, translation)
	// orderRouter.SetupRouter()

	// paymentRouter := payment.NewPaymentRouter(apiRoute, mainConfig, paymentService, translation)
	// paymentRouter.SetupRouter()

	// transactionRouter := transaction.NewTransactionRouter(apiRoute, mainConfig, transactionService)
	// transactionRouter.Setup()

	// log.Println("the server is running")

	// log.Fatalln(server.Start(port))

	// server := configAppServer()
	// apiRoute := server.Group("/api/v1")

	mainConfig := config.NewMainConfig()
	mainConfig.LoadConfigs()

	appServer := NewAppServer(mainConfig)
	appServer.ConfigAppServer()

	diContainer := dig.New()

	apiRoute := appServer.GetApiPath()

	diContainer.Provide(translations.NewTranslation)
	diContainer.Provide(events.NewEventsContainer2)
	diContainer.Provide(storage.NewStorage)
	diContainer.Provide(utils.NewUtil)
	diContainer.Provide(func() config.MainConfig {
		return mainConfig
	})

	authModule := auth.NewAuthModule(diContainer, apiRoute)
	authModule.LoadModule()

	categoryModule := category.NewCategoryModule(diContainer, apiRoute)
	categoryModule.LoadModule()

	productModule := product.NewProductModule(diContainer, apiRoute)
	productModule.LoadModule()

	userModule := user.NewUserModule(diContainer, apiRoute)
	userModule.LoadModule()

	authModule.Run()
	userModule.Run()
	productModule.Run()
	categoryModule.Run()
	appServer.StartApp()
}

type AppServer struct {
	config config.MainConfig
	app    *echo.Echo
}

func NewAppServer(config config.MainConfig) AppServer {
	return AppServer{config: config, app: echo.New()}
}

func (appServer *AppServer) ConfigAppServer() {
	appServer.app.Use(middleware.Logger())
	appServer.app.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{"*"},
			},
		),
	)
	appServer.app.Validator = validation.NewInputValidator()
	appServer.app.HTTPErrorHandler = errorhandling.GlobalErrorHandling
}

func (appServer *AppServer) StartApp() {
	port := appServer.config.GetAppPort()
	log.Fatalln(appServer.app.Start(port))
}

func (appServer AppServer) GetApiPath() *echo.Group {
	return appServer.app.Group("/api/v1")
}

// func diUser(container *dig.Container) {
// 	container.Provide(userservice.NewUserService)
// 	container.Provide(userservice.NewUserAddressService)
// 	container.Provide(userrepository.NewUserRepository)
// 	container.Provide(userrepository.NewUserAddressRepository)
// }

// func diAuth(container *dig.Container) {
// 	container.Provide(auth.NewAuthHandler)
// 	container.Provide(authservice.NewAuthService)
// 	container.Provide(authservice.NewJwtService)
// }
