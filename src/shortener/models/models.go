package models

import (
	"appengine"
	"appengine/datastore"
	"time"
)

type UrlHistory struct {
	OriginalUrl string
	ShortUrl string
	Created time.Time
}


func (his *UrlHistory) CreateKey(ctx appengine.Context) datastore.Key {
	key := datastore.NewKey(ctx, "UrlHistory", "stringID", 0, nil)
	return *key
}
