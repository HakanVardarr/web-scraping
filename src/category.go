package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type category struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Products products
}

func (c *category) View(m model) string {
	var content string

	if c == nil {
		content += fmt.Sprintf("%s\n\n", defaultText.Render("Choose your category"))

		for i, category := range m.categories {
			if m.cursor == i {
				content += fmt.Sprintf("> %s\n", chosenCategory.Render(category.Name))
			} else {
				content += fmt.Sprintf("%s\n", defaultText.Render(category.Name))
			}
		}
	} else {
		content += fmt.Sprintf("%s%s\n%s%s\n\n", defaultText.Render("Listing: "), chosenCategory.Render(m.category.Name), defaultText.Render("Page: "), chosenCategory.Bold(true).Render(fmt.Sprint(m.page)))
		content += c.Products.View(m)

	}

	content += fmt.Sprintf("\n%s", defaultText.Render("Press Q to quit!"))

	return container.Render(content)
}

func categories() ([]category, error) {
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
