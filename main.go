package main

import (
	"fmt"
	"os"
	"os/signal"
	str "strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := "Njg2Mzk1MjY1NzI0NDQ4ODAw.XsIqig.GYiMEI7R-LHl5SAXyisbCuNRshY"
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("oh fuck we couldn't create the discord session,", err)
		return
	}
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		fmt.Println("oopsie, couldn't open the connection,", err)
		return
	}
	fmt.Println("we vibin")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, os.Interrupt, os.Kill)
	<-c

	dg.Close()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "cock and ball torture bruh"
	if m.Author.ID == s.State.User.ID {
		return
	}
	if str.Contains(m.Content, "bruh") {
		s.ChannelMessageSend(m.ChannelID, message)
	}
}