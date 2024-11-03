package todo

import "time"

/*
*
BOOL = Bulyon
*/
type bulyon bool

type Todo struct {
	Completed bulyon
	CreatedAt time.Time
	Title     string
	Gorev     string
}

type TodoList struct {
	Todos []Todo
}
