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

func ConcurrentManager(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	agent := r.Header.Get("User-Agent")
	if blockCurl(agent) {
		handleError(fmt.Errorf("Not allowed to access"), "Oopsies, you are not allowd to access", w)
		return
	}

	var wg sync.WaitGroup
	var posts types.ParentPosts
	var quotes types.ParentQuotes
	var todos types.ParentTodos

	wg.Add(3)

	/*
		Each function will assign the result to its local variable only.
		We will be combining all of them into one combined struct later.
	*/
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
