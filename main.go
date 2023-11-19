package main

import (
	"clyeung/har-to-csv/har"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var dataDir string = "data"
var origin string = ""
var priority string = ""
var outputPath string = "output/result.csv"

func main() {
	files := readDir(dataDir)

	fmt.Println("--START read data dir--")
	var req []har.Request
	for _, f := range files {
		fmt.Printf("--READ %s--\n", f)
		b := readFile(filepath.Join("data", f))
		h := har.ReadHar(b)
		req = append(req, har.FilterRequest(h, origin, priority)...)

	}

	fmt.Println("--FINISH read data dir--")

	sort.SliceStable(req, func(i, j int) bool {
		return req[i].URL < req[j].URL
	})

	fmt.Println("--START save csv--")
	har.SaveRequestAsCSV(req, outputPath)
	fmt.Println("--FINISH save csv--")
}

func readDir(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	list, err := file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	var res []string
	for _, f := range list {
		slice := strings.Split(f.Name(), ".")
		if len(slice) == 2 && slice[1] == "har" {
			res = append(res, f.Name())
		}
	}

	return res
}

func readFile(path string) []byte {
	b, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	return b
}
