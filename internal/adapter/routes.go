package adapter

import (
	"lucasdorazio/golang-example/config"
	"lucasdorazio/golang-example/internal"
	"lucasdorazio/golang-example/internal/adapter/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// MIDDLEWARES
	mux.Use(middleware.Recoverer)
	mux.Use(internal.WriteToConsole)

	// ROUTES
	mux.Get("/", controller.ControllersRepo.Home)
	mux.Get("/about", controller.ControllersRepo.About)
	mux.Get("/addition", controller.ControllersRepo.AdditionController)
	mux.Get("/divide", controller.ControllersRepo.DivideController)

	return mux
}
