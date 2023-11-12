package api

import (
	"encoding/json"
	"net/http"

	"github.com/Threqt1/architecture-go/user-microservice/config"
	"github.com/Threqt1/architecture-go/user-microservice/library/snowflake"
	"github.com/Threqt1/architecture-go/user-microservice/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIv1 struct {
	router    *mux.Router
	snowflake snowflake.SnowflakeProvider
}

func (api *APIv1) Start() {
	api.router = mux.NewRouter()

	api.router.HandleFunc(config.USERS_MS_USERS_ROUTE, api.CreateUser).Methods(http.MethodPut)

	handler := cors.Default().Handler(api.router)
	http.ListenAndServe(config.USERS_MS_PORT, handler)
}

func (api *APIv1) CreateUser(w http.ResponseWriter, r *http.Request) {
	id := api.snowflake.Generate()
	user := models.User{ID: id}

	json.NewEncoder(w).Encode(user)
}

func CreateAPIv1(snowflake snowflake.SnowflakeProvider) APIv1 {
	return APIv1{snowflake: snowflake}
}
