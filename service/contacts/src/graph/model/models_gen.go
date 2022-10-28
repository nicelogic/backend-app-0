// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddContactsApply struct {
	ID         string  `json:"id"`
	ContactsID string  `json:"contactsId"`
	Message    *string `json:"message"`
}

type ApplyAddContactsInput struct {
	ContactsID string `json:"contactsId"`
	RemarkName string `json:"remarkName"`
}

type Contacts struct {
	ID         string `json:"id"`
	RemarkName string `json:"remarkName"`
}

type ContactsConnection struct {
	TotalCount int       `json:"totalCount"`
	Edges      []*Edge   `json:"edges"`
	PageInfo   *PageInfo `json:"pageInfo"`
}

type Edge struct {
	Node   *Contacts `json:"node"`
	Cursor string    `json:"cursor"`
}

type PageInfo struct {
	EndCursor   *string `json:"endCursor"`
	HasNextPage bool    `json:"hasNextPage"`
}

type ReplyAddContactsInput struct {
	ID         string  `json:"id"`
	Ack        bool    `json:"ack"`
	RemarkName *string `json:"remarkName"`
}
