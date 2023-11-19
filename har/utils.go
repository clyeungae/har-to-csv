package har

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func ReadHar(raw []byte) Har {
	result := make(map[string]interface{})
	json.Unmarshal(raw, &result)

	var res Har
	err := json.Unmarshal(raw, &res)

	if err != nil {
		fmt.Println("error when ReadHar")
		fmt.Println(err)
	}

	return res
}

func ReadURL(h Har, origin string, priority string) []string {
	var urls []string
	for _, entry := range h.Log.Entries {
		url := strings.ReplaceAll(entry.Request.URL, origin, "")
		if len(priority) > 0 && entry.Priority != priority {
			continue
		}

		urls = append(urls, url)
	}

	return urls
}

func FilterRequest(h Har, origin string, priority string) []Request {
	var r []Request

	for _, entry := range h.Log.Entries {
		if len(priority) > 0 && entry.Priority != priority {
			continue
		}

		or := entry.Request
		or.URL = strings.ReplaceAll(entry.Request.URL, origin, "")
		r = append(r, or)
	}

	return r
}

func SaveRequestAsCSV(rArr []Request, path string) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var data [][]string
	for _, r := range rArr {
		row := []string{r.Method, r.URL}
		data = append(data, row)
	}

	writer.WriteAll(data)

	return nil
}
