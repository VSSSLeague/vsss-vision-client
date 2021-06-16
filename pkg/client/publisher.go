package client

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/VSSSLeague/vsss-vision-client/pkg/vision"
	"github.com/gorilla/websocket"
)

const publishDt = 50 * time.Millisecond
const visionSource = "vision"

type PublishType int

type Publisher struct {
	upgrader          websocket.Upgrader
	DetectionProvider func() *vision.Frame
	GeometryProvider  func() *vision.Field
}

func NewPublisher() (p Publisher) {
	p.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(*http.Request) bool { return true },
	}

	// Return new p, Providers will be set in the main module

	return p
}

type PublisherClient struct {
	conn                *websocket.Conn
	activeTrackedSource string
	mutex               sync.Mutex
}

func (p *Publisher) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := p.upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Could not upgrade connection: ", err)
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("Could not close websocket connection: ", err)
		}
		log.Println("Client disconnected")
	}()

	log.Println("Client connected")

	client := &PublisherClient{}
	client.conn = conn
	client.activeTrackedSource = r.URL.Query().Get("sourceId")
	go client.handleClientRequests()

	for {
		pack := new(Package)

		client.mutex.Lock()

		// Set detection frame to Package
		detectionFrame := p.DetectionProvider()
		pack.AddDetectionFrame(detectionFrame)

		client.mutex.Unlock()

		// Set geometry detection to Package
		geometry := p.GeometryProvider()
		pack.AddGeometryShapes(geometry)

		// Set pack to Publisher
		p.addVisualization(pack)

		// Convert struct into json
		payload, err := json.Marshal(*pack)

		// Check errors in conversion
		if err != nil {
			return
		}

		// Write message to websocket (broadcast)
		if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
			log.Println(err)
			return
		}

		time.Sleep(publishDt)
	}
}

func (p *PublisherClient) handleClientRequests() {
	for {
		messageType, data, err := p.conn.ReadMessage()
		if err != nil {
			log.Println("Could not read mesage: ", err)
			return
		}

		if messageType == websocket.TextMessage {
			var request Request
			if err := json.Unmarshal(data, &request); err != nil {
				log.Println("Could not deserialize message: ", string(data))
			} else {
				p.mutex.Lock()
				p.activeTrackedSource = request.ActiveSourceId
				p.mutex.Unlock()
			}
		} else {
			log.Println("Got non-text message")
		}
	}
}

func (p *Publisher) addVisualization(pack *Package) {
	pack.SortShapes()
}
