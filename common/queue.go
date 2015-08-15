package common

import "container/list"

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

func (q *Queue) Enqueue(obj interface{}) {
	q.list.PushBack(obj)
}

func (q *Queue) Dequeue() interface{} {
	element := q.list.Front()
	if element == nil {
		return nil
	}
	value := element.Value
	q.list.Remove(element)
	return value
}

func (q *Queue) Len() int {
	return q.list.Len()
}
