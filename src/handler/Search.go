package handler

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/firozmi/csv-read/src/conf"
	"bitbucket.org/firozmi/csv-read/src/service"
	"github.com/hifx/bingo/infra/log"
	"goji.io/pat"
)

//Search
type Search struct {
	dbService service.DBService
	conf      conf.Vars
	log       log.Logger
}

type Resp struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

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
	} else {
		resp := &Resp{
			Key: key,
			Val: val,
		}
		body, err = json.Marshal(resp)
		if err != nil {
			s.log.Error("SearchKey", err.Error())
			return
		}
	}
	s.dbService.Printall()
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

	return
}
