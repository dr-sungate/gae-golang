package entity

import (
	"time"
)
const GAEUSERS_ENTITYNAME = "gae-users"

type GaeUsers struct {
	Id       int64 `datastore:"-" goon:"id"`
	Name  string  `datastore:"name"`
	Email string    `datastore:"email"`
	CreateDate time.Time `datastore:"create_date,noindex"`
	UpdateDate time.Time `datastore:"update_date,noindex"`
}