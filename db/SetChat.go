package db

import (
	"errors"
	"github.com/number571/hiddenlake/models"
	"github.com/number571/hiddenlake/settings"
)

func SetChat(user *models.User, chat *models.Chat) error {
	id := GetUserId(user.Username)
	if id < 0 {
		return errors.New("User id undefined")
	}
	clientid := GetClientId(id, chat.Companion)
	if clientid < 0 {
		return errors.New("Client id undefined")
	}
	for index := range chat.Messages {
		encryptMessage(user, &chat.Messages[index])
		_, err := settings.DB.Exec(
			"INSERT INTO Chat (IdUser, IdClient, Name, Message, LastTime) VALUES ($1, $2, $3, $4, $5)",
			id,
			clientid,
			chat.Messages[index].Name,
			chat.Messages[index].Text,
			chat.Messages[index].Time,
		)
		if err != nil {
			panic("exec 'setchat' failed")
		}
	}
	return nil
}
