package dataStructures

import (
	"fmt"
)

func New() *LinkedList {
	return &LinkedList{}
}

type LinkedList struct {
	First  *Element
	Length int
}

func (l *LinkedList) Add(value interface{}) {
	element := l.First
	if element == nil {
		l.First = &Element{value: value}
		l.Length++
		return
	}

	for {
		if element.next == nil {
			element.next = &Element{value: value}
			l.Length++
			return
		}
		element = element.next
	}
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

	if l.Length < index+1 {
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
