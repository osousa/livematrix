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
	matrix_user := os.Getenv("MATRIX_USERNAME")
	matrix_pass := os.Getenv("MATRIX_PASSWORD")

	db, err := chat.ConnectSQL(db_user, db_pass, db_name)

	// If one wishes, they can move this to another file, but not database.go
	query := `CREATE TABLE if not exists Session(
			  id int(11) NOT NULL AUTO_INCREMENT,
			  session varchar(100) NOT NULL,
			  expirity varchar(100) DEFAULT NULL,
			  alias varchar(100) DEFAULT NULL,
			  email varchar(100) DEFAULT NULL,
			  ip varchar(100) DEFAULT NULL,
			  RoomID varchar(256) CHARACTER SET utf8 DEFAULT NULL,
			  PRIMARY KEY (id)
			) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4;`

	db.RawQuery(query)

	App := chat.NewApp()
	go App.Connect(matrix_user, matrix_pass)

	// websocket server
	server := chat.NewServer("/entry", App)
	go server.Listen()

	// static files
	http.Handle("/", http.FileServer(http.Dir("webroot")))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
