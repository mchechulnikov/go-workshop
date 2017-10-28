package linkedList

import (
	"fmt"
)

func New() *LinkedList {
	return &LinkedList{}
}

type LinkedList struct {
	First  *Element
	Last   *Element
	Length int
}

func (l *LinkedList) Add(value interface{}) {
	element := l.First
	if element == nil {
		l.First = &Element{value: value}
		l.Last = l.First
		l.Length++
		return
	}

	l.Last.next = &Element{value: value}
	l.Last = l.Last.next
	l.Length++
}

func (l *LinkedList) AddByIndex(value interface{}, index int) error {
	element := l.First
	if l.Length < index {
		return fmt.Errorf("Index %d out of range: max length of linked list equal %d", index, l.Length)
	}

	var prevElement *Element
	counter := 0
	for {
		if counter == index {
			if prevElement == nil {
				l.First = &Element{value: value, next: element}
			} else if element == nil {
				element = &Element{value: value}
				prevElement.next = element
				l.Last = element
			} else {
				prevElement.next = &Element{value: value, next: element}
			}
			l.Length++
			return nil
		}
		prevElement = element
		element = element.next
		counter++
	}
}

func (l *LinkedList) Get(index int) (*Element, error) {
	element := l.First
	if element == nil {
		return nil, fmt.Errorf("Couldn't get element from empty linked list")
	}

	if l.Length < index {
		return nil, fmt.Errorf("Index %d out of range: max length of linked list equal %d", index, l.Length)
	}

	counter := 0
	for {
		if counter == index {
			return element, nil
		}
		if element.next == nil {
			return nil, fmt.Errorf("Index %d out of range: max length of linked list equal %d", index, counter)
		}
		element = element.next
		counter++
	}
}

func (l *LinkedList) Delete(index int) error {
	element := l.First
	if element == nil {
		return fmt.Errorf("Couldn't delete any element from empty linked list")
	}

	if l.Length < index {
		return fmt.Errorf("Index %d out of range: max length of linked list equal %d", index, l.Length)
	}

	var prevElement *Element
	counter := 0
	for {
		if counter == index {
			if prevElement == nil {
				l.First = l.First.next
			} else if element == nil {
				prevElement.next = nil
				l.Last = prevElement
			} else {
				prevElement.next = element.next
			}
			l.Length--
			return nil
		}
		prevElement = element
		element = element.next
		counter++
	}
}

func (l* LinkedList) HasCycle() bool {
	if l.First == nil {
		return false
	}

	// Floyd's algorithm
	tortoise := l.First
	hare := l.First
	for tortoise == hare {
		if hare == nil || hare.next == nil {
			return false
		}
		tortoise = tortoise.next
		hare = hare.next.next
	}
	return true
}

type Element struct {
	value interface{}
	next  *Element
}

func (e *Element) GetValue() interface{} {
	return e.value
}

func (e *Element) GetNext() *Element {
	return e.next
}
