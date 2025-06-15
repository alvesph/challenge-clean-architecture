package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alvesph/challenge-clean-architecture/configs"
	"github.com/alvesph/challenge-clean-architecture/graph"
	"github.com/alvesph/challenge-clean-architecture/graph/generated"
	"github.com/alvesph/challenge-clean-architecture/internal/entity"
	"github.com/alvesph/challenge-clean-architecture/internal/repository"
	"github.com/alvesph/challenge-clean-architecture/internal/service"
	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(&entity.Order{})

	orderRepo := repository.NewOrder(db)
	orderService := service.NewOrderService(orderRepo)

	resolver := &graph.Resolver{
		OrderService: orderService,
	}

	srv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.GraphqlPort)
	log.Fatal(http.ListenAndServe(":"+cfg.GraphqlPort, nil))
}
