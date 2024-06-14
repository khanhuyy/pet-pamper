package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"net/http"
	"os"
	"pet-pamper/gateway/graphql/resolver"

	"pet-pamper/gateway/graphql/api"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//svr := handler.NewDefaultServer(api.NewExecutableSchema(api.Config{Resolvers: resolver.NewResolver()}))
	cfg := api.Config{Resolvers: resolver.NewResolver()}

	//cfg.Directives.Binding = directives.BindingValidator
	//cfg.Directives.Secured = directives.SecuredDirective
	//cfg.Directives.StaffSecured = directives.StaffSecuredDirective
	//cfg.Directives.SellerSecured = directives.SellerSecuredDirective

	es := api.NewExecutableSchema(cfg)
	srv := handler.NewDefaultServer(es)
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		var defErr = graphql.DefaultErrorPresenter(ctx, err)
		//var st = StatusFromError(err)
		//if st != nil {
		//	st, err := st.WithDetails(&errdetails.DebugInfo{
		//		Detail: err.Error(),
		//	})
		//	if err != nil {
		//		return defErr
		//	}
		//	defErr.Message = st.Message()
		//	defErr.Extensions = map[string]interface{}{"details": st.Details()}
		//}
		return defErr
	})
	//srv.Use(otelgqlgen.Middleware())

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
