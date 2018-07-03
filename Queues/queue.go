//+build disable

package main

type Queue struct{
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	length := len(q.items)
	if length == 0 {
		return -1
	}

	item, items := q.items[0], q.items[1:]
	q.items = items
	return item
}

func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(4)
	queue.Enqueue(8)

	println(queue.Dequeue())
	println(queue.Dequeue())
	println(queue.Dequeue())
	println(queue.Dequeue())
	println(queue.Dequeue())
	println(queue.Dequeue())

}
