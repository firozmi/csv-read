package main

import (
	"net/http"

	"bitbucket.org/firozmi/csv-read/src/conf"
	"bitbucket.org/firozmi/csv-read/src/handler"
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

	statusHandle := handler.NewServerStatus(*conf, errlog)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/api/status"), statusHandle.GetServerStatus)

	http.ListenAndServe(conf.Port, mux)
}
