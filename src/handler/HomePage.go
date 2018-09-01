package handler

import (
	"encoding/csv"
	"encoding/json"
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

//Home handles the uploading and loading of csv
type Home struct {
	dbService service.DBService
	conf      conf.Vars
	log       log.Logger
}

//RespUp contains response
type RespUp struct {
	Upload string `json:"upload"`
	Api    string `json:"api"`
}

//ErrorUp contains error
type ErrorUp struct {
	Upload string `json:"upload"`
	Error  string `json:"error"`
}

type HomeVars struct{}

//NewHomeHandle returns a new Home handler
func NewHomeHandle(c conf.Vars, l log.Logger, ds service.DBService) Home {
	return Home{dbService: ds, conf: c, log: l}
}

//GetHome shows homepage
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

//Upload uploads & loads csv into goleveldb
func (h Home) Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		h.log.Error("error", err.Error())
		return
	}
	defer file.Close()
	f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		h.log.Error("error", err.Error())
		return
	}
	defer f.Close()
	io.Copy(f, file)

	if handler.Header["Content-Type"][0] == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		//load csv to go level db
		keyval, err := getXlsData("./uploads/" + handler.Filename)
		if err != nil {
			h.log.Error("Upload", err.Error())
			return
		}
		err = saveCsvData(h.dbService, keyval)
		if err != nil {
			h.log.Error("Upload", err.Error())
			return
		}
	} else if handler.Header["Content-Type"][0] == "text/csv" {
		keyval, err := getCsvData("./uploads/" + handler.Filename)
		if err != nil {
			h.log.Error("Upload", err.Error())
			return
		}
		err = saveCsvData(h.dbService, keyval)
		if err != nil {
			h.log.Error("Upload", err.Error())
			return
		}
	} else {
		resp := &ErrorUp{
			Upload: "failed",
			Error:  "Incompatible format",
		}
		body, err := json.Marshal(resp)
		if err != nil {
			h.log.Error("Upload", err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", body)
		return
	}

	resp := &RespUp{
		Upload: "success",
		Api:    "http://localhost" + h.conf.Port + "/api/key",
	}
	body, err := json.Marshal(resp)
	if err != nil {
		h.log.Error("Upload", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

	return
}

func getXlsData(fiName string) (map[string]string, error) {
	keyval := map[string]string{}
	xlFile, err := xlsx.OpenFile(fiName)
	if err != nil {
		return nil, err
	}

	sheet := xlFile.Sheets[0]
	for _, row := range sheet.Rows[1:] {
		key := row.Cells[0].String()
		val := row.Cells[1].String()
		if key != "" && val != "" {
			keyval[val] = key
		}
	}
	return keyval, nil
}

func getCsvData(fiName string) (map[string]string, error) {
	keyval := map[string]string{}
	file, err := os.Open(fiName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		records, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if len(records) != 2 {
			continue
		}

		key := records[0]
		val := records[1]

		if key == "key" {
			continue
		}

		if key != "" && val != "" {
			keyval[val] = key
		}
	}
	return keyval, nil
}

func saveCsvData(ds service.DBService, keyval map[string]string) error {
	for key, val := range keyval {
		err := ds.SaveKeyValue(key, val)
		if err != nil {
			return err
		}
	}
	return nil
}
