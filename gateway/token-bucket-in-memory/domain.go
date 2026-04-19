package tokenbucketinmemory

import "time"

type bucket struct {
	token          int
	lastRefillTime time.Time
}
