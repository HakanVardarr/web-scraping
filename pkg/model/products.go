package model

import (
	"fmt"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

type product struct {
	name, price, link string
}

type products struct {
	items []product
}

func (p products) IsEmpty() bool {
	return len(p.items) == 0
}

func (p products) Length() int {
	return len(p.items)
}

func (p *products) Clear() {
	p.items = make([]product, 0)
}

func (p *products) GetItems(m model) {
	resp, err := soup.Get(m.category.Link + fmt.Sprint(m.page))
	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)
	for _, pr := range doc.FindAll("div", "class", "p-card-wrppr") {
		name := pr.Find("span", "class", "prdct-desc-cntnr-name").Text()
		price := pr.Find("div", "class", "prc-box-dscntd").Text()
		link := pr.Find("a").Attrs()["href"]

		p.items = append(p.items, product{name, price, link})

	}

}

func (p products) Draw(m model, content *strings.Builder) {
	for i, pr := range p.items {
		if m.cursor == i {
			content.WriteString(fmt.Sprintf("> %s %s\n", chosenCategory.Render(pr.name), price.Render(pr.price)))
		} else {
			content.WriteString(fmt.Sprintf("%s %s\n", defaultText.Render(pr.name), defaultText.Render(pr.price)))
		}
	}

}

func (p *products) View(m model, content *strings.Builder) {
	if p.IsEmpty() {
		p.GetItems(m)
	}

	p.Draw(m, content)
}
