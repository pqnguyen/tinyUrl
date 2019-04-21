package cache

import "time"

type Repository interface {
	GetOriginalUrl(hash string) (string, bool)
	SetUrl(hash string, originalUrl string, d time.Duration)
}
