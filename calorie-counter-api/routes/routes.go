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

	v1 := goji.SubMux()
	v1.UseC(middleware.BearerAuth)
	root.HandleC(pat.New("/v1/*"), v1)

	// users routes
	usersCtrl := controllers.Users{}
	v1.HandleFuncC(pat.Get("/users"), usersCtrl.Get)
	v1.HandleFuncC(pat.Put("/users/:id"), usersCtrl.Update)
	v1.HandleFuncC(pat.Delete("/users/:id"), usersCtrl.Disable)

	// intakes routes
	intakesCtrl := controllers.Intakes{}
	v1.HandleFuncC(pat.Post("/intakes"), intakesCtrl.Create)
	v1.HandleFuncC(pat.Get("/intakes"), intakesCtrl.GetAll)
	v1.HandleFuncC(pat.Get("/intakes/:id"), intakesCtrl.Get)
	v1.HandleFuncC(pat.Put("/intakes/:id"), intakesCtrl.Update)
	v1.HandleFuncC(pat.Delete("/intakes/:id"), intakesCtrl.Disable)

	admin := goji.SubMux()
	admin.UseC(middleware.BearerAuth)
	admin.UseC(middleware.RequireAdmin)
	root.HandleC(pat.New("/admin/*"), admin)

	// admin routes
	adminCtrl := controllers.Admin{}
	admin.HandleFuncC(pat.Get("/users"), adminCtrl.GetUsers)
	admin.HandleFuncC(pat.Post("/users"), adminCtrl.CreateUser)
	admin.HandleFuncC(pat.Get("/users/:user_id"), adminCtrl.GetUser)
	admin.HandleFuncC(pat.Put("/users/:user_id"), adminCtrl.UpdateUser)
	admin.HandleFuncC(pat.Delete("/users/:user_id"), adminCtrl.DisableUser)

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
