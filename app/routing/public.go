package routing

import (
	"github.com/gofiber/fiber/v2"
	"dockerized-go-to-heroku/app/controllers"
)

func CreatePublicRoutes(router fiber.Router) {
	demoController := controllers.NewDemoController("A random string passed into demo controller")

	router.Use(func(c *fiber.Ctx) error {
		c.Set("Allow", "GET, POST, PUT")
		return c.Next()
	})

	router.Get("/demo", demoController.ReturnString)
}