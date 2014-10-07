package models

import (
	"appengine"
	"appengine/datastore"
	"time"
)

func GuestbookKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}

type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}
