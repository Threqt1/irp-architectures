package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Threqt1/architecture-go/messages-microservice/config"
	"github.com/Threqt1/architecture-go/messages-microservice/library/snowflake"
	"github.com/Threqt1/architecture-go/messages-microservice/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIv1 struct {
	router    *mux.Router
	snowflake snowflake.SnowflakeProvider
}

func (api *APIv1) Start() {
	api.router = mux.NewRouter()

	api.router.HandleFunc(config.MESSAGE_MS_MESSAGES_ROUTE, api.CreateMessage).Methods(http.MethodPut)

	handler := cors.Default().Handler(api.router)
	http.ListenAndServe(config.MESSAGES_MS_PORT, handler)
}

func (api *APIv1) CreateMessage(w http.ResponseWriter, r *http.Request) {
	message := models.APIMessage{}

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		log.Println("failed to decode JSON")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	message.ID = api.snowflake.Generate()

	json.NewEncoder(w).Encode(message)
}

func CreateAPIv1(snowflake snowflake.SnowflakeProvider) APIv1 {
	return APIv1{snowflake: snowflake}
}
