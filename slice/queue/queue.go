package queue

import "fmt"

type Queue struct {
	storage []interface{}
}

func (q *Queue) String() string {
	return fmt.Sprintf("Q:%v", q.storage)
}

func (q *Queue) Clear() {
	q.storage = make([]interface{}, 0, 10)
}

func (q *Queue) Size() int {
	return len(q.storage)
}

func (q *Queue) Add(items ...interface{}) {
	for _, item := range items {
		q.storage = append(q.storage, item)
	}
}

func (q *Queue) Next() interface{} {
	if len(q.storage) == 0 {
		return nil
	}
	result := q.storage[0]
	q.storage = q.storage[1:]
	return result
}
