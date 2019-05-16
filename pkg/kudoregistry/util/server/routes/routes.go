package routes

import (
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/storage"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, b storage.Backend) {
	// Handle the routes
	initializePublicRoutes(r, b)
}
