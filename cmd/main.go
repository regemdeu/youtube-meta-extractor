package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Usage = "YouTube meta keywords extractor"
	app.Description = "YouTube meta keywords extractor"

	app.Action = func(c *cli.Context) error {
		u := c.Args().Get(0)
		if u == "" {
			return cli.ShowAppHelp(c)
		}

		fmt.Println(ExtrtactKeyWords(u))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func ExtrtactKeyWords(url string) []string {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader((string(html))))
	if err != nil {
		log.Fatal(err)
	}

	r := make([]string, 0, 20)

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {

		if name, _ := s.Attr("name"); name == "keywords" {
			content, _ := s.Attr("content")
			r = append(r, content)
		}
	})

	return r
}
