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
	"maunium.net/go/mautrix/format"
	"maunium.net/go/mautrix/id"
	mid "maunium.net/go/mautrix/id"
)

type BotPlexer struct {
	client *mautrix.Client
	//configuration Configuration
	timewait   float64
	olmMachine *mcrypto.OlmMachine
	//stateStore *store.StateStore

	// Most recent send to room.
	mostRecentSend map[mid.RoomID]time.Time
	ch             chan *Client
}

var App BotPlexer
var username string

func NewApp(ch chan *Client) *BotPlexer {
	return &BotPlexer{
		nil,
		1,
		nil,
		make(map[mid.RoomID]time.Time),
		ch,
	}
}

func (b *BotPlexer) Connect(username, password string) {
	b.timewait = 1
	b.mostRecentSend = make(map[mid.RoomID]time.Time)
	username = mid.UserID("@osousa:privex.io").String()
	password = "6VrdT8DCsa1xDvyaOghxT"

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
			Password:                 password,
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

// lol there's no goroutine running this function... you have to spawn it somewhere duh...
func (b *BotPlexer) CreateRoom(client *Client) (resp mid.RoomID, err error) {
	response, err := b.client.CreateRoom(&mautrix.ReqCreateRoom{
		Preset:        "public_chat",
		RoomAliasName: "mao_tseng" + (*client.session.SessionId)[:5],
		Topic:         "Topic_yeah",
		Invite:        []id.UserID{id.UserID("@osousa:matrix.org")},
	})

	if err != nil {
		return "", err
	}

	return response.RoomID, nil
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
			log.Info(description, " succeeded")
			return val, nil
		}
		nextDuration, stop := b.Next()
		log.Debugf("  %s failed. Retrying in %f seconds...", description, nextDuration.Seconds())
		if stop {
			log.Debugf("  %s failed. Retry limit reached. Will not retry.", description)
			err = errors.New("%s failed. Retry limit reached. Will not retry.")
			break
		}
		time.Sleep(nextDuration)
	}
	return nil, err
}
func (b *BotPlexer) HandleMessage(source mautrix.EventSource, event *mevent.Event) {
	if event.Sender.String() == username {
		log.Infof("Event %s is from us, so not going to respond.", event.ID)
		return
	} else {
		log.Infof("Event %s is from someone else.", event.ID)
	}

	//if !b.configuration.RespondToGroups && len(App.stateStore.GetRoomMembers(event.RoomID)) != 2 {
	//	log.Infof("Event %s is not from in a DM, so not going to respond.", event.ID)
	//	return
	//}

	now := time.Now()
	if now.Sub(b.mostRecentSend[event.RoomID]).Minutes() < b.timewait {
		log.Infof("Already sent a vacation message to %s in the past %f minutes.", event.RoomID, b.timewait)
		return
	}
	b.mostRecentSend[event.RoomID] = now

	content := format.RenderMarkdown("Test Message goes here...", true, true)
	b.SendMessage(event.RoomID, &content)
}

func (b *BotPlexer) SendMessage(roomId mid.RoomID, content *mevent.MessageEventContent) (resp *mautrix.RespSendEvent, err error) {
	eventContent := &mevent.Content{Parsed: content}
	r, err := DoRetry(fmt.Sprintf("send message to %s", roomId), func() (interface{}, error) {
		log.Debugf("Sending unencrypted event to %s", roomId)
		return b.client.SendMessageEvent(roomId, mevent.EventMessage, eventContent)
	})
	if err != nil {
		// give up
		log.Errorf("Failed to send message to %s: %s", roomId, err)
		return nil, err
	}
	return r.(*mautrix.RespSendEvent), err
}
