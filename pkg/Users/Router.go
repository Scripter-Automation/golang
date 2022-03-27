package Users

import "github.com/go-chi/chi/v5"

func UserRouts() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/getUser", getHandler)
	router.Post("/updateUser", postHandler)

	return router
}
