package main

import (
	"fmt"
	"log"
	"net/http"
	"recipeServer/db"
	"recipeServer/handlers"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	initEnv()

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.Handle("/api/recipes/", &handlers.RecipeHandler{Prefix: "/api/recipes/"})

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodOptions, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
		AllowedOrigins:   []string{db.GetEnvVar("CORS")},
		Debug:            true,
	})

	db.Connect()
	// http.ListenAndServe(":5000", http.FileServer(http.Dir("public")))PORT
	PORT := db.GetEnvVar("PORT")
	fmt.Printf("listening on port %s\n", PORT)
	http.ListenAndServe(":"+PORT, c.Handler(mux))
}
