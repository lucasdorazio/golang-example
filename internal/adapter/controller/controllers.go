package controller

import (
	"errors"
	"fmt"
	"log"
	"lucasdorazio/golang-example/config"
	render "lucasdorazio/golang-example/internal"
	"net/http"
)

type ControllersRepository struct {
	App *config.AppConfig
}

// Repository used by the controllers
var ControllersRepo *ControllersRepository

// Creates a new repository
func NewControllersRepo(a *config.AppConfig) *ControllersRepository {
	return &ControllersRepository{a}
}

// Sets the repository for the controllers
func NewControllers(r *ControllersRepository) {
	ControllersRepo = r
}

// --------------------- HANDLERS ----------------

func (m *ControllersRepository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *ControllersRepository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

func (m *ControllersRepository) HomeController(w http.ResponseWriter, r *http.Request) {
	// Respuesta
	n, err := fmt.Fprintf(w, "Hello World! :)")

	if err != nil {
		log.Println(err)
	}

	// Consola
	numberOfBytesText := fmt.Sprintf("Number of bytes written: %d", n)
	fmt.Println(numberOfBytesText)

}

func (m *ControllersRepository) AboutController(w http.ResponseWriter, r *http.Request) {
	// Respuesta
	n, err := fmt.Fprintf(w, "We are Think Software")

	if err != nil {
		log.Println(err)
	}

	// Consola
	numberOfBytesText := fmt.Sprintf("Number of bytes written: %d", n)
	fmt.Println(numberOfBytesText)

}

func (m *ControllersRepository) AdditionController(w http.ResponseWriter, r *http.Request) {
	num1 := 10
	num2 := 27
	sum := addValues(num1, num2)

	// Respuesta
	fmt.Fprintf(w, "%d + %d = %d", num1, num2, sum)
}

func (m *ControllersRepository) DivideController(w http.ResponseWriter, r *http.Request) {
	var num1, num2 float32
	num1 = 12.1
	num2 = 0
	value, err := divide(num1, num2)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprintf(w, "%f / %f = %f", num1, num2, value)
}

// --------------------- Funciones privadas ----------------

func addValues(x, y int) int {
	return x + y
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
