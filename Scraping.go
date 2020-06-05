package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type field struct {
	id  int
	url string
}

func main() {
	serve := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", SearchFrom)
	http.HandleFunc("/ok", okPage)

	serve.ListenAndServe()
}

func SearchFrom(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/Search.html")
	t.Execute(w, nil)
}

func okPage(w http.ResponseWriter, r *http.Request) {
	Searches := r.FormValue("Search")
	file := r.FormValue("filename")
	filename := file + ".csv"
	URLencode := url.QueryEscape(Searches)
	url := "https://search.yahoo.co.jp/image/search?p=" + URLencode + "&op=&ei=UTF-8&b="

	a := fordeal(url)
	data := GetPage(a)
	EncodingCSV(data, filename)

	t, _ := template.ParseFiles("html/ok.html")
	t.Execute(w, nil)
}

func fordeal(url string) []string {
	result := []string{}
	for i := 1; i < 1001; i += 20 {
		query := strconv.Itoa(i)
		a := url + query
		result = append(result, a)
	}
	return result
}

func GetPage(url []string) []field {
	result := []field{}
	i := 1
	for _, l := range url {
		doc, _ := goquery.NewDocument(l)
		doc.Find("img").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("src")
			f := field{
				id:  i,
				url: url,
			}
			result = append(result, f)
			i += 1
		})
	}
	return result
}

func EncodingCSV(data []field, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Create fileに失敗:", err)
		return
	}
	defer file.Close()
	write := csv.NewWriter(file)
	for _, k := range data {
		row := []string{strconv.Itoa(k.id), k.url}
		err := write.Write(row)
		if err != nil {
			log.Fatal("can't encoding csv file: ", err)
			return
		}
	}
	write.Flush()
}
