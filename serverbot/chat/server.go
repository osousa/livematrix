package chat

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Chat server.
type Server struct {
	pattern   string
	messages  []*JSONMessage
	clients   map[int]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan *JSONMessage
	doneCh    chan bool
	errCh     chan error
}

// Create new chat server.
func NewServer(pattern string) *Server {
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
	}
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

		// del a client
		case c := <-s.delCh:
			log.Println("Delete client")
			delete(s.clients, c.id)

		// broadcast message for all clients
		case msg := <-s.sendAllCh:
			log.Println("Send all:", msg)
			s.messages = append(s.messages, msg)
			s.sendAll(msg)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}
