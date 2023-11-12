package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Threqt1/architecture-go/api-gateway/adapter"
	"github.com/Threqt1/architecture-go/api-gateway/config"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIv1 struct {
	router                 *mux.Router
	userServiceAdapter     adapter.UserServiceAdapter
	matchServiceAdapter    adapter.MatchServiceAdapter
	messagesServiceAdapter adapter.MessagesServiceAdapter
}

func (api *APIv1) Start() {
	api.router = mux.NewRouter()

	api.router.HandleFunc("/test", api.RunTest).Methods(http.MethodPost)

	handler := cors.Default().Handler(api.router)
	http.ListenAndServe(config.ROOT_PORT, handler)
}

func (api *APIv1) RunTest(w http.ResponseWriter, r *http.Request) {
	user, err := api.userServiceAdapter.CreateUser()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	match, err := api.matchServiceAdapter.MatchUser(user.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	for match.WaitFor > 0 {
		time.Sleep(time.Duration(match.WaitFor))
		match, err = api.matchServiceAdapter.GetMatch(user.ID)
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	message, err := api.messagesServiceAdapter.SendMessage(user.ID, match.MatchedUserID, "Hello Other User")
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(message)
}

func CreateAPIv1() APIv1 {
	return APIv1{userServiceAdapter: adapter.CreateUserServiceAdapter(), matchServiceAdapter: adapter.CreateMatchServiceAdapter(), messagesServiceAdapter: adapter.CreateMessagesServiceAdapter()}
}
