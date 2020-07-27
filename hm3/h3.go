// Реализовать очередь. Это структура данных, работающая по принципу FIFO (First Input — first output, или «первым зашел — первым вышел»).
package main

import "fmt"

type elem struct {
	id   int
	next *elem
	prev *elem
}

var lastElem, firstElem *elem

func push(id int) {
	if lastElem != nil {
		tmpElem := elem{
			id:   id,
			next: lastElem,
			prev: nil,
		}
		lastElem.prev = &tmpElem
		lastElem = &tmpElem
	} else {
		lastElem = &elem{
			id:   id,
			next: nil,
			prev: nil,
		}
		firstElem = lastElem
	}
}

func pop() {
	if firstElem != nil {
		fmt.Println(firstElem.id)
		firstElem = firstElem.prev
		if firstElem == nil {
			lastElem = nil
		}
	} else {
		fmt.Println("Queue is empty")
	}
}
func main() {
	pop()
	pop()
	push(10)
	push(20)
	pop()
	pop()
	pop()
	push(30)
	pop()
	pop()
	pop()
}
