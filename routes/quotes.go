package routes

import (
	"api/types"
	"encoding/json"
	"net/http"
	"time"
)

func QuotesHandler() (types.ParentQuotes, error) {
	start := time.Now()

	// creating a new request
	req, err := http.NewRequest("GET", "https://dummyjson.com/quotes", nil)
	req.Header.Set("Cache-Header", "no-cache, no-store, must-revalidate")
	req.Header.Set("Vary", "*")
	if err != nil {
		return types.ParentQuotes{}, err
	}

	// performing the request
	res, err := client.Do(req)
	if err != nil {
		return types.ParentQuotes{}, err
	}
	defer res.Body.Close()

	var quotes types.ParentQuotes

	// reading the body and converting it to struct
	err = json.NewDecoder(res.Body).Decode(&quotes)
	if err != nil {
		return types.ParentQuotes{}, err
	}

	end := time.Since(start)
	quotes.Time = end.String()

	return quotes, nil
}
