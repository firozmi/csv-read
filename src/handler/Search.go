package handler

import (
	"fmt"
	"net/http"

	"bitbucket.org/firozmi/csv-read/src/conf"
	"github.com/hifx/bingo/infra/log"
)

//Search
type Search struct {
	conf conf.Vars
	log  log.Logger
}

//NewSearchHandle returns a new Search Handler
func NewSearchHandle(c conf.Vars, l log.Logger) Search {
	return Search{conf: c, log: l}
}

//SearchKey fetches the value for the key
func (s Search) SearchKey(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("hi"))
	return
}
