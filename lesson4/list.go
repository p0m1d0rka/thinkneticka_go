// Реализуация двусвязного списка вместе с базовыми операциями.
package list

import (
	"fmt"
)

// List - двусвязный список.
type List struct {
	root *Elem
}

// Elem - элемент списка.
type Elem struct {
	Val        interface{}
	next, prev *Elem
}

// String реализует интерфейс fmt.Stringer представляя список в виде строки.
func (e *Elem) String() string {
	var s string
	s += fmt.Sprintf("%v ", e.Val)
	return s
}

// New создаёт список и возвращает указатель на него.
func New() *List {
	var l List
	l.root = &Elem{}
	l.root.next = l.root
	l.root.prev = l.root
	return &l
}

// Push вставляет элемент в начало списка.
func (l *List) Push(e Elem) *Elem {
	e.prev = l.root
	e.next = l.root.next
	l.root.next = &e
	if e.next != l.root {
		e.next.prev = &e
	}
	return &e
}

// String реализует интерфейс fmt.Stringer представляя список в виде строки.
func (l *List) String() string {
	el := l.root.next
	var s string
	for el != l.root {
		s += fmt.Sprintf("%v ", el.Val)
		el = el.next
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

// Pop удаляет первый элемент списка.
func (l *List) Pop() *List {
	l.root.next = l.root.next.next
	return l
}

// Reverse разворачивает список.
func (l *List) Reverse() *List {
	el := l.root.next
	root := l.root.next
	// el.next, el.prev = el.prev, el.next
	for i := 0; i < 2; i++ {
		fmt.Println(el)
		el = el.next
	}
	for i := 0; i < 2; i++ {
		fmt.Println(el, root)
		root.next = el
		el = el.prev

		root = root.next
	}
	// fmt.Println(reversed_l)
	return l
}
