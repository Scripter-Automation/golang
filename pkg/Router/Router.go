package Router

import (
	"fmt"
	"log"
	database "main/pkg/Database"
	"main/pkg/Users"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func StartServer() *chi.Mux {
	database.SetupDBConnection()
	router := chi.NewRouter()
	router.Mount("/api/users", Users.UserRouts())
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	return router
}
