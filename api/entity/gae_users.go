package entity

import (
	"time"
)
const GAEUSERS_ENTITYNAME = "GaeUsers"

type GaeUsers struct {
	Id       int64 `datastore:"-" goon:"id" json:"id,omitempty"`
	Name  string  `datastore:"name"  json:"name,omitempty"`
	Email string    `datastore:"email" json:"email,omitempty"`
	CreateDate time.Time `datastore:"create_date,noindex" json:"createDate,omitempty"`
	UpdateDate time.Time `datastore:"update_date,noindex" json:"updateDate,omitempty"`
}