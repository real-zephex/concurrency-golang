package scripts

import (
	"api/routes"
	"api/types"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SequentialManager(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	posts, err := routes.PostsHandler()
	if err != nil {
		handleError(err, "Posts Handler", w)
		return
	}

	quotes, err := routes.QuotesHandler()
	if err != nil {
		handleError(err, "Quotes Handler", w)
		return
	}

	todos, err := routes.TodoHandler()
	if err != nil {
		handleError(err, "Todo Handler", w)
		return
	}
	end := time.Since(start)

	w.Header().Set("Content-Type", "application/json")

	combined := types.Combined{Posts: posts, Quotes: quotes, Todos: todos}

	combined.Time = end.String()

	data, err := json.Marshal(combined)
	if err != nil {
		handleError(err, "JSON Marshaller", w)
		return
	}

	fmt.Fprintf(w, "%s", data)

}
