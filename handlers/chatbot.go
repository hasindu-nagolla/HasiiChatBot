package handlers

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/hasindu-nagolla/HasiiChatBot/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	tele "gopkg.in/telebot.v3"
)

type ChatWord struct {
	Word  string `bson:"word"`
	Text  string `bson:"text"`
	Check string `bson:"check"`
	ID    string `bson:"id,omitempty"`
}

type Hasii struct {
	ChatID int64 `bson:"chat_id"`
}

func RegisterChatbot(b *tele.Bot, db *database.Database) {
	handleQuery := func(c tele.Context, queryWord string) error {
		// check if disabled in this chat
		var hasii Hasii
		err := db.Hasii.FindOne(context.Background(), bson.M{"chat_id": c.Chat().ID}).Decode(&hasii)
		if err == nil {
			// bot is off here
			return nil
		}

		cursor, err := db.WordDb.Find(context.Background(), bson.M{"word": queryWord})
		if err != nil {
			return err
		}

		var results []ChatWord
		if err = cursor.All(context.Background(), &results); err != nil {
			return err
		}

		if len(results) > 0 {
			rand.Seed(time.Now().UnixNano())
			selected := results[rand.Intn(len(results))]

			isSticker := false
			if selected.Check == "sticker" {
				isSticker = true
			} else if selected.Check == "none" && len(selected.Text) > 30 && !strings.Contains(selected.Text, " ") {
				// handle old python bot data
				isSticker = true
			}

			if isSticker {
				sticker := &tele.Sticker{File: tele.File{FileID: selected.Text}}
				return c.Reply(sticker)
			} else {
				return c.Reply(selected.Text)
			}
		}

		return nil
	}

	learn := func(c tele.Context) {
		msg := c.Message()
		if !msg.IsReply() {
			return
		}
		replyTo := msg.ReplyTo
		// ignore self
		if replyTo.Sender.ID == b.Me.ID {
			return
		}

		var word string
		if replyTo.Text != "" {
			word = replyTo.Text
		} else if replyTo.Sticker != nil {
			word = replyTo.Sticker.UniqueID
		} else {
			return
		}

		var text string
		var check string
		if msg.Text != "" {
			text = msg.Text
			check = "text"
		} else if msg.Sticker != nil {
			text = msg.Sticker.FileID
			check = "sticker"
		} else {
			return
		}

		// avoid duplicates
		var existing ChatWord
		err := db.WordDb.FindOne(context.Background(), bson.M{"word": word, "text": text}).Decode(&existing)
		if err == mongo.ErrNoDocuments {
			// save new reply
			_, _ = db.WordDb.InsertOne(context.Background(), bson.M{
				"word":  word,
				"text":  text,
				"check": check,
			})
		}
	}

	b.Handle(tele.OnText, func(c tele.Context) error {
		text := c.Message().Text

		// skip commands
		if strings.HasPrefix(text, "/") || strings.HasPrefix(text, "!") || strings.HasPrefix(text, "?") {
			return nil
		}

		// learn if reply
		learn(c)

		msg := c.Message()
		if msg.IsReply() && msg.ReplyTo.Sender.ID != b.Me.ID {
			// ignore replies to other people (only learn)
			return nil
		}

		return handleQuery(c, text)
	})

	b.Handle(tele.OnSticker, func(c tele.Context) error {
		if c.Message().Sticker == nil {
			return nil
		}

		// learn if reply
		learn(c)

		msg := c.Message()
		if msg.IsReply() && msg.ReplyTo.Sender.ID != b.Me.ID {
			// ignore replies to other people (only learn)
			return nil
		}

		// use sticker unique id
		stickerID := c.Message().Sticker.UniqueID
		return handleQuery(c, stickerID)
	})
}
