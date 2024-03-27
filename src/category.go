package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type category struct {
	Name string
	Link string
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
