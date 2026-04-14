package scripts

import (
	"api/routes"
	"api/types"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func getPosts(ch *types.Combined, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := routes.PostsHandler()
	if err != nil {
		ch.Posts = types.ParentPosts{}
		return
	}
	ch.Posts = res
}

func getQuotes(ch *types.Combined, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := routes.QuotesHandler()
	if err != nil {
		ch.Quotes = types.ParentQuotes{}
		return
	}
	ch.Quotes = res
}

func getTodos(ch *types.Combined, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := routes.TodoHandler()
	if err != nil {
		ch.Todos = types.ParentTodos{}
		return
	}
	ch.Todos = res
}

func ConcurrentManager(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	agent := r.Header.Get("User-Agent")
	if blockCurl(agent) {
		handleError(fmt.Errorf("Not allowed to access"), "Oopsies, you are not allowd to access", w)
		return
	}

	// Run the three fetches concurrently but avoid writing to the same
	// Combined struct from multiple goroutines to prevent data races.
	var wg sync.WaitGroup
	var posts types.ParentPosts
	var quotes types.ParentQuotes
	var todos types.ParentTodos

	wg.Add(3)

	// Each goroutine writes to its own local variable and reports via Done.
	go func() {
		defer wg.Done()
		res, err := routes.PostsHandler()
		if err == nil {
			posts = res
		}
	}()

	go func() {
		defer wg.Done()
		res, err := routes.QuotesHandler()
		if err == nil {
			quotes = res
		}
	}()

	go func() {
		defer wg.Done()
		res, err := routes.TodoHandler()
		if err == nil {
			todos = res
		}
	}()

	wg.Wait()

	end := time.Since(start)

	combined := types.Combined{Posts: posts, Quotes: quotes, Todos: todos}
	combined.Time = end.String()

	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(combined)
	if err != nil {
		handleError(err, "JSON Handling in Concurrent Handler", w)
		return
	}

	fmt.Fprint(w, string(data))
}
