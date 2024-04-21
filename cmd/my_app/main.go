package main

import (
	"fmt"
	"log"
	"lucasdorazio/golang-example/config"
	render "lucasdorazio/golang-example/internal"
	"lucasdorazio/golang-example/internal/adapter"
	"lucasdorazio/golang-example/internal/adapter/controller"
	"net/http"
)

const port = ":8084"

func main() {

	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := controller.NewControllersRepo(&app)
	controller.NewControllers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    port,
		Handler: adapter.Routes(&app),
	}

	fmt.Printf("Starting application on port %s", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	// http.HandleFunc("/", controller.ControllersRepo.Home)
	// http.HandleFunc("/about", controller.ControllersRepo.About)
	// http.HandleFunc("/addition", controller.ControllersRepo.AdditionController)
	// http.HandleFunc("/divide", controller.ControllersRepo.DivideController)

	/*
		Print: 	print estandar, convierte a cadena e imprime en la salida estandar todos los argumentos
		Fprint:	imprime en un io.writer
		Sprint:	devuelve la cadena (util si se quiere formatear con Sprintf)

		Cada uno tiene su version de printf para formatear
	*/

	// http.ListenAndServe(port, nil) // Probar con curl http://localhost:8084
}
