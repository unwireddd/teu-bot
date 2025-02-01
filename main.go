package main

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	pref := tele.Settings{
		Token:  "",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/s", func(c tele.Context) error {

		arguments := c.Message().ReplyTo.Text
		fmt.Println(c.Message().ReplyTo.Text)

		return c.Send(arguments)

	})

	b.Start()

}
