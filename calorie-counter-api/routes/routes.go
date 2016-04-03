package routes

import (
	"goji.io"
	"goji.io/pat"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/controllers"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/routes/middleware"
)

func Init() *goji.Mux {
	root := goji.NewMux()

	// auth routes
	authMux := goji.SubMux()
	root.HandleC(pat.New("/auth/*"), authMux)
	authCtrl := controllers.Auth{}
	authMux.HandleFuncC(pat.Post("/signup"), authCtrl.Signup)
	authMux.HandleFuncC(pat.Post("/login"), authCtrl.Login)

	// users routes
	usersMux := goji.SubMux()
	root.HandleC(pat.New("/users/*"), usersMux)
	usersCtrl := controllers.Users{}
	usersMux.HandleFuncC(pat.Get("/:id"), usersCtrl.Get)
	usersMux.HandleFuncC(pat.Put("/:id"), usersCtrl.Update)
	usersMux.HandleFuncC(pat.Delete("/:id"), usersCtrl.Disable)

	// intakes routes
	intakesMux := goji.SubMux()
	root.HandleC(pat.New("/intakes/*"), intakesMux)
	intakesCtrl := controllers.Intakes{}
	intakesMux.HandleFuncC(pat.Post("/"), intakesCtrl.Create)
	intakesMux.HandleFuncC(pat.Get("/"), intakesCtrl.GetAll)
	intakesMux.HandleFuncC(pat.Get("/:id"), intakesCtrl.Get)
	intakesMux.HandleFuncC(pat.Put("/:id"), intakesCtrl.Update)
	intakesMux.HandleFuncC(pat.Delete("/:id"), intakesCtrl.Disable)

	// intakes routes
	adminMux := goji.SubMux()
	root.HandleC(pat.New("/admin/*"), adminMux)
	adminCtrl := controllers.Admin{}
	adminMux.HandleFuncC(pat.Get("/users"), adminCtrl.GetUsers)
	adminMux.HandleFuncC(pat.Post("/users"), adminCtrl.CreateUser)
	adminMux.HandleFuncC(pat.Get("/users/:user_id"), adminCtrl.GetUser)
	adminMux.HandleFuncC(pat.Put("/users/:user_id"), adminCtrl.UpdateUser)
	adminMux.HandleFuncC(pat.Delete("/users/:user_id"), adminCtrl.DisableUser)

	/*
	   //	For later consideration:
	   	adminMux.HandleFuncC(pat.Get("/users/:user_id/intakes"), adminCtrl.GetUserIntakes)
	   	adminMux.HandleFuncC(pat.Post("/users/:user_id/intakes"), adminCtrl.CreateUserIntake)
	   	adminMux.HandleFuncC(pat.Put("/users/:user_id/intakes/:intake_id"), adminCtrl.UpdateUserIntake)
	   	adminMux.HandleFuncC(pat.Delete("/users/:user_id/intakes/:intake_id"), adminCtrl.DisableUserIntake)
	*/

	// hook middleware
	root.UseC(middleware.HTTPLogger)
	root.UseC(middleware.JSONHeader)
	root.Use(middleware.CORSHeader)

	return root
}
