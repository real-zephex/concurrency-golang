package routes

import (
	"api/types"
	"encoding/json"
	"net/http"
	"time"
)

func TodoHandler() (types.ParentTodos, error) {
	start := time.Now()

	// creating a new request
	req, err := http.NewRequest("GET", "https://dummyjson.com/todos", nil)
	if err != nil {
		return types.ParentTodos{}, err
	}

	// performing the request
	res, err := client.Do(req)
	if err != nil {
		return types.ParentTodos{}, err
	}
	defer res.Body.Close()

	var todos types.ParentTodos

	// reading the body and converting it to struct
	err = json.NewDecoder(res.Body).Decode(&todos)
	if err != nil {
		return types.ParentTodos{}, err
	}

	end := time.Since(start)
	todos.Time = end.String()

	return todos, nil 
}
