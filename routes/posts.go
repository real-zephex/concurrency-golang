package routes

import (
	"api/types"
	"encoding/json"
	"net/http"
	"time"
)

var client = &http.Client{}

func PostsHandler() (types.ParentPosts, error) {
	start := time.Now()

	// creating a new request
	req, err := http.NewRequest("GET", "https://dummyjson.com/posts", nil)
	req.Header.Set("Cache-Header", "no-cache, no-store, must-revalidate")
	req.Header.Set("Vary", "*")
	if err != nil {
		return types.ParentPosts{}, err
	}

	// performing the request
	res, err := client.Do(req)
	if err != nil {
		return types.ParentPosts{}, err
	}
	defer res.Body.Close()

	var posts types.ParentPosts

	// reading the body and converting it to struct
	err = json.NewDecoder(res.Body).Decode(&posts)
	if err != nil {
		return types.ParentPosts{}, err
	}

	end := time.Since(start)
	posts.Time = end.String()

	// taking the struct and converting it to actual json string
	// marhsal is used when you want to convert map or struct to json string

	// fmt.Fprintln(w, string(data))
	return posts, nil
}
