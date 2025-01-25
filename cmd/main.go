package main

import (
	"github.com/VictorNevola/config"
	"github.com/VictorNevola/internal/domain/company"
	"github.com/VictorNevola/internal/domain/promotion"
	"github.com/VictorNevola/internal/domain/user"
	userinpromotion "github.com/VictorNevola/internal/domain/userInPromotion"
	"github.com/VictorNevola/internal/infra/adapters/postgresql"
	customvalidator "github.com/VictorNevola/internal/pkg/utils/custom-validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "America/Sao_Paulo",
		Format:     "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		Next: func(c *fiber.Ctx) bool {
			return c.Path() != "/metrics"
		},
	}))

	app.Get("/metrics", monitor.New())

	validator := customvalidator.NewCustomValidator()

	// Dependencies
	db := config.NewDB(&config.DBConfigParams{
		Dsn: "postgres://postgres:postgres@localhost:5432/fidelis?sslmode=disable",
	})

	// Repositories
	companyRepository := postgresql.NewCompanyRepository(db)
	promotionRepository := postgresql.NewPromotionRepository(db)
	userRepository := postgresql.NewUserRepository(db)
	userInPromotionRepository := postgresql.NewUserInPromotionRepository(db)

	// Services
	companyService := company.NewService(&company.ServiceParams{
		CompanyRepository: companyRepository,
	})
	promotionService := promotion.NewService(&promotion.ServiceParams{
		PromotionRepository: promotionRepository,
	})
	userService := user.NewService(&user.ServiceParams{
		UserRepository: userRepository,
	})
	userInPromotionService := userinpromotion.NewService(&userinpromotion.ServiceParams{
		UserInPromotionRepository: userInPromotionRepository,
	})

	// Routers
	company.NewHTTPHandler(&company.HttpHandlerParams{
		App:       app,
		Service:   companyService,
		Validator: validator,
	})
	promotion.NewHTTPHandler(&promotion.HttpHandlerParams{
		App:       app,
		Service:   promotionService,
		Validator: validator,
	})
	user.NewHTTPHandler(&user.HttpHandlerParams{
		App:       app,
		Service:   userService,
		Validator: validator,
	})
	userinpromotion.NewHTTPHandler(&userinpromotion.HttpHandlerParams{
		App:       app,
		Service:   userInPromotionService,
		Validator: validator,
	})

	routes := app.GetRoutes()

	for _, route := range routes {
		log.Info(route.Method, route.Path)
	}

	app.Listen(":3000")
}
