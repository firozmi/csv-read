package handler

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"bitbucket.org/firozmi/csv-read/src/conf"
	"bitbucket.org/firozmi/csv-read/src/service"
	"github.com/hifx/bingo/infra/log"
	"github.com/tealeg/xlsx"
)

//Home
type Home struct {
	dbService service.DBService
	conf      conf.Vars
	log       log.Logger
}

type HomeVars struct{}

//NewHomeHandle returns a new Home handler
func NewHomeHandle(c conf.Vars, l log.Logger, ds service.DBService) Home {
	return Home{dbService: ds, conf: c, log: l}
}

func (h Home) GetHome(w http.ResponseWriter, r *http.Request) {
	HomeVars := HomeVars{}

	t, err := template.ParseFiles("src/static/home.html")
	if err != nil {
		h.log.Error("error", err.Error())
	}
	err = t.Execute(w, HomeVars)
	if err != nil {
		h.log.Error("error", err.Error())
	}
}

func (h Home) Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		h.log.Error("error", err.Error())
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		h.log.Error("error", err.Error())
		return
	}
	defer f.Close()
	io.Copy(f, file)

	//load csv to go level db
	h.loadCsvData("./uploads/" + handler.Filename)

	fmt.Fprintf(w, "Api active now: http://localhost"+h.conf.Port+"/api/key")
	return
}

func (h Home) loadCsvData(fiName string) {
	xlFile, err := xlsx.OpenFile(fiName)
	if err != nil {
		return err
	}

	sheet := xlFile.Sheets[0]
	for _, row := range sheet.Rows[1:] {
		key := row.Cells[0].String()
		val := row.Cells[1].String()
		if key != "" && val != "" {
			/* saving value as key, and key as value
			ie, "1" is key and "one" is value */
			err = h.dbService.SaveKeyValue(val, key)
			if err != nil {
				h.log.Error("dbService", err.Error())
			}
		}
	}
}