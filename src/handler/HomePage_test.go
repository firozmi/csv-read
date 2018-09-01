package handler

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetXlsData(t *testing.T) {
	fiName := os.Getenv("GOPATH") + "/src/bitbucket.org/firozmi/csv-read/src/test_files/Corpus.CSV.XLSX"
	keyval := [][]string{
		{"one", "1"},
		{"three", "3"},
		{"hundred", "100"},
		{"five", "5"},
		{"twenty", "20"},
		{"fifty", "50"},
	}

	assert := assert.New(t)

	csvKeyVal, err := getXlsData(fiName)
	if err != nil {
		t.Errorf("Reading xls got error: %s", err.Error())
	}

	for _, kval := range keyval {
		assert.Equal(csvKeyVal[kval[1]], kval[0], "They should be equal.")
	}
}

func TestGetCsvData(t *testing.T) {
	fiName := os.Getenv("GOPATH") + "/src/bitbucket.org/firozmi/csv-read/src/test_files/Corpus.CSV.csv"
	keyval := [][]string{
		{"one", "1"},
		{"three", "3"},
		{"hundred", "100"},
		{"five", "5"},
		{"twenty", "20"},
		{"fifty", "50"},
	}

	assert := assert.New(t)

	csvKeyVal, err := getCsvData(fiName)
	if err != nil {
		t.Errorf("Reading xls got error: %s", err.Error())
	}

	for _, kval := range keyval {
		assert.Equal(csvKeyVal[kval[1]], kval[0], "They should be equal.")
	}
}
