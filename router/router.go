package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func init() {
	router := chi.NewRouter()

	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		panic(err)
	}
}

// import (
// 	"goapi/internal/middleware"

// 	"github.com/go-chi/chi"
// 	chimiddle "github.com/go-chi/chi/middleware"
// )

// func Handler(r *chi.Mux) {
// 	r.Use(chimiddle.StripSlashes)
// 	r.Route("/account", func(router chi.Router) {
// 		router.Use(middleware.Authorization)
// 		router.Get("/coins", GetCoinBalance)
// 	})
// }

// func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
// 	var params = api.CoinBalanceParams{}
// 	var decoder *schema.Decoder = schema.NewDecoder()
// 	var err error

// 	err = decoder.Decode(&params, r.URL.Query())

// 	if err != nil {
// 		log.Error(err)
// 		api.InternalErrorHandler(w)
// 		return
// 	}

// 	var database *tools.DatabaseInterface
// 	database, err = tools.NewDatabase()
// 	if err != nil {
// 		api.InternalErrorHandler(w)
// 		return
// 	}

// 	var tokenDetails *tools.CoinDetails
// 	tokenDetails = (*database).GetUserCoins(params.Username)
// 	if tokenDetails == nil {
// 		log.Error(err)
// 		api.InternalErrorHandler(w)
// 		return
// 	}

// 	var response = api.CoinBalanceResponce{
// 		Balance: (*tokenDetails).Coins,
// 		Code:    http.StatusOK,
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	err = json.NewEncoder(w).Encode(response)
// 	if err != nil {
// 		log.Error(err)
// 		api.InternalErrorHandler(w)
// 		return
// 	}
// }

// package middleware

// import (
// 	"errors"
// 	"goapi/api"
// 	"goapi/internal/tools"
// 	"net/http"

// 	log "github.com/sirupsen/logrus"
// )

// var UnAuthorizedError = errors.New("Invalid username or token.")

// func Authorization(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var username string = r.URL.Query().Get("username")
// 		var token = r.Header.Get("Authorization")

// 		var err error

// 		if username == "" || token == "" {
// 			log.Error(UnAuthorizedError)
// 			api.RequestErrorHandler(w, UnAuthorizedError)
// 			return
// 		}

// 		var database *tools.DatabaseInterface
// 		database, err = tools.NewDatabase()
// 		if err != nil {
// 			api.InternalErrorHandler(w)
// 			return
// 		}

// 		var loginDetails *tools.LoginDetails
// 		loginDetails = (*database).GetUserLoginDetails(username)

// 		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
// 			log.Error(UnAuthorizedError)
// 			api.RequestErrorHandler(w, UnAuthorizedError)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

// package main

// import (
// 	"fmt"
// 	"goapi/internal/handlers"
// 	"net/http"

// 	"github.com/go-chi/chi"
// 	log "github.com/sirupsen/logrus"
// )

// func main() {

// 	log.SetReportCaller(true)
// 	var r *chi.Mux = chi.NewRouter()
// 	handlers.Handler(r)

// 	fmt.Println("Starting GO API service...")

// 	fmt.Println("!!! goapi !!!")

// 	err := http.ListenAndServe("localhost:8000", r)
// 	if err != nil {
// 		log.Error(err)
// 	}
// }
