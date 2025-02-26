package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type SensorData struct {
	Type      string  `json:"type"`
	Value     float64 `json:"value"`
	Unit      string  `json:"unit"`
	Timestamp string  `json:"timestamp"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(clients, ws)
			break
		}

		var receivedData SensorData
		if err := json.Unmarshal(message, &receivedData); err != nil {
			log.Printf("Invalid JSON format: %v", err)
			continue
		}

		responseMessage := fmt.Sprintf("Received sensor data: %s with a value of %.2f at %s", receivedData.Type, receivedData.Value, receivedData.Timestamp)
		if err := ws.WriteMessage(websocket.TextMessage, []byte(responseMessage)); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	go simulateSensors()

	http.ListenAndServe(":8083", nil)
}

func simulateSensors() {
	timer := time.NewTicker(5 * time.Second)

	for range timer.C {
		now := time.Now()
		isoTime := now.Format(time.RFC3339)

		humidityData := SensorData{
			Type:      "humidity",
			Value:     rand.Float64() * 100,
			Unit:      "%",
			Timestamp: isoTime,
		}

		lightData := SensorData{
			Type:      "light",
			Value:     500 + rand.Float64()*1000,
			Unit:      "lux",
			Timestamp: isoTime,
		}

		data := []SensorData{humidityData, lightData}
		sendData(data)
	}
}

func sendData(data []SensorData) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshaling sensor data: %v", err)
		return
	}
	for client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}
