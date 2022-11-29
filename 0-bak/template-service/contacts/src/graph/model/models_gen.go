// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddContactsApply struct {
	UserID     string `json:"user_id"`
	ContactsID string `json:"contacts_id"`
	UpdateTime string `json:"update_time"`
	Message    string `json:"message"`
	Reply      string `json:"reply"`
}

type AddContactsApplyConnection struct {
	TotalCount int                           `json:"totalCount"`
	Edges      []*AddContactsApplyEdge       `json:"edges"`
	PageInfo   *AddContactsApplyEdgePageInfo `json:"pageInfo"`
}

type AddContactsApplyEdge struct {
	Node   *AddContactsApply `json:"node"`
	Cursor *string           `json:"cursor"`
}

type AddContactsApplyEdgePageInfo struct {
	EndCursor   *string `json:"endCursor"`
	HasNextPage bool    `json:"hasNextPage"`
}

type ApplyAddContactsInput struct {
	ContactsID string `json:"contactsId"`
	RemarkName string `json:"remarkName"`
	Message    string `json:"message"`
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
	Cursor *string   `json:"cursor"`
}

type PageInfo struct {
	EndCursor   *string `json:"endCursor"`
	HasNextPage bool    `json:"hasNextPage"`
}

type ReplyAddContactsInput struct {
	ContactsID string `json:"contacts_id"`
	Reply      string `json:"reply"`
	RemarkName string `json:"remarkName"`
}