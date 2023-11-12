package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/Threqt1/architecture-go/match-microservice/config"
	"github.com/Threqt1/architecture-go/match-microservice/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIv1 struct {
	router  *mux.Router
	queue   chan string
	matches map[string]string
	mutex   sync.RWMutex
}

func (api *APIv1) Start() {
	api.router = mux.NewRouter()

	api.router.HandleFunc(config.USERS_MS_USERS_ROUTE, api.MatchUser).Methods(http.MethodPut)
	api.router.HandleFunc(config.USERS_MS_USERS_ROUTE, api.GetMatch).Methods(http.MethodGet)

	handler := cors.Default().Handler(api.router)
	http.ListenAndServe(config.USERS_MS_PORT, handler)
}

func (api *APIv1) MatchUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("bad id query param")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	select {
	case match := <-api.queue:
		api.mutex.Lock()
		api.matches[match] = id
		api.mutex.Unlock()
		json.NewEncoder(w).Encode(models.APIMatch{MatchedUserID: match, WaitFor: 0})
	default:
		api.queue <- id

		json.NewEncoder(w).Encode(models.APIMatch{WaitFor: int(config.DEFAULT_RETRY_TIME)})
	}
}

func (api *APIv1) GetMatch(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("bad id query param")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	api.mutex.RLock()
	match, exists := api.matches[id]
	api.mutex.RUnlock()
	if !exists {
		json.NewEncoder(w).Encode(models.APIMatch{WaitFor: int(config.DEFAULT_RETRY_TIME)})
		return
	}

	json.NewEncoder(w).Encode(models.APIMatch{MatchedUserID: match, WaitFor: 0})
}

func CreateAPIv1() APIv1 {
	channel := make(chan string, 1000)
	channel <- "0"
	return APIv1{queue: channel, matches: make(map[string]string), mutex: sync.RWMutex{}}
}
