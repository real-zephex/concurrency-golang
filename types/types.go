package types

type ParentPosts struct {
	Posts []PostsItems `json:"posts"`
	Time  string       `json:"time"`
}

type PostsItems struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Reactions struct {
		Likes    int `json:"likes"`
		Dislikes int `json:"dislikes"`
	} `json:"reactions"`
	Views  int `json:"views"`
	UserID int `json:"userId"`
}

type ParentQuotes struct {
	Quotes []Quotes `json:"quotes"`
	Time   string   `json:"time"`
}

type Quotes struct {
	ID     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type ParentTodos struct {
	Todos []Todos `json:"todos"`
	Time  string  `json:"time"`
}

type Todos struct {
	ID        int    `json:"id"`
	ToDo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserID    int    `json:"userId"`
}

type Combined struct {
	Time   string       `json:"duration"`
	Posts  ParentPosts  `json:"posts"`
	Quotes ParentQuotes `json:"quotes"`
	Todos  ParentTodos  `json:"todos"`
}
