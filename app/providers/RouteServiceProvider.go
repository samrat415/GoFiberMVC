package providers

import (
	"eventsapp/app/routes"
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers routes and middleware for your application
func RegisterRoutes(app *fiber.App) {
	// Import route files
	importRoutes := []func(*fiber.App){
		routes.RegisterWebRoutes,
	}
	// Register routes
	for _, importRoute := range importRoutes {
		importRoute(app)
	}
}
func NotFoundHandler(c *fiber.Ctx) error {
	// Custom 404 not found message
	url := c.OriginalURL()
	message := "Oops! The requested route '" + url + "' was not found."

	// Return a custom 404 response
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}
