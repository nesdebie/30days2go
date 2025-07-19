package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func Run() {
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("Error message")
	}

	discord.AddHandler(newMessage)

	discord.Open()
	defer discord.Close()

	// keep bot running until 'ctrl + C'
	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// prevent bot responding to itself own message
	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
		case strings.Contains(message.Content, "!hello"):
			discord.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Hello %s (≖ ͜ʖ≖)", message.Author.Username))
		case strings.Contains(message.Content, "!help"):
			discord.ChannelMessageSend(message.ChannelID, "RTFM (ﾒ￣▽￣)︻┳═一	- - -")
		case strings.Contains(message.Content, "!bye"):
			discord.ChannelMessageSend(message.ChannelID, "Adios (￣▽￣)ノ")
	}
}