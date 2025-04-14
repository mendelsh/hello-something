package server

import (
	"log"
	"math/rand"
	"sync"
	"github.com/gorilla/websocket"
)

type Client struct {
	Host   bool
	Conn   *websocket.Conn
	Name   string
}

type Room struct {
	Clients map[*Client]bool
}

func (r *Room) removeClient(client *Client) {
	delete(r.Clients, client)
	client.Conn.Close()
}

type RoomMap struct {
	Mutex sync.RWMutex
	Rooms map[string]*Room
}

func NewRoomMap() *RoomMap {
	return &RoomMap{
		Rooms: make(map[string]*Room),
	}
}

func (rm *RoomMap) GetClientss(roomID string) map[*Client]bool {
	rm.Mutex.RLock()
	defer rm.Mutex.RUnlock()
	
	return rm.Rooms[roomID].Clients
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (rm *RoomMap) CreateRoom() string {
	rm.Mutex.Lock()
	defer rm.Mutex.Unlock()

	roomID := randSeq(10)
	for _, ok := rm.Rooms[roomID]; ok; {
		roomID = randSeq(10)
	}

	rm.Rooms[roomID] = &Room{
		Clients: make(map[*Client]bool),
	}

	return roomID
}

func (rm *RoomMap) AddClient(roomID string, client *Client) {
	rm.Mutex.Lock()
	defer rm.Mutex.Unlock()

	room, ok := rm.Rooms[roomID]
	if !ok {
		log.Println("Room not found")
		return
	}
	room.Clients[client] = true
}

func (rm *RoomMap) RemoveClient(roomID string, client *Client) {
	rm.Mutex.Lock()
	defer rm.Mutex.Unlock()

	room, ok := rm.Rooms[roomID]
	if !ok {
		log.Println("Room not found")
		return
	}
	if _, ok := room.Clients[client]; !ok {
		log.Println("Client not found in room")
		return
	}

	room.removeClient(client)
}

func (rm *RoomMap) DeleteRoom(roomID string) {
	rm.Mutex.Lock()
	defer rm.Mutex.Unlock()

	delete(rm.Rooms, roomID)
}
