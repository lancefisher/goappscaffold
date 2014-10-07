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

func (g *Greeting) GetRecent(c appengine.Context, greetings *[]Greeting, count int) error {
	q := datastore.NewQuery("Greeting").
		Ancestor(GuestbookKey(c)).
		Order("-Date").
		Limit(count)
	if _, err := q.GetAll(c, greetings); err != nil {
		return err
	}
	return nil
}
