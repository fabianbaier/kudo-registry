package start

import (
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/config"
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/storage"
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/util/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// CmdErrorProcessor returns the errors associated with cmd env
func CmdErrorProcessor(cmd *cobra.Command, args []string) error {

	logLevel := os.Getenv("KUDO_REGISTRY_LOGLEVEL")
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Infof("[KUDO REGISTRY] LOG_LEVEL set to debug.")
	}

	backend := backendFromConfig()

	logrus.Infof("Configured backend: %v", vars.StorageBackend)

	server := server.NewServer(backend)

	server.Start()

	return nil
}

func backendFromConfig() storage.Backend {

	var backend storage.Backend

	storageFlag := strings.ToLower(vars.StorageBackend)
	switch storageFlag {
	case "local":
		backend = storage.LocalBackendFromConfig()
	default:
		logrus.Fatal("Unsupported storage backend: ", storageFlag)
	}

	return backend
}
