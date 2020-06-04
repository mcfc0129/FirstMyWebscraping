package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type feild struct {
	id  int
	url string
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

func GetPage(url []string) []feild {
	result := []feild{}
	i := 1
	for _, l := range url {
		doc, _ := goquery.NewDocument(l)
		doc.Find("img").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("src")
			f := feild{
				id:  i,
				url: url,
			}
			result = append(result, f)
			i += 1
		})
	}
	return result
}

func EncodingCSV(data []feild, filename string) {
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

func main() {
	url := "https://search.yahoo.co.jp/image/search?p=%E3%82%B7%E3%83%9C%E3%83%AC%E3%83%BC%E3%82%AB%E3%83%9E%E3%83%AD&oq=%E3%82%B7%E3%83%9C%E3%83%AC%E3%83%BC&ei=UTF-8&b="
	a := fordeal(url)
	// fmt.Println(a)

	b := GetPage(a)
	// fmt.Println(b)
	EncodingCSV(b, "Camaro.csv")
}
