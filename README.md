## Buffalo rate-limiter
Rate-limiter Middleware for Buffalo based on https://github.com/didip/tollbooth

## Usage and config
Add the middleware in your App() like this:

```
import (
...
limiter "github.com/alcalbg/buffalo-rate-limiter-mw"
...
)

...

// List of places to look up IP addresses
// If your application is behind a proxy, set "X-Forwarded-For" first
// If you use CloudFlare, set "CF-Connecting-Ip" first  
IPLookups := []string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}

// Maximum 5 requests per second
maxRequestsPerSecond = 5

app.Use(limiter.Limiter(maxRequestsPerSecond, IPLookups))
```

