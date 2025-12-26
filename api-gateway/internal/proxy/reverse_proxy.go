package proxy

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Forward(target string) gin.HandlerFunc {
	u, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(u)

	return func(c *gin.Context) {
		c.Request.Host = u.Host
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
