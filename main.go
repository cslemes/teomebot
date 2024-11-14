package main

import (
	"flag"
	"os"
	"teomebot/chat"
	"teomebot/models"
	"teomebot/utils"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/joho/godotenv"
)

func main() {

	migration := flag.Bool("migrations", false, "Realizar migrations do banco de dados")
	flag.Parse()

	godotenv.Load()
	user := os.Getenv("TWITCH_BOT")
	//oauth := os.Getenv("TWITCH_OAUTH_BOT")
	oauth := utils.AccessSecretVersion("projects/551619572964/secrets/twitch-token/versions/latest")
	channel := os.Getenv("TWITCH_CHANNEL")

	if *migration {
		con, err := utils.OpenDBConnection()
		if err != nil {
			panic("erro na conexão com banco")
		}

		con.AutoMigrate(&models.PresentUser{}, &models.StreakPresentUser{}, &models.TwitchUser{}, &models.ProfileUser{})
		return
	}

	client := twitch.NewClient(user, oauth)

	go chat.GetChat(client, channel)
	go chat.RandomWarnings(client, channel)

	for {
		time.Sleep(time.Hour * 1)
	}
}
