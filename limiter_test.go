package limiter

import (
	"testing"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/httptest"
	"github.com/stretchr/testify/require"
)

func limitedApp() *buffalo.App {
	h := func(c buffalo.Context) error {
		return c.Render(200, nil)
	}

	a := buffalo.New(buffalo.Options{})
	a.Use(Limiter(3, []string{"X-Forwarded-For"}))
	a.GET("/protected", h)
	return a
}

func Test_Limiter(t *testing.T) {
	r := require.New(t)

	w := httptest.New(limitedApp())
	req := w.HTML("/protected")
	req.Headers["X-Forwarded-For"] = "10.10.10.10"

	res := req.Get()
	r.Equal(200, res.Code)

	res = req.Get()
	res = req.Get()
	res = req.Get() // treshold reached!
	res = req.Get()

	r.Equal(429, res.Code)
}
