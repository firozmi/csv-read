package handler

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"bitbucket.org/firozmi/csv-read/src/conf"
	"github.com/hifx/bingo/infra/log"
	"github.com/tealeg/xlsx"
)

//Home
type Home struct {
	conf conf.Vars
	log  log.Logger
}

type HomeVars struct{}

//NewHomeHandle returns a new Home handler
func NewHomeHandle(c conf.Vars, l log.Logger) Home {
	return Home{conf: c, log: l}
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
	err = loadCsvData("./uploads/" + handler.Filename)

	if err != nil {
		h.log.Error("error", err.Error())
		return
	}

	return
}

func loadCsvData(fiName string) error {
	xlFile, err := xlsx.OpenFile(fiName)
	if err != nil {
		return err
	}

	sheet := xlFile.Sheets[0]
	for _, row := range sheet.Rows {
		key := row.Cells[0].String()
		val := row.Cells[1].String()

	}
}
