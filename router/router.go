package router

import (
	"encoding/json"
	"mg-weather-service/orchestrator"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RunServer() {
	router := chi.NewRouter()
	Handler(router)
	println("Started serving requests")
	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		panic(err)
	}
}

func Handler(router *chi.Mux) {
	router.Route("/api/v1", func(router chi.Router) {
		router.Get("/city/{city}", GetWeatherForCity)
	})
}

func GetWeatherForCity(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "city")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orchestrator.GetWeather(param))
}
