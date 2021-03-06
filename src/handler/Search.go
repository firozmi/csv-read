package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/firozmi/csv-read/src/conf"
	"bitbucket.org/firozmi/csv-read/src/service"
	"github.com/hifx/bingo/infra/log"
	"goji.io/pat"
)

//Search searches for key
type Search struct {
	dbService service.DBService
	conf      conf.Vars
	log       log.Logger
}

//Resp contains response
type Resp struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

//ErrorResp contains response
type ErrorResp struct {
	Error string `json:"error"`
}

//NewSearchHandle returns a new Search Handler
func NewSearchHandle(c conf.Vars, l log.Logger, ds service.DBService) Search {
	return Search{dbService: ds, conf: c, log: l}
}

//SearchKey fetches the value for the key
func (s Search) SearchKey(w http.ResponseWriter, r *http.Request) {
	key := pat.Param(r, "key")
	val, err := s.dbService.GetKeyValue(key)
	var body []byte

	if err != nil {
		s.log.Error("SearchKey", err.Error())
		resp := &ErrorResp{
			Error: "Unable to fetch value",
		}
		body, err = json.Marshal(resp)
		if err != nil {
			s.log.Error("SearchKey", err.Error())
			return
		}
		w.WriteHeader(http.StatusNotFound)
	} else {
		resp := &Resp{
			Key: val,
			Val: key,
		}
		body, err = json.Marshal(resp)
		if err != nil {
			s.log.Error("SearchKey", err.Error())
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", body)

	return
}
