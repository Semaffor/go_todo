package todo_demo

type TodoList struct {
	Id          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int32
	UserId int32
	ListId int32
}

type TodoItem struct {
	Id          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

type ListsItem struct {
	Id     int32
	ListId int32
	ItemId int32
}
