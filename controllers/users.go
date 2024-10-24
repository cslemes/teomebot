package controllers

import (
	"log"
	"teomebot/models"
	"teomebot/services"

	"github.com/gempir/go-twitch-irc/v4"
)

func ExecCreateOrUpdateUser(twitchUser *twitch.User) error {

	user, err := models.GetUserByField("twitch_id", twitchUser.ID, conDB)
	if err != nil {

		user, err = models.GetUserByField("twitch_nick", twitchUser.Name, conDB)
		if err != nil {

			userID, err := services.CreateUser(twitchUser.ID)
			if err != nil {
				log.Println(err)
				return err
			}

			user = &models.TwitchUser{
				UUID:       userID,
				TwitchId:   twitchUser.ID,
				TwitchNick: twitchUser.Name,
			}

			if err := user.Create(conDB); err != nil {
				return err
			}

			return nil
		}
	}

	user.TwitchId = twitchUser.ID
	user.TwitchNick = twitchUser.Name
	return user.Update(conDB)
}
