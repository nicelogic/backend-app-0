package sql

import (
	"encoding/json"
	"log"
	"message/constant"
	"message/graph/model"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

func ChatRowToChatModel(chatRow pgx.Row) (chatModel *model.Chat, err error) {
	var id, Type string
	var memberIds []string
	var name, last_message *string
	// var last_message_time *time.Time
	var last_message_time pgtype.Timestamptz
	var priority int
	err = chatRow.Scan(&id, &Type, &memberIds, &name, &last_message, &last_message_time, &priority)
	if err != nil {
		log.Printf("query row err(%v)\n", err)
		return nil, err
	} else {
		members := []*model.User{}
		for _, memberId := range memberIds {
			members = append(members, &model.User{ID: memberId})
		}
		var lastMessage model.Message
		if last_message != nil {
			err = json.Unmarshal([]byte(*last_message), &lastMessage)
			if err != nil {
				log.Printf("json unmarshal last_message(%s) err(%v)\n", *last_message, err)
				return nil, err
			}
		}
		pinned := false
		if priority == constant.PriorityPinned {
			pinned = true
		}
		last_message_time_string := ""
		if last_message_time.Status != pgtype.Null {
			last_message_time_string = last_message_time.Time.Format(time.RFC3339)
		}
		return &model.Chat{
			ID:              id,
			Type:            model.ChatType(Type),
			Members:         members,
			Name:            name,
			LastMessage:     &lastMessage,
			LastMessageTime: &last_message_time_string,
			Pinned:          pinned,
		}, nil
	}
}
