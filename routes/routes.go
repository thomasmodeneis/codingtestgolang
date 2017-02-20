package routes

import (
	"codingtestgolang/model"
	"fmt"
	"github.com/gorilla/pat"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"github.com/modeneis/gocustomparser"
	"golang.org/x/text/encoding/charmap"
)

func Handler() http.Handler {
	m := pat.New()

	m.Get("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		csvParser := gocustomparser.CustomParser{
			File:          "Workbook2.csv",
			CustomDecoder: charmap.ISO8859_1.NewDecoder(),
			Separator:     ',',
			SkipFirstLine: true,
		}
		var csvParsedItems, err = csvParser.Parse(model.Workbook{})
		if err != nil {
			log.Println("Error: Failed to parse Workbook ->", err)
			return
		}

		prnParser := gocustomparser.CustomParser{
			File:          "Workbook2.prn",
			CustomDecoder: charmap.ISO8859_1.NewDecoder(),
			SkipFirstLine: true,
			PRNReader: func(raw string) (line []string, err error) {
				runes := []rune(raw)
				if len(runes) < 74 {
					err = fmt.Errorf("ReadPrnLine detected Wrong data -> %s", raw)
					return
				}
				line = append(line, strings.TrimSpace(string(runes[0:16])))
				line = append(line, strings.TrimSpace(string(runes[16:38])))
				line = append(line, strings.TrimSpace(string(runes[38:47])))
				line = append(line, strings.TrimSpace(string(runes[47:61])))
				line = append(line, strings.TrimSpace(string(runes[61:74])))
				line = append(line, strings.TrimSpace(string(runes[74:])))
				return
			},
		}

		prnParsedItems, err := prnParser.Parse(model.Workbook{})
		if err != nil {
			log.Println("Error: Failed to parse Workbook ->", err)
			return
		}

		dir, _ := os.Getwd()
		t, err := template.ParseFiles(filepath.Join(dir, "templates/index.html"))

		if err != nil {
			fmt.Println(err)
			fmt.Fprintln(w, err)
			return
		}
		t.Execute(w, map[string]interface{}{
			"Csv": csvParsedItems,
			"Prn": prnParsedItems,
		})
	}))

	return m
}
