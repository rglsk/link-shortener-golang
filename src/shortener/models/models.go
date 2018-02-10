package models

import (
	"time"
)

type UrlHistory struct {
	OriginalUrl string
	ShortUrl string
	Created time.Time
}
