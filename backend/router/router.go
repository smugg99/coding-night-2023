package router

import (
	"crypto/rand"
	"crypto/rsa"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/smugg99/coding-night-2023/config"
	"github.com/smugg99/coding-night-2023/controllers"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *fiber.App {
	a := fiber.New()

	// Middleware
	a.Use(logger.New())
	a.Use(recover.New())
	a.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	a.Use(cors.New(cors.Config{
		AllowOrigins: config.Conf.SrvConfig.AllowOrigins,
		AllowMethods: "GET, POST, OPTIONS",
		AllowHeaders: "Origin, Host, Content-Type, Accept",
	}))
	a.Get("/monitor", monitor.New(monitor.Config{
		Title: "Backend monitor stats",
	}))
 
    // JWT
    rng := rand.Reader
    var err error
    config.PrivateKey, err = rsa.GenerateKey(rng, 2048)
    if err != nil {
        panic(err)
    }
    jwtMiddleware := jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{
            JWTAlg: jwtware.RS256,
            Key:    config.PrivateKey.Public(),
        },
    })

	// Controller
	api := controllers.Controller{Db: db}

	// Routing
    auth := a.Group("/auth")
    {
        auth.Post("/login", api.Login)
        auth.Post("/register", api.Register)
    }

    incident := a.Group("/incident")
    {
        incident.Get("/", api.Incidents)
        incident.Post("/category", api.CreateCategory)       
        protected := incident.Group("/").Use(jwtMiddleware)
        {
            protected.Post("/", api.CreateIncident)            
        }
    }

	return a
}
