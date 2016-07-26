package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

const INITIAL_ITEMS_API_URL string = "https://my.cl.ly/v3/items"
const DOWNLOADS_FOLDER string = "~/Downloads"

type Item struct {
	slug         string
	name         string
	created_at   string
	updated_at   string
	long_link    bool
	item_type    string
	view_counter int
	content_url  string
	redirect_url string
	remote_url   string
	source       string
}

type Page struct {
	method string
	url    string
}

type ItemsResponse struct {
	data  []Item
	links []Page
	meta  []string
}

func loadItemsPage(url string) ([]Item, string) {
	resp, err := http.Get(url)
	if err {
		panic(err)
	}

	json_resp := json.Unmarshal(resp, ItemsResponse)

	return json_resp.data, links.next_page
}

func downloadItem(url, fn) {
	// @todo
}

func main() {
	items_url := INITIAL_ITEMS_API_URL

	for {
		items, next_page := loadItemsPage()

		// Download items in parallel
		for range items {
			go downloadItem(item.remote_url, item.name)
		}

		// Load next page
		if !next_page {
			break
		}

		// Update endpoint for next page load
		items_url := next_page
	}
}
