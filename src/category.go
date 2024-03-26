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

func categories() []category {
	bytes, err := os.ReadFile("data/categories.json")

	if err != nil {
		fmt.Println("categories: unable to load config file.")
		os.Exit(1)
	}

	var categories []category

	err = json.Unmarshal(bytes, &categories)

	if err != nil {
		fmt.Println("JSON decode error!")
		os.Exit(1)
	}

	return categories
}
