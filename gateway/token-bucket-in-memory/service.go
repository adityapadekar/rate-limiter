package tokenbucketinmemory

import (
	"fmt"
	"log/slog"
	"sync"
)

type service struct {
	data map[string]map[string]bucket
	sync.Mutex
}

type Service interface {
	RateLimitRequest(ipAddr, pattern string) bool
}

func NewService() Service {
	return &service{
		data: make(map[string]map[string]bucket),
	}
}

func (s *service) RateLimitRequest(ipAddr, pattern string) bool {
	fmt.Println(s.data)
	s.Lock()
	defer s.Unlock()
	if _, ok := s.data[ipAddr]; !ok {
		slog.Info("no entry found", "ipAddr", ipAddr)
		s.data[ipAddr] = make(map[string]bucket)
	}

	if _, ok := s.data[ipAddr][pattern]; !ok {
		slog.Info("no bucket found", "pattern", pattern)
		s.data[ipAddr][pattern] = rateLimiterConfig[pattern]
	}

	rateLimitBucket := s.data[ipAddr][pattern]
	if rateLimitBucket.token == 0 {
		slog.Info("rate limit reached", "ipAddr", ipAddr, "pattern", pattern)
		return false
	}

	rateLimitBucket.token--
	s.data[ipAddr][pattern] = rateLimitBucket

	return true
}
