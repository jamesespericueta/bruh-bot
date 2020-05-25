package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	str "strings"
	"syscall"

	toml "github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
)

var counter int
var numBruhs int

type baseConfig struct {
	Bruh bruhStruct
}

type bruhStruct struct {
	Token string
}
type statsStruct struct {
	BruhsServer1 int
	BruhsServer2 int
}

func main() {

	counter = 0

	dg, err := discordgo.New("Bot " + decodeConfig())

	if err != nil {
		fmt.Println("Couldn't create discord session,", err)
		return
	}
	stats := strconv.FormatInt(int64(decodeBruhStats()), 10)
	fmt.Println(stats)
	dg.AddHandler(messageCreate)

	//just another error check
	err = dg.Open()

	if err != nil {
		fmt.Println("oopsie, couldn't open the connection,", err)
		return
	}

	fmt.Println("Working")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, os.Interrupt, os.Kill)
	<-c
	dg.Close()
}

/*
first store the
*/
func decodeBruhStats() int {
	var bruhStats statsStruct
	if _, err := toml.DecodeFile(".\\bruhStats.toml", &bruhStats); err != nil {
		fmt.Println(err)
	}
	return bruhStats.BruhsServer2
}

func decodeConfig() string {
	var config baseConfig
	if _, err := toml.DecodeFile(".\\config.toml", &config); err != nil {
		fmt.Println(err)
	}
	return config.Bruh.Token
}

func presenceUpdate(s *discordgo.Session, m *discordgo.PresenceUpdate) {

}

/*
updates the bruhs from the last time that the bot was shutdown
*/
func updateBruhs() int {
	return 0
}

/*
messageCreate simply in the event that a message is created
*/
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "there's been "
	if m.Author.Bot {
		return
	} else if m.Content == "!bruh" {
		finalMessage := fmt.Sprintf("%s%d%s", message, numBruhs, " bruhs")
		if numBruhs == 1 {
			finalMessage = fmt.Sprintf("%s%d%s", "there's been ", numBruhs, " bruh")
		}
		s.ChannelMessageSend(m.ChannelID, finalMessage)
		fmt.Println(m.GuildID)
	} else if str.Contains(m.Content, "bruh") {
		numBruhs += str.Count(m.Content, "bruh")
	}
	if numBruhs == 69 {
		s.ChannelMessageSend(m.Content, "69 bruhs, nice.")
	}
	if numBruhs == 420 {
		s.ChannelMessageSend(m.Content, "420 bruhs, ahaha.")
	}
	if numBruhs == 420 {
		s.ChannelMessageSend(m.Content, "69420 bruhs, type bruh again and I'll find you...")
	}
}

/*TODO:
create an array which stores the guildID and then a corresponding
array which stores the amount of messages
*/
