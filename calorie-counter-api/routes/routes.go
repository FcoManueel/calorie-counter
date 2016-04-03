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
	intakesMux.HandleFuncC(pat.Get("/"), intakesCtrl.GetAll)

	// intakes routes
	adminMux := goji.SubMux()
	root.HandleC(pat.New("/admin/*"), adminMux)
	adminCtrl := controllers.Admin{}
	adminMux.HandleFuncC(pat.Get("/"), adminCtrl.GetUsers)

	// hook middleware
	root.UseC(middleware.HTTPLogger)
	root.UseC(middleware.JSONHeader)

	return root
}
