package main

import (
	"log"
	"strings"
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

		if !c.Message().IsReply() {
			return c.Send("Syntax error")
		}

		arguments := c.Message().ReplyTo.Text
		switches := c.Args()

		if len(switches) != 2 {
			return c.Send("Bad syntax")
		}

		if strings.Contains(arguments, switches[0]) {
			newArgs := strings.ReplaceAll(arguments, switches[0], switches[1])
			sender := c.Message().ReplyTo.Sender.FirstName
			//fmt.Println(sender)

			owtput := "<" + sender + ">" + ":" + " " + newArgs

			return c.Send(owtput)
		}

		return c.Send("Syntax error")

	})

	b.Start()

}
