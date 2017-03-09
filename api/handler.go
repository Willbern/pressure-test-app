package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Dataman-Cloud/pressure-test-app/model"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func (api *Api) Pong(ctx *gin.Context) {
	HttpOkResponse(ctx, "pong")
}

func (api *Api) Add(ctx *gin.Context) {
	var app model.App
	err := ctx.BindJSON(&app)
	if err != nil {
		log.Error("bind json error: ", err.Error())
		HttpErrorResponse(ctx, http.StatusBadRequest, err)
	}

	app.CreateAt = time.Now()
	app.UpdateAt = time.Now()
	err = api.Store.SaveApp(&app)
	if err != nil {
		log.Error("save app error: ", err.Error())
		HttpErrorResponse(ctx, http.StatusServiceUnavailable, err)
	}

	HttpOkResponse(ctx, "OK")
}

func (api *Api) Get(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("convert strin to int error: ", err.Error())
		HttpErrorResponse(ctx, http.StatusServiceUnavailable, err)
	}

	app, err := api.Store.GetApp(id)
	if err != nil {
		log.Error("save app error: ", err.Error())
		HttpErrorResponse(ctx, http.StatusServiceUnavailable, err)
	}

	HttpOkResponse(ctx, app)
}

func (api *Api) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("convert strin to int error: ", err.Error())
		HttpErrorResponse(ctx, http.StatusServiceUnavailable, err)
	}

	var app model.App
	app.ID = id
	err = api.Store.DeleteApp(&app)
	if err != nil {
		log.Error("save app error: ", err.Error())
		HttpErrorResponse(ctx, http.StatusServiceUnavailable, err)
	}

	HttpOkResponse(ctx, "OK")
}

func (api *Api) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("convert strin to int error: ", err.Error())
		HttpErrorResponse(ctx, http.StatusServiceUnavailable, err)
	}

	var app model.App
	err = ctx.BindJSON(&app)
	if err != nil {
		log.Error("update bind json error: ", err.Error())
		HttpErrorResponse(ctx, http.StatusServiceUnavailable, err)
	}
	app.ID = id
	app.UpdateAt = time.Now()

	err = api.Store.UpdateApp(&app)
	if err != nil {
		log.Error("update bind json error: ", err.Error())
		HttpErrorResponse(ctx, http.StatusServiceUnavailable, err)
	}

	HttpOkResponse(ctx, "OK")
}

func HttpOkResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
	return
}

func HttpErrorResponse(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, err.Error())
	return
}
