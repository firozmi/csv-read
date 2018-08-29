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
	homeHandle := handler.NewHomeHandle(*conf, errlog)
	searchHandle := handler.NewSearchHandle(*conf, errlog)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/api/status"), statusHandle.GetServerStatus)

	mux.HandleFunc(pat.Get("/"), homeHandle.GetHome)
	mux.HandleFunc(pat.Post("/upload"), homeHandle.Upload)

	mux.HandleFunc(pat.Get("/api/:key"), searchHandle.SearchKey)

	http.ListenAndServe(conf.Port, mux)
}
