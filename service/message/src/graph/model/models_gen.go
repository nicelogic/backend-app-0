// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Chat struct {
	ID              string   `json:"id"`
	Type            ChatType `json:"type"`
	Members         []*User  `json:"members"`
	Pinned          bool     `json:"pinned"`
	LastMessage     *Message `json:"lastMessage"`
	LastMessageTime *string  `json:"lastMessageTime"`
	Name            *string  `json:"name"`
}

type ChatConnection struct {
	TotalCount int       `json:"totalCount"`
	Edges      []*Edge   `json:"edges"`
	PageInfo   *PageInfo `json:"pageInfo"`
}

type Edge struct {
	Node   *Chat   `json:"node"`
	Cursor *string `json:"cursor"`
}

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Sender  *User  `json:"sender"`
	Date    string `json:"date"`
}

type MessageConnection struct {
	TotalCount int              `json:"totalCount"`
	Edges      []*MessageEdge   `json:"edges"`
	PageInfo   *MessagePageInfo `json:"pageInfo"`
}

type MessageEdge struct {
	Node   *Message `json:"node"`
	Cursor *string  `json:"cursor"`
}

type MessagePageInfo struct {
	EndCursor   *string `json:"endCursor"`
	HasNextPage bool    `json:"hasNextPage"`
}

type NewChatMessage struct {
	ID      string   `json:"id"`
	Message *Message `json:"message"`
}

type PageInfo struct {
	EndCursor   *string `json:"endCursor"`
	HasNextPage bool    `json:"hasNextPage"`
}

type User struct {
	ID string `json:"id"`
}

type ChatType string

const (
	ChatTypeP2p   ChatType = "P2P"
	ChatTypeGroup ChatType = "GROUP"
)

var AllChatType = []ChatType{
	ChatTypeP2p,
	ChatTypeGroup,
}

func (e ChatType) IsValid() bool {
	switch e {
	case ChatTypeP2p, ChatTypeGroup:
		return true
	}
	return false
}

func (e ChatType) String() string {
	return string(e)
}

func (e *ChatType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ChatType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ChatType", str)
	}
	return nil
}

func (e ChatType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
