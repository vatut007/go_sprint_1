package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

var cars = map[string]string{
	"id1": "Renault Logan",
	"id2": "Renault Duster",
	"id3": "BMW X6",
	"id4": "BMW M5",
	"id5": "VW Passat",
	"id6": "VW Jetta",
	"id7": "Audi A4",
	"id8": "Audi Q7",
}

// carsListFunc — вспомогательная функция для вывода всех машин.
func carsListFunc() []string {
	var list []string
	for _, c := range cars {
		list = append(list, c)
	}
	return list
}

// carFunc — вспомогательная функция для вывода определённой машины.
func carFunc(id string) string {
	if c, ok := cars[id]; ok {
		return c
	}
	return "unknown identifier " + id
}

func carBrandHandle(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(chi.URLParam(r, "brand")))
}

func carBrandModelHandle(rw http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	model := chi.URLParam(r, "model")
	final := fmt.Sprintf("%s %s", brand, model)
	rw.Write([]byte(final))
}

func carsHandle(rw http.ResponseWriter, r *http.Request) {
	carsList := carsListFunc()
	io.WriteString(rw, strings.Join(carsList, ", "))
}

func carHandle(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(carFunc(chi.URLParam(r, "id"))))
}

func main() {
	r := chi.NewRouter()
	r.Get("/", carsHandle)
	r.Get("/car/{id}", carHandle)
	r.Route("/cars", func(r chi.Router) {
		r.Get("/", carsHandle)
		r.Route("/{brand}", func(r chi.Router) {
			r.Get("/", carBrandHandle)
			r.Get("/{model}", carBrandModelHandle)
		})
	})
	// r.Get("/cars/{brand}", carBrandHandle)
	http.ListenAndServe(":8080", r)
}
