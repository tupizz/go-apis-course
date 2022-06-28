package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tupizz/go-api-rest/controllers"
	"github.com/tupizz/go-api-rest/middleware"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home)

	router.Use(middleware.ContentTypeMiddleware)

	router.
		HandleFunc("/api/personas", controllers.GetAllPersonas).
		Methods(http.MethodGet)

	router.
		HandleFunc("/api/personas/{personaId:[0-9]+}", controllers.GetOnePersona).
		Methods(http.MethodGet)

	router.
		HandleFunc("/api/personas", controllers.CreatePersona).
		Methods(http.MethodPost)

	router.
		HandleFunc("/api/personas/{personaId:[0-9]+}", controllers.DeletePersona).
		Methods(http.MethodDelete)

	router.
		HandleFunc("/api/personas/{personaId:[0-9]+}", controllers.EditPersona).
		Methods(http.MethodPut)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(
		http.ListenAndServe(
			":8585",
			handlers.CORS(allowedOrigins)(router),
		),
	)
}
