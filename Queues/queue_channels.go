package main

// Queue structure but with channels
type Queue struct {
	items chan int
}

// Enqueue an item into queue of channels
func (q *Queue) Enqueue(item int)  {
	q.items <- item // write to a channel
}

// Dequeue an item from a queue of channels
func (q *Queue) Dequeue() int {
	return <- q.items // read from a channel
}

func main(){
	q := Queue{
		items: make (chan int, 16),
	}

	q.Enqueue(1)
	q.Enqueue(2)

	println(q.Dequeue())
	println(q.Dequeue())
	// println(q.Dequeue()) // Will fix later

}