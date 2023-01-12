package todo_demo

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:required"`
	Password string `json:"password" binding:required`
	Name     string `json:"name" binding:required`
}
