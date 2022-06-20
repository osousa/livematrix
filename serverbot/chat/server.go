package chat

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
	"maunium.net/go/mautrix/format"
	mid "maunium.net/go/mautrix/id"
)

// Chat server.
type Server struct {
	pattern        string
	messages       []*JSONMessage
	clients        map[int]*Client
	addCh          chan *Client
	delCh          chan *Client
	sendAllCh      chan *JSONMessage
	doneCh         chan bool
	errCh          chan error
	Mautrix_client *BotPlexer
}

// Create new chat server.
func NewServer(pattern string, mautrix_client *BotPlexer) *Server {
	messages := []*JSONMessage{}
	clients := make(map[int]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	sendAllCh := make(chan *JSONMessage)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		pattern,
		messages,
		clients,
		addCh,
		delCh,
		sendAllCh,
		doneCh,
		errCh,
		mautrix_client,
	}
}

func (s *Server) FindClientByRoomID(roomid mid.RoomID) (Client, error) {

	for _, client := range s.clients {
		if c := client.GetRoomId(); mid.RoomID(*c) == roomid {
			return *client, nil
		}
	}
	return Client{}, error(fmt.Errorf("No clients with such RoomID"))
}

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) Del(c *Client) {
	s.delCh <- c
}

func (s *Server) SendAll(msg *JSONMessage) {
	s.sendAllCh <- msg
}

func (s *Server) Done() {
	s.doneCh <- true
}

func (s *Server) Err(err error) {
	s.errCh <- err
}

func (s *Server) SendMatrixMessage(c *Client, msg JSONMessage) {
	var r mid.RoomID
	r = mid.RoomID(*c.session.RoomID)
	log.Println(r)
	content := format.RenderMarkdown(msg.Body, true, true)
	s.Mautrix_client.SendMessage(r, &content)
}

func (s *Server) sendPastMessages(c *Client) {
	for _, msg := range c.GetJSONMessages() {
		c.Write(msg)
	}
}

func (s *Server) sendAll(msg *JSONMessage) {
	for _, c := range s.clients {
		c.Write(msg)
	}
}

// Trying to access the original request before it upgrades the http connection
// to a websocket one. Use this to apply any middlewares, as for authentication
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	tokenCookie, err := r.Cookie("session_id")
	if err != nil {
		log.Panic("No session cookie, abort!")
	}

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		client := NewClient(ws, s, tokenCookie)
		s.Add(client)
		client.Listen()
	}
	handler := websocket.Handler(onConnected)
	handler.ServeHTTP(w, r)
}

// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen() {
	session := NewSession(nil, nil)
	http.Handle("/session", session)
	http.Handle(s.pattern, s)
	log.Println("Created handler")

	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			log.Println("Added new client")
			s.clients[c.id] = c
			log.Println("Now", len(s.clients), "clients connected.")
			s.sendPastMessages(c)
			if rid := c.GetRoomId(); *rid != "" {
				log.Println(*rid != "")
				s.Mautrix_client.JoinRoomByID(mid.RoomID(*rid))
			} else {
				roomid, err := s.Mautrix_client.CreateRoom(c)
				if err != nil {
					log.Println("Could not create room, abort!")
				} else {
					*c.session.RoomID = string(roomid)
					DB.UpdateRow(c.session)
				}
			}

		// del a client
		case c := <-s.delCh:
			log.Println("Delete client")
			delete(s.clients, c.id)

		// broadcast message for all clients
		case msg := <-s.sendAllCh:
			s.messages = append(s.messages, msg)
			//s.sendAll(msg)

		// listens to matrix events
		case matrix_evt := <-s.Mautrix_client.Ch:
			client, err := s.FindClientByRoomID(matrix_evt.RoomID)
			if err == nil {
				log.Println()
				jsonmsg := NewJSONMessage(matrix_evt.Content.Raw["body"].(string), "0")
				client.Write(jsonmsg)
			}

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}
