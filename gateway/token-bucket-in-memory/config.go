package tokenbucketinmemory

var rateLimiterConfig = map[string]bucket{
	"GET /route1": {
		token: 10,
	},
	"GET /route2": {
		token: 5,
	},
	"GET /route3": {
		token: 20,
	},
	"GET /route4": {
		token: 15,
	},
}
