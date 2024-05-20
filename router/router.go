package router

import (
	"app/handler"
	"app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)
	///api/auth/login
	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// Group
	group := api.Group("/group")
	group.Post("/", handler.CreateGroup)
	group.Get("/", handler.GetGroups)
	group.Get("/:id", handler.GetGroup)
	group.Patch("/:id", handler.UpdateGroup)
	group.Delete("/:id", handler.DeleteGroup)

	// Note
	note := api.Group("/note")
	note.Get("/:id", handler.GetNote)
	note.Get("/", handler.GetNotes)
	note.Post("/", handler.CreateNote)
	note.Patch("/:id", handler.UpdateNote)
	note.Delete("/:id", handler.DeleteNote)

	// User
	user := api.Group("/user")
	user.Post("/auth", handler.AuthUser)
	user.Get("/:id", handler.GetUser)
	user.Get("/", handler.GetUsers)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser)
}
