package chat

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Session struct {
	Id        int     `db:"id"`
	SessionId *string `db:"session"`
	Expirity  *string `db:"expirity"`
	Alias     *string `db:"alias"`
	Email     *string `db:"email"`
	IpAddr    *string `db:"ip"`
	RoomID    *string `db:"RoomID"`
	Messages  *[]Message
	mutex     *sync.Mutex
}

func NewSession(msgs *[]Message, sess_id []byte, args ...*string) *Session {
	if len(args) == 0 {
		return &Session{
			0,
			new(string),
			new(string),
			new(string),
			new(string),
			new(string),
			new(string),
			msgs,
			new(sync.Mutex),
		}
	} else if len(args) == 5 {
		return &Session{
			0,
			args[0], //SessionId
			args[1], //Expirity
			args[2], //Alias
			args[3], //Email
			args[4], //IpAddr
			args[5], //IpAddr
			msgs,
			new(sync.Mutex),
		}
	}
	return nil
}

func (e *Session) GetRoomId() *string {
	return e.RoomID
}

func (e *Session) GetById(id int) error {
	tmp := NewSession(nil, nil)
	err := DB.GetById(tmp, id)
	if err != nil {
		return err
	}
	*e = *tmp
	return nil
}

func Hash254(args ...string) string {
	hash := sha256.New()
	for _, arg := range args {
		hash.Write([]byte(arg))
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// Mutexes might not be enough , so we sleep for a bit in order to make sure we
// get a different token every time, avoiding race conditions that could result
// in a duplicate token
func (s *Session) createCookie(name string) *http.Cookie {
	format := "2006-01-02 15:04:05 -0700"
	*s.Expirity = time.Now().Add(365 * 24 * time.Hour).Format(format)
	//Not that secure...
	*s.SessionId = Hash254(strconv.Itoa(rand.Intn(128000)) + *s.Expirity)
	time, _ := time.Parse(format, *s.Expirity)
	return &http.Cookie{
		Value:   *s.SessionId,
		Name:    "session_id",
		Expires: time,
	}
}

func (s *Session) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie("session_id")
	if err != nil {
		http.SetCookie(w, s.createCookie("session_id"))
	} else {
		log.Println(tokenCookie.Value)
	}
}
