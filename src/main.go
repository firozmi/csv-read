package main

import (
	"net/http"
	"os"

	"bitbucket.org/firozmi/csv-read/src/conf"
	"bitbucket.org/firozmi/csv-read/src/handler"
	"bitbucket.org/firozmi/csv-read/src/service"
	"github.com/hifx/bingo/infra/log"
	goji "goji.io"
	"goji.io/pat"
)

const appname = "csvread"

func main() {
	conf := conf.Read(appname)

	//Setup logging
	var errlog log.Logger
	errlog = log.NewLogfmtLogger(conf.Log.Error)
	errlog = errlog.With("app", conf.App)

	dbService, err := service.NewDBService(*conf, errlog)
	if err != nil {
		errlog.Error("leveldb", err.Error())
		os.Exit(1)
	}

	statusHandle := handler.NewServerStatus(*conf, errlog)
	homeHandle := handler.NewHomeHandle(*conf, errlog, dbService)
	searchHandle := handler.NewSearchHandle(*conf, errlog, dbService)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/api/status"), statusHandle.GetServerStatus)

	mux.HandleFunc(pat.Get("/"), homeHandle.GetHome)
	mux.HandleFunc(pat.Post("/upload"), homeHandle.Upload)

	mux.HandleFunc(pat.Get("/api/:key"), searchHandle.SearchKey)

	http.ListenAndServe(conf.Port, mux)
}
