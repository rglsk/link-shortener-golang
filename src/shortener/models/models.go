package models

import (
	"time"
	"appengine"
	"appengine/datastore"
	"shortener/generator"
)

type UrlHistory struct {
	OriginalUrl string
	ShortUrl string
	Created time.Time
}


func (h *UrlHistory) CreateKey(ctx appengine.Context) (key *datastore.Key ){
	urlSuffix := generator.GenerateUrlSuffix()
	k := datastore.NewKey(ctx, "UrlHistory", urlSuffix, 0, nil)
	if err := datastore.Get(ctx, k, *h); err != nil {
		h.ShortUrl = urlSuffix
		return k
	}
	return h.CreateKey(ctx)
}
