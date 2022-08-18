package routes

import (
	"waysbuck/handlers"
	"waysbuck/pkg/mysql"
	"waysbuck/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	profileRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(profileRepository)

	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
}
