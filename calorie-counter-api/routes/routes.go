package routes

import (
	"goji.io"
	"goji.io/pat"
)

func Init() *goji.Mux {
	root := goji.NewMux()

	// auth routes
	authMux := goji.SubMux()
	root.HandleC(pat.New("/auth/*"), authMux)

	// users routes
	usersMux := goji.SubMux()
	root.HandleC(pat.New("/users/*"), usersMux)

	// intakes routes
	intakesMux := goji.SubMux()
	root.HandleC(pat.New("/intakes/*"), intakesMux)

	return root
}
