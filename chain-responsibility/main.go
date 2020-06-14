package main

import (
	"fmt"
)

type request struct {
	id int
}

type handler interface {
	next(*request)
	process(*request)
}

type handlerB struct {
	nextHandler handler
}

func (h *HandlerA) next(r *request) {
	fmt.Println("next1.....")
	if h.nextHandler != nil {
		h.nextHandler.process(r)
	}
}

func (h *handlerA) process(r *request) {
	fmt.Println("proces1.....")
	h.next(r)
}

func (h *handlerB) next(r *request) {
	fmt.Println("next2.....")
	if h.nextHandler != nil {
		h.nextHandler.process(r)
	}
}

func (h *handlerB) process(r *request) {
	fmt.Println("proces2.....")
	h.next(r)
}

func main() {
	var b *handlerB = &handlerB{nextHandler: nil}
	var a *handlerA = &handlerA{nextHandler: b}
	var r *request = &request{id: 1234}
	a.process(r)
	fmt.Println(r.id)
}
