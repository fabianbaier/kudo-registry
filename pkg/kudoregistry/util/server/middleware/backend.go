package middleware

import (
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/storage"
	"github.com/gin-gonic/gin"
)

func BackendMiddleware(b storage.Backend) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("backend", b)
		c.Next()
	}
}
