package gateway

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nightwolf93/transmitter-server/net"
	log "github.com/sirupsen/logrus"
)

// Type of the client for this gateway
var typeClient = 1

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// InitWSGatewayServer create the gateway with a websocket entry
func InitWSGatewayServer(host string, port int) {
	// Register the Websocket endpoint
	http.HandleFunc("/websocket_gateway", func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Incoming websocket connection ..")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Errorf("Can't handle the connection of the client: %s", err)
			return
		}

		client := &net.Client{Alive: true, Type: typeClient, WSConn: conn, WSReceiveMessageChan: make(chan []byte)}
		handleNewClient(client)
	})

	// Listen on the websocket gateway port
	log.Infof("Websocket Gateway listen on %s:%d", host, port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		log.Fatalf("Can't initialze the websocket gateway: %s", err)
		return
	}
}

func handleNewClient(client *net.Client) {
	AddClient <- client

	go func() {
		for client.Alive {
			select {
			case bytes := <-client.WSReceiveMessageChan:
				ReceiveClientData <- struct {
					client *net.Client
					bytes  []byte
				}{client: client, bytes: bytes}
				break
			}
		}
	}()

	// Receive messages
	go func() {
		for {
			_, message, err := client.WSConn.ReadMessage()
			if err != nil {
				log.Errorf("Connection with the client has been closed: %s", err)
				RemoveClient <- client
				return
			}
			client.WSReceiveMessageChan <- message
		}
	}()
}
