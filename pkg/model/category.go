package model

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type category struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	products products
}

func (c *category) View(m model) string {
	var content strings.Builder

	if c == nil {
		content.WriteString(fmt.Sprintf("%s\n\n", defaultText.Render("Choose your category")))

		for i, category := range m.categories {
			if m.cursor == i {
				content.WriteString(fmt.Sprintf("> %s\n", chosenCategory.Render(category.Name)))
			} else {
				content.WriteString(fmt.Sprintf("%s\n", defaultText.Render(category.Name)))
			}
		}

	} else {
		content.WriteString(fmt.Sprintf("%s%s\n%s%s\n\n", defaultText.Render("Listing: "), chosenCategory.Render(m.category.Name), defaultText.Render("Page: "), chosenCategory.Bold(true).Render(fmt.Sprint(m.page))))
		c.products.View(m, &content)

	}

	content.WriteString(fmt.Sprintf("\n%s", defaultText.Render("Press Q to quit!")))

	return container.Render(content.String())
}

func Categories() ([]category, error) {
	bytes, err := os.ReadFile("data/categories.json")

	if err != nil {
		return nil, fmt.Errorf("categories: unable to read JSON file: %w", err)
	}

	var categories []category
	err = json.Unmarshal(bytes, &categories)

	if err != nil {
		return nil, fmt.Errorf("categories: unable to read JSON file: %w", err)
	}

	return categories, nil
}
