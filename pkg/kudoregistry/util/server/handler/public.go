package handler

import (
	"fmt"
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/storage"
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/util/server/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HealthPublic(c *gin.Context) {
	c.JSON(http.StatusOK, types.Response{
		Data: "Welcome to the KUDO Framework Registry of the future!",
	})
}

func GetFile(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		logrus.Debug("[HANDLER] filename empty.")
		c.JSON(http.StatusNotFound, types.ResponseError{
			Data: types.FileNotFound,
		})
		return
	}

	backendInterface, exists := c.Get("backend")
	if !exists {
		logrus.Debug("[HANDLER] Storage Backend does not exist.")
		c.JSON(http.StatusNotFound, types.ResponseError{
			Data: types.BackendNotFound,
		})
		return
	}

	b, err := castBackend(backendInterface)
	if err != nil {
		logrus.Debugf("[HANDLER] Cannot cast Backend Interface: %+v.", err)
		c.JSON(http.StatusNotFound, types.ResponseError{
			Data: types.BackendNotFound,
		})
		return
	}

	obj, err := storage.Backend.GetObject(b, filename)
	if err != nil {
		logrus.Infof("[HANDLER] Error: %+v", err)
		c.JSON(http.StatusNotFound, types.ResponseError{
			Data: err.Error(),
		})
		return
	}
	c.Data(200, obj.ContentType, obj.Content)

}

func castBackend(i interface{}) (storage.Backend, error) {
	switch i.(type) {
	case storage.Backend:
		b, _ := i.(storage.Backend)
		return b, nil
	default:
		return nil, fmt.Errorf("Interface is not a Storage Backend.")
	}
}
