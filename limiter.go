package limiter

import (
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gobuffalo/buffalo"
)

// Limiter is a middleware that performs rate-limiting
func Limiter(maxPerSecond float64, IPLookups []string) buffalo.MiddlewareFunc {
	lmt := tollbooth.NewLimiter(maxPerSecond, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	lmt.SetIPLookups(IPLookups)
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
