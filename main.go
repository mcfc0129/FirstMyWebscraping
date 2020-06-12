package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const starturl1 = "https://search.yahoo.co.jp/image/search?p="
const starturl2 = "&op=&ei=UTF-8&b="

func main() {
	urls, word := CreateURL()
	pageurl := GetPage(urls)
	EncodingCSV(pageurl, word)
}

func CreateURL() ([]string, string) {
	var word string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word = s.Text()
		break
	}
	enword := url.QueryEscape(word)

	send := []string{}
	for i := 1; i < 700; i += 20 {
		a := strconv.Itoa(i)
		urls := starturl1 + enword + starturl2 + a
		send = append(send, urls)
	}
	return send, word
}

func GetPage(urls []string) []string {
	send := []string{}
	for _, i := range urls {
		doc, _ := goquery.NewDocument(i)
		doc.Find("img").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("src")
			send = append(send, url)
		})
	}
	return send
}

func EncodingCSV(urls []string, word string) {
	filename := word + ".csv"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("cannot create file:", err)
	}
	defer file.Close()

	record := []string{}
	csvfile := csv.NewWriter(file)
	for i, k := range urls {
		id := strconv.Itoa(i + 1)
		record = []string{
			id,
			k,
		}
		err := csvfile.Write(record)
		if err != nil {
			log.Fatal("cannot write record:", err)
		}
	}
}
