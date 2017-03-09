package main

import (
	"flag"
	"net/http"

	"github.com/Dataman-Cloud/pressure-test-app/api"
	"github.com/Dataman-Cloud/pressure-test-app/config"
	"github.com/Dataman-Cloud/pressure-test-app/store"

	log "github.com/Sirupsen/logrus"
)

var (
	envFile = flag.String("config", "./deploy/env_file", "")
)

func main() {
	flag.Parse()

	conf := config.InitConfig(*envFile)

	store := store.New(conf.DbDriver, conf.DbDSN)
	api := &api.Api{
		Config: conf,
		Store:  store,
	}

	server := &http.Server{
		Addr:           conf.DemoAddr,
		Handler:        api.ApiRouter(),
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("can't start server: ", err)
	}
}
