package routes

import (
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/storage"
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/util/server/handler"
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/util/server/middleware"
	"github.com/gin-gonic/gin"
)

func initializePublicRoutes(r *gin.Engine, b storage.Backend) *gin.RouterGroup {

	public := r.Group("/api/v1")
	public.Use(middleware.BackendMiddleware(b))
	{
		public.GET("/health", handler.HealthPublic)
		public.GET("/repo/:filename", handler.GetFile)
	}
	return public
}
