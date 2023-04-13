package main

import (
	"encoding/csv"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	
	file, err := os.Create("data.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	c := colly.NewCollector()

	writer.Write([]string{
		"titulo",
		"figura",
		"descrição",
		"preço",
		"avaliação",
	})

	c.OnHTML(".thumbnail", func(h *colly.HTMLElement) {
		title := h.ChildText(".title")
		figure := h.ChildAttrs(".img-responsive", "src")
		description := h.ChildText(".description")
		price := h.ChildText(".price")
		rating := h.ChildAttr(".ratings > p ~ p", "data-rating")



		writer.Write([]string{
			title,
			figure[0],
			description,
			price,
			rating,
		})
	})

	c.Visit("https://webscraper.io/test-sites/e-commerce/allinone/computers/laptops")

}