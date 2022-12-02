package sql

import (
	"encoding/json"
	"log"
	"message/graph/model"

	"github.com/jackc/pgx/v4"
)

func ChatRowToChatModel(chatRow pgx.Row)(chatModel *model.Chat, err error){
	var id, Type string
	var memberIds []string
	var name, last_message *string
	err = chatRow.Scan(&id, &Type, &memberIds, &name, &last_message)
	if err != nil {
		log.Printf("query row err: %v\n", err)
		return nil, err
	} else {
		members := []*model.User{}
		for _, memberId := range memberIds {
			members = append(members, &model.User{ID: memberId})
		}
		var lastMessage *model.Message
		if last_message != nil {
			err = json.Unmarshal([]byte(*last_message), lastMessage)
			if err != nil {
				log.Printf("json unmarshal last_message(%s) err(%v)\n", *last_message, err)
				return nil, err
			}
		}
		return &model.Chat{
			ID:          id,
			Type: model.ChatType(Type),
			Members:     members,
			Name:        name,
			LastMessage: lastMessage,
		}, nil
	}
}