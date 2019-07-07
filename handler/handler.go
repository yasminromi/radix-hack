package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"radix-hack/model"
	"radix-hack/service"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan model.Message)

type Handler struct {
	Upgrader websocket.Upgrader
	Service  *service.ElasticService
}

func (h *Handler) SendViaPost(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		var payload model.Message

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Printf("ERROR: %s", err)
			http.Error(w, "Bad request", http.StatusTeapot)
			return
		}

		defer r.Body.Close()

		log.Printf("payload sending via post: %v", payload)
		broadcast <- payload

		ctx := r.Context()

		error := h.Service.SaveToElastic(ctx, payload)
		if error != nil {
			log.Printf("error saving to ES: %v", error)
			return
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) HandleConnections(w http.ResponseWriter, r *http.Request) {

	ws, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true
	ctx := r.Context()

	for {
		var msg model.Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error reading json: %v", err)
			delete(clients, ws)
			return
		}

		log.Printf("payload reading: %v", msg)
		broadcast <- msg

		error := h.Service.SaveToElastic(ctx, msg)
		if error != nil {
			log.Printf("error saving to ES: %v", error)
			delete(clients, ws)
			return
		}

	}

}

func (h *Handler) HandleMessages() {
	for {

		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			log.Printf("payload writting: %v", msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

//func (i18n *I18n) GetTranslation(language, group, key string) (string, error) {

//	applicationId := "35"
//	group = language + "." + group
//	cacheKey := applicationId + group + key
//	value, _ := i18n.Cache.Get(cacheKey)

//	if value == "" {
//		valueFromMetadata, err := client.GetMetadataValue(applicationId, group, key)
//
//		if err != nil {
//			return "", err
//		}
//
//		i18n.Cache.Set(cacheKey, value, 30 * time.Minute)
//		value = valueFromMetadata
//	}
//	return value, nil
//}
