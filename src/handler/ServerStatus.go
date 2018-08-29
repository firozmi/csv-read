package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/firozmi/csv-read/src/conf"
	"golang.org/x/net/context"

	"github.com/firozmi/movie-api/core"
	"github.com/hifx/bingo/infra/log"
)

//ServerStatus gives status of server
type ServerStatus struct {
	conf conf.Vars
	log  log.Logger
}

//NewServerStatus returns a new ServerStatus
func NewServerStatus(c conf.Vars, l log.Logger) ServerStatus {
	return ServerStatus{conf: c, log: l}
}

func (s ServerStatus) GetServerStatus(w http.ResponseWriter, r *http.Request) {
	var serverOb core.ServerStatus

	serverOb.Code = 200
	serverOb.Message = "Status oK"

	//Marshalling Status
	body, err := json.Marshal(serverOb)
	if err != nil {
		s.log.Error(err.Error())
		return
	}
	fmt.Fprintf(w, string(body))
	return
}

func (e ServerStatus) MethodNotAllowedHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	return
}
