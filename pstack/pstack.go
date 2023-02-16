package pstack

import (
	"errors"
	"fmt"
)

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{value: value}
}

type PersistentStack[T any] struct {
	bottom *node[T]
	top    *node[T]

	memory map[string]*node[T]
}

func NewPersistentStack[T any]() PersistentStack[T] {
	return PersistentStack[T]{memory: map[string]*node[T]{}}
}

func (p *PersistentStack[T]) Push(value T) {
	node := newNode(value)

	if p.top == nil {
		p.bottom = node
		p.top = node
		return
	}

	p.top.next = node
	node.prev = p.top
	p.top = node
}

func (p *PersistentStack[T]) Pop() (T, error) {
	if p.top != nil {
		value := p.top.value
		p.top = p.top.prev
		return value, nil
	}

	var none T
	return none, errors.New("Stack empty")
}

func (p *PersistentStack[T]) popNode() (*node[T], error) {
	if p.top != nil {
		value := p.top
		p.top = p.top.prev
		return value, nil
	}

	return nil, errors.New("Stack empty")
}

func (p *PersistentStack[T]) pushNode(node *node[T]) {
	if node == nil {
		return
	}

	if p.top == nil {
		p.bottom = node
		p.top = node
		return
	}

	p.top.next = node
	node.prev = p.top
	p.top = node
}

func (p *PersistentStack[T]) PopAll() (items []T) {
	for item, err := p.Pop(); err == nil; item, err = p.Pop() {
		items = append(items, item)
	}
	return
}

// Adds a remember point at the top of the stack.
func (p *PersistentStack[T]) RememberPoint(label string) error {
	if p.top != nil {
		p.memory[label] = p.top
		return nil
	}
	return errors.New("Cannot set remember point in an empty stack")
}

// Stops popping just before the remember point. Does not pop the items added before the remember point
func (p *PersistentStack[T]) PopTill(label string) (items []T, err error) {
	stopPoint, ok := p.memory[label]
	if !ok {
		err = errors.New(fmt.Sprintf("Remember point %s not found", label))
		return
	}

	stopped := false
	for itemNode, err := p.popNode(); err == nil; itemNode, err = p.popNode() {
		if itemNode == stopPoint {
			stopped = true
			p.pushNode(itemNode)
			break
		}
		items = append(items, itemNode.value)
	}

	if !stopped {
		err = errors.New("Remember point was never reached while popping")
	} else {
		err = nil
	}
	return
}

// Only peeks all the elements till the remember point and returns them
func (p *PersistentStack[T]) PeekTill(label string) (items []T, err error) {
	stopPoint, ok := p.memory[label]
	if !ok {
		err = errors.New(fmt.Sprintf("Remember point %s not found", label))
		return
	}

	stopped := false
	for node := p.top; node != nil; node = node.prev {
		if node == stopPoint {
			stopped = true
			break
		}
		items = append(items, node.value)
	}

	if !stopped {
		err = errors.New("Remember point was nver reached while popping")
	}
	return
}
