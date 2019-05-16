package server

import (
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/storage"
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/util/server/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type (
	// ServerOptions are options for constructing a Server
	ServerOptions struct {
		StorageBackend storage.Backend
	}
)

func NewServer(backend storage.Backend) ServerOptions {
	return ServerOptions{
		StorageBackend: backend,
	}
}

func (s *ServerOptions) Start() {

	ginMode := os.Getenv("KUDO_REGISTRY_GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
		logrus.Debugf("[KUDO REGISTRY] Switched to production mode: GIN_MODE=%+v", ginMode)
	}

	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	c.AddAllowHeaders("Authorization")

	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		cors.New(c),
	)

	routes.InitializeRoutes(r, s.StorageBackend)

	logrus.Infof("Listening on http://localhost:50051")
	http.ListenAndServe("0.0.0.0:50051", r)

}
