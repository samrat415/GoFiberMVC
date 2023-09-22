package main

import (
	"eventsapp/app/initializers"
	"eventsapp/app/providers"
	"eventsapp/app/routes"
)

func main() {

	// Initialize the Fiber app using the AppProvider
	app := providers.AppProvider()
	// Register web routes
	routes.RegisterWebRoutes(app)
	initializers.DbConnection()
	initializers.OauthDatabaseConnection()
	// Start the server
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
