package chat

import (
	"errors"
	"fmt"
	"time"

	retry "github.com/sethvargo/go-retry"
	log "github.com/sirupsen/logrus"
	"maunium.net/go/mautrix"
	mcrypto "maunium.net/go/mautrix/crypto"
	mevent "maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
	mid "maunium.net/go/mautrix/id"
)

type BotPlexer struct {
	username       *string
	password       *string // only kept until connect
	client         *mautrix.Client
	timewait       float64
	olmMachine     *mcrypto.OlmMachine
	mostRecentSend map[mid.RoomID]time.Time
	Ch             chan *mevent.Event
	//stateStore *store.StateStore
}

var App BotPlexer
var username string

func NewApp() *BotPlexer {
	return &BotPlexer{
		new(string),
		new(string),
		nil,
		1,
		nil,
		make(map[mid.RoomID]time.Time),
		make(chan *mevent.Event, 8),
	}
}

func (b *BotPlexer) Connect(uname, passwd string) {
	b.timewait = 30
	b.mostRecentSend = make(map[mid.RoomID]time.Time)
	username = mid.UserID(uname).String()
	*b.username = uname
	*b.password = passwd

	log.Infof("Logging in %s", username)

	var err error
	b.client, err = mautrix.NewClient("matrix.privex.io", "", "")
	if err != nil {
		panic(err)
	}
	_, err = DoRetry("login", func() (interface{}, error) {
		return b.client.Login(&mautrix.ReqLogin{
			Type: mautrix.AuthTypePassword,
			Identifier: mautrix.UserIdentifier{
				Type: mautrix.IdentifierTypeUser,
				User: username,
			},
			Password:                 *b.password,
			InitialDeviceDisplayName: "vacation responder",
			//DeviceID:                 deviceID,
			StoreCredentials: true,
		})
	})
	if err != nil {
		log.Fatalf("Couldn't login to the homeserver.")
	}

	syncer := b.client.Syncer.(*mautrix.DefaultSyncer)
	syncer.OnEventType(mevent.EventMessage, func(source mautrix.EventSource, event *mevent.Event) { go b.HandleMessage(source, event) })

	log.Infof("Logged in as %s/%s", b.client.UserID, b.client.DeviceID)

	for {
		log.Debugf("Running sync...")
		err = b.client.Sync()
		if err != nil {
			log.Errorf("Sync failed. %+v", err)
		}
	}
}

func (b *BotPlexer) GetMessages(roomid mid.RoomID, offset int) []*JSONMessage {
	//b.client.Messages
	//TODO
	return nil
}

// There's no goroutine running this function... you have to spawn it somewhere
func (b *BotPlexer) CreateRoom(client *Client) (resp mid.RoomID, err error) {
	response, err := b.client.CreateRoom(&mautrix.ReqCreateRoom{
		Preset:        "public_chat",
		RoomAliasName: (*client.session.Alias) + "_" + (*client.session.SessionId)[:6],
		Topic:         "livechat",
		Invite:        []id.UserID{id.UserID("@osousa:matrix.org")},
	})

	if err != nil {
		return "", err
	}

	return response.RoomID, nil
}

func (b *BotPlexer) JoinRoomByID(rid mid.RoomID) (*mautrix.RespJoinRoom, error) {
	return b.client.JoinRoomByID(rid)
}

func DoRetry(description string, fn func() (interface{}, error)) (interface{}, error) {
	var err error
	b := retry.NewFibonacci(1 * time.Second)
	b = retry.WithMaxRetries(5, b)
	for {
		log.Info("trying: ", description)
		var val interface{}
		val, err = fn()
		if err == nil {
			// Success
			return val, nil
		}
		nextDuration, stop := b.Next()
		// Retrying...
		if stop {
			err = errors.New("%s failed. Retry limit reached. Will not retry.")
			break
		}
		time.Sleep(nextDuration)
	}
	return nil, err
}
func (b *BotPlexer) HandleMessage(source mautrix.EventSource, event *mevent.Event) {
	// If event is from ourselves, ignore
	if event.Sender.String() == *b.username {
		return
	} else {
		b.Ch <- event
	}
	now := time.Now()
	if now.Sub(b.mostRecentSend[event.RoomID]).Minutes() < b.timewait {
		log.Infof("Already sent greetings message to %s in the past %f minutes.", event.RoomID, b.timewait)
		return
	}
	b.mostRecentSend[event.RoomID] = now
}

func (b *BotPlexer) SendMessage(roomId mid.RoomID, content *mevent.MessageEventContent) (resp *mautrix.RespSendEvent, err error) {
	eventContent := &mevent.Content{Parsed: content}
	r, err := DoRetry(fmt.Sprintf("send message to %s", roomId), func() (interface{}, error) {
		//Sending unencrypted event
		return b.client.SendMessageEvent(roomId, mevent.EventMessage, eventContent)
	})
	if err != nil {
		log.Errorf("Failed to send message to %s: %s", roomId, err)
		return nil, err
	}
	return r.(*mautrix.RespSendEvent), err
}
