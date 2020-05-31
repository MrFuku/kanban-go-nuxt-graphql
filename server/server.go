package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MrFuku/kanban-go-nuxt-graphql/server/graph"
	"github.com/MrFuku/kanban-go-nuxt-graphql/server/graph/generated"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi"
)

const dataSource = "localuser:localpass@tcp(db:3306)/localdb?charset=utf8&parseTime=True&loc=Local"
const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(err)
	}
	defer func() {
		if db != nil {
			if err := db.Close(); err != nil {
				panic(err)
			}
		}
	}()
	defer db.Close()
	db.LogMode(true)

	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(cors.Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))
	var mb int64 = 1 << 20
	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     128 * mb,
		MaxUploadSize: 100 * mb,
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
