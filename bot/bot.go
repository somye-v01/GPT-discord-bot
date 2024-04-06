package bot

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/somye-v01/discord-gpt-bot/config"
	cohere       "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot running")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	client := cohereclient.NewClient(cohereclient.WithToken(config.AuthToken))
	// ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	// defer cancel()
	response, err := client.Chat(
	context.TODO(),
	&cohere.ChatRequest{
		Message: m.Content,
  	},
	)
	if(err!=nil){
		fmt.Println(err.Error())
		return
	}
	
	// if m.Content == "ping" {
		_,_ = s.ChannelMessageSend(m.ChannelID,response.Text)
	// }
}
