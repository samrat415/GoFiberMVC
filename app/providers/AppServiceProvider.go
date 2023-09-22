package providers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/spf13/viper"
)

// AppProvider performs application bootstrapping
func AppProvider() *fiber.App {
	//Using HTML Engine
	engine := html.New("app/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	// Load the application configuration LoadAppConfig()
	LoadAppConfig()
	// Register routes
	RegisterRoutes(app)
	//404 Handler
	app.Use(NotFoundHandler)
	return app
}

// LoadAppConfig loads the application configuration from the config file
func LoadAppConfig() {
	viper.SetConfigFile("app/config/config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func StorageConfig() {

}
