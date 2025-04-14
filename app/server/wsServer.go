package server

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"encoding/json"
)

var Rooms = NewRoomMap()

type Message struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Data interface{} `json:"data"`
}

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roomID := Rooms.CreateRoom()

	log.Printf("Room created: %s", roomID)

	type res struct {
		RoomID string `json:"roomID"`
	}

	json.NewEncoder(w).Encode(res{RoomID: roomID})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var broadcastChan = make(chan Message)

func broadcast(roomID string) {
	for {
		msg := <-broadcastChan

		for client := range Rooms.GetClientss(roomID) {
			if err := client.Conn.WriteJSON(msg); err != nil {
				log.Printf("Error broadcasting to client: %v", err)
				Rooms.RemoveClient(roomID, client)
			}
		}
	}
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	ri, ok := r.URL.Query()["roomID"]
	roomID := ri[0]
	if !ok {
		http.Error(w, "Room ID is required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Error upgrading connection", http.StatusInternalServerError)
		return
	}

	client := &Client{Host: false, Conn: conn}
	Rooms.AddClient(roomID, client)
	defer Rooms.RemoveClient(roomID, client)

	var msg Message
	err = conn.ReadJSON(&msg)

	// first message must be a name
	if err != nil || msg.Type != "name" {
		log.Println(err)
		return
	}

	client.Name = msg.Data.(string)

	go broadcast(roomID)

	for {
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}

		msg.Name = client.Name
		msg.Type = "message"

		broadcastChan <- msg
	}
}
