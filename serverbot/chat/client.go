package chat

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

const channelBufSize = 100

var maxId int = 0

// Chat client.
type Client struct {
	id      int
	ws      *websocket.Conn
	server  *Server
	ch      chan *JSONMessage
	doneCh  chan bool
	session *Session
}

// Create new chat client.
func NewClient(ws *websocket.Conn, server *Server, token *http.Cookie) *Client {
	if ws == nil {
		panic("ws cannot be nil")
	}
	if server == nil {
		panic("server cannot be nil")
	}

	maxId++
	ch := make(chan *JSONMessage, channelBufSize)
	doneCh := make(chan bool)
	session := NewSession(nil, nil)
	err := DB.GetByPk(session, token.Value, "session")
	if err != nil {
		log.Println(err)
	}

	return &Client{maxId, ws, server, ch, doneCh, session}
}

func (c *Client) Conn() *websocket.Conn {
	return c.ws
}

func (c *Client) Write(msg *JSONMessage) {
	select {
	case c.ch <- msg:
	default:
		c.server.Del(c)
		err := fmt.Errorf("client %d is disconnected.", c.id)
		c.server.Err(err)
	}
}

func (c *Client) Done() {
	c.doneCh <- true
}

// Listen Write and Read request via chanel
func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

func (c *Client) GetJSONMessages() []*JSONMessage {
	var jsonlist []*JSONMessage
	messages := c.session.Messages
	if messages != nil {
		for _, msg := range *c.session.Messages {
			jsonlist = append(jsonlist, &JSONMessage{Author: *msg.Author, Body: *msg.Body})
		}
	}
	return jsonlist
}

// Listen write request via chanel
func (c *Client) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {

		// send message to the client
		case msg := <-c.ch:
			log.Println("Send:", msg)
			websocket.JSON.Send(c.ws, msg)

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

// Listen read request via chanel
func (c *Client) listenRead() {
	log.Println("Listening read from client")
	for {
		select {

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenWrite method
			return

		// read data from websocket connection
		default:
			var msg JSONMessage
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF {
				c.doneCh <- true
			} else if err != nil {
				c.server.Err(err)
			} else {
				c.server.SendMatrixMessage(c, msg)
				c.server.SendAll(&msg)
			}
		}
	}
}
