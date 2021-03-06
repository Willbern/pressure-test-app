package api

import (
	"github.com/Dataman-Cloud/pressure-test-app/config"
	"github.com/Dataman-Cloud/pressure-test-app/store"

	//"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Config *config.Config
	Store  *store.Store
}

func (api *Api) ApiRouter() *gin.Engine {
	router := gin.New()

	router.GET("/ping", api.Pong)
	router.GET("/json", api.ResponseJson)
	router.POST("/add", api.Add)
	router.GET("/get/:id", api.Get)
	router.DELETE("/delete/:id", api.Delete)
	router.PUT("/update/:id", api.Update)

	//router.GET("/static", api.StaticHandler)

	return router
}
