package main

//estrutura das tarefas
type Task struct {
	ID        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Status    string `json:"status"`
}

type TaskStore struct {
	Tasks []Task `json:"tasks"`
}