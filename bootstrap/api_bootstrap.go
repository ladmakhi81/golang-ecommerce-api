package bootstrap

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/auth"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/cart"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/category"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	errorhandling "github.com/ladmakhi81/golang-ecommerce-api/internal/common/error_handling"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/validation"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/order"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/payment"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/product"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/transaction"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/user"
	vendorincome "github.com/ladmakhi81/golang-ecommerce-api/internal/vendor_income"
	pkgemail "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/service"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/logger"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
	"go.uber.org/dig"
)

type AppServerModule struct {
	authModule         auth.AuthModule
	categoryModule     category.CategoryModule
	productModule      product.ProductModule
	userModule         user.UserModule
	cartModule         cart.CartModule
	orderModule        order.OrderModule
	paymentModule      payment.PaymentModule
	transactionModule  transaction.TransactionModule
	vendorIncomeModule vendorincome.VendorIncomeModule
}

type AppServer struct {
	config    config.MainConfig
	app       *echo.Echo
	container *dig.Container
	logger    logger.ILogger
	AppServerModule
}

func NewAppServer(
	config config.MainConfig,
	container *dig.Container,
	logger logger.ILogger,
) AppServer {
	return AppServer{
		config:    config,
		app:       echo.New(),
		container: container,
		logger:    logger,
	}
}

func (appServer *AppServer) configAppServer() {
	mainConfig := config.NewMainConfig()
	mainConfig.LoadConfigs()

	appServer.app.Use(
		middleware.RequestLoggerWithConfig(
			middleware.RequestLoggerConfig{
				LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
					appServer.logger.Info(
						fmt.Sprintf("REQUEST:%s, STATUS:%d", c.Request().URL, c.Response().Status),
					)
					return nil
				},
			},
		),
	)
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
	apiPath := appServer.app.Group("/api/v1")

	appServer.configAppServer()
	appServer.loadModules(apiPath)
	appServer.runModules()

	log.Fatalln(appServer.app.Start(port))
}

func (appServer AppServer) loadBasicDependency() {
	appServer.container.Provide(translations.NewTranslation)
	appServer.container.Provide(events.NewEventsContainer)
	appServer.container.Provide(storage.NewStorage)
	appServer.container.Provide(pkgemail.NewEmailService)
	appServer.container.Provide(utils.NewUtil)
	appServer.container.Provide(func() config.MainConfig {
		return appServer.config
	})
}

func (appServer *AppServer) loadModules(apiPath *echo.Group) {
	appServer.loadBasicDependency()

	appServer.authModule = auth.NewAuthModule(appServer.container, apiPath)
	appServer.authModule.LoadModule()

	appServer.categoryModule = category.NewCategoryModule(appServer.container, apiPath)
	appServer.categoryModule.LoadModule()

	appServer.productModule = product.NewProductModule(appServer.container, apiPath)
	appServer.productModule.LoadModule()

	appServer.userModule = user.NewUserModule(appServer.container, apiPath)
	appServer.userModule.LoadModule()

	appServer.cartModule = cart.NewCartModule(appServer.container, apiPath)
	appServer.cartModule.LoadModule()

	appServer.orderModule = order.NewOrderModule(appServer.container, apiPath)
	appServer.orderModule.LoadModule()

	appServer.paymentModule = payment.NewPaymentModule(appServer.container, apiPath)
	appServer.paymentModule.LoadModule()

	appServer.transactionModule = transaction.NewTransactionModule(appServer.container, apiPath)
	appServer.transactionModule.LoadModule()

	appServer.vendorIncomeModule = vendorincome.NewVendorIncomeModule(appServer.container, apiPath)
	appServer.vendorIncomeModule.LoadModule()
}

func (appServer AppServer) runModules() {
	appServer.authModule.Run()
	appServer.categoryModule.Run()
	appServer.productModule.Run()
	appServer.userModule.Run()
	appServer.cartModule.Run()
	appServer.orderModule.Run()
	appServer.paymentModule.Run()
	appServer.transactionModule.Run()
	appServer.vendorIncomeModule.Run()
}
