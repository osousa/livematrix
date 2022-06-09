package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"osousa.me/chat"

	"github.com/joho/godotenv"
)

var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)
var Warning = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)
var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
var Debug = log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Does it exist?")
	}

	db_pass := os.Getenv("DATABASE_PASSWORD")
	db_name := os.Getenv("DATABASE_NAME")
	db_user := os.Getenv("DATABASE_USER")

	_, _ = chat.ConnectSQL(db_user, db_pass, db_name)

	roomch := make(chan *chat.Client, 8)
	App := chat.NewApp(roomch)
	go App.Connect("this", "that")

	// websocket server
	server := chat.NewServer("/entry", App)
	go server.Listen()

	// static files
	http.Handle("/", http.FileServer(http.Dir("webroot")))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
