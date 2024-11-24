package main

import (
	"errors"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) edit(command string) {

}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		error := errors.New("Index is not present")
		return error
	} else {
		return nil
	}
}

func (todos *Todos) del(index int) {

}

func (todos *Todos) list() string {
	return "foobar"
}
