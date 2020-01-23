package limiter

import (
	"time"

	"github.com/alcalbg/buffalo"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
)

// LimitHandler is a middleware that performs request rate-limiting on max per second basis
func Limiter(maxPerSecond float64) buffalo.MiddlewareFunc {
	lmt := tollbooth.NewLimiter(maxPerSecond, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	lmt.SetIPLookups([]string{"CF-Connecting-Ip", "X-Forwarded-For", "RemoteAddr", "X-Real-IP"})
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			httpError := tollbooth.LimitByRequest(lmt, c.Response(), c.Request())
			if httpError != nil {
				lmt.ExecOnLimitReached(c.Response(), c.Request())
				return c.Redirect(401, "")
			}
			return next(c)
		}
	}
}
