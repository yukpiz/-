package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"unicode/utf8"
)

var (
	tmplFile  = "README.tmpl"
	outFile   = "README.md"
	logWriter = os.Stdout
	errCode   = 1
	mdPattern = `^.*\.(md|markdown)$`
	targetDir = flag.String("t", "", "Target directory")
)

func main() {
	flag.Parse()
	if len(strings.TrimSpace(*targetDir)) == 0 {
		fmt.Fprintf(logWriter, "Empty target directory: -t\n")
		os.Exit(errCode)
	}

	paths, err := getMarkdownFiles(*targetDir)
	if err != nil {
		os.Exit(errCode)
	}

	type Entity struct {
		YMD        string
		Year       int
		Month      int
		Day        int
		LinesCount int
		CharCount  int
	}

	var entities []*Entity
	for _, p := range paths {
		b, err := ioutil.ReadFile(p)
		if err != nil {
			fmt.Fprintf(logWriter, "Read file error: %v\n", err)
			continue
		}
		lines := strings.Split(string(b), "\n")

		ymd := strings.Split(p, "/")
		if len(ymd) != 3 {
			fmt.Fprintf(logWriter, "Skip file: %v\n", p)
			continue
		}
		entities = append(entities, &Entity{
			YMD:        fmt.Sprintf("%s-%s-%s", ymd[0], ymd[1], strings.Split(ymd[2], ".")[0]),
			Year:       parseInt(ymd[0]),
			Month:      parseInt(ymd[1]),
			Day:        parseInt(strings.Split(ymd[2], ".")[0]),
			LinesCount: len(lines),
			CharCount:  utf8.RuneCountInString(string(b)),
		})
	}

	type YearSummary struct {
		Year       int
		LinesCount int
		CharCount  int
	}

	type MonthSummary struct {
		YM         string
		Year       int
		Month      int
		LinesCount int
		CharCount  int
	}
	var years []*YearSummary
	var months []*MonthSummary
	var prvYear, prvMonth int
	var y *YearSummary
	var m *MonthSummary
	for _, e := range entities {
		if prvYear != e.Year {
			if y != nil {
				years = append(years, y)
			}
			y = &YearSummary{
				Year:       e.Year,
				LinesCount: e.LinesCount,
				CharCount:  e.CharCount,
			}
		} else if y != nil {
			y.LinesCount += e.LinesCount
			y.CharCount += e.CharCount
		}
		prvYear = e.Year

		if prvMonth != e.Month {
			if m != nil {
				months = append(months, m)
			}
			m = &MonthSummary{
				YM:         fmt.Sprintf("%d-%02d", e.Year, e.Month),
				Year:       e.Year,
				Month:      e.Month,
				LinesCount: e.LinesCount,
				CharCount:  e.CharCount,
			}
		} else if m != nil {
			m.LinesCount += e.LinesCount
			m.CharCount += e.CharCount
		}
		prvMonth = e.Month
	}
	if y != nil {
		years = append(years, y)
	}
	if m != nil {
		months = append(months, m)
	}

	type EmbedData struct {
		Entities       []*Entity
		YearSummaries  []*YearSummary
		MonthSummaries []*MonthSummary
	}

	f, err := os.Create(outFile)
	if err != nil {
		fmt.Fprintf(logWriter, "Create output file error: %v\n", err)
		os.Exit(errCode)
	}
	t, err := template.ParseFiles(tmplFile)
	if err != nil {
		fmt.Fprintf(logWriter, "Read template file error: %v\n", err)
		os.Exit(errCode)
	}
	if err := t.Execute(f, EmbedData{
		Entities:       entities,
		YearSummaries:  years,
		MonthSummaries: months,
	}); err != nil {
		fmt.Fprintf(logWriter, "Execute template error: %v\n", err)
		os.Exit(errCode)
	}
}

func parseInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		fmt.Fprintf(logWriter, "Parse int error: %v\n", err)
		return 0
	}
	return i
}

func getMarkdownFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			ps, err := getMarkdownFiles(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			paths = append(paths, ps...)
			continue
		}

		reg := regexp.MustCompile(mdPattern)
		if reg.MatchString(file.Name()) {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths, nil
}
