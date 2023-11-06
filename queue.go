package main

type BoundedQueue struct {
	queue chan string
}

func (bq *BoundedQueue) init(capacity int) {
	bq.queue = make(chan string, capacity)
}

func (bq *BoundedQueue) enqueue(item string) int {
	bq.queue <- item
	return 0
}

func (bq *BoundedQueue) dequeue() (string, int) {
	item, err := <-bq.queue
	if err {
		return item, 0
	} else {
		// fmt.Println(err)
		return item, -1
	}
}

func (bq *BoundedQueue) size() int {
	return len(bq.queue)
}

func (bq *BoundedQueue) capacity() int {
	return cap(bq.queue)
}

func (bq *BoundedQueue) close() int {
	close(bq.queue)
	return 0
}
