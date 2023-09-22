package routes

import (
	"eventsapp/app/controllers"
	"eventsapp/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

// RegisterWebRoutes registers web routes
func RegisterWebRoutes(app *fiber.App) {
	// Create an instance of the ClientAuthMiddleware
	clientAuthMiddleware := &middlewares.ClientAuthMiddleware{}
	throttleMiddleware := &middlewares.ThrottleMiddleware{}
	eventController := &controllers.EventController{}
	userController := &controllers.UserController{}
	transactionController := &controllers.TransactionController{}

	//Defining Route Group as api
	//api := app.Group("", clientAuthMiddleware.Auth)
	app.Get("", userController.Index)
	api := app.Group("api", clientAuthMiddleware.Auth)
	api.Get("/events", throttleMiddleware.Limit, eventController.ApiIndex)
	app.Get("/transactions", transactionController.Index)
	api.Get("/denomination/list", throttleMiddleware.Limit, eventController.ApiDenominationList)
	api.Get("/event/group/list", throttleMiddleware.Limit, eventController.ApiGroupList)
	api.Get("/event/participant/list", throttleMiddleware.Limit, eventController.ApiParticipantList)
	api.Post("/vote/participant", throttleMiddleware.Limit, eventController.ApiVoteParticipant)
	api.Get("/vote/history", throttleMiddleware.Limit, eventController.ApiVoteHistory)
	api.Get("/vote/history", throttleMiddleware.Limit, eventController.ApiVoteHistory)
	api.Get("/vote/category", throttleMiddleware.Limit, eventController.ApiEventCategory)

}
