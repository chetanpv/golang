package revision

import "fmt"

func Queues() {
	qpush(1)
	qpush(2)
	qpush(3)
	fmt.Println(queue)
	qpop()
	fmt.Println(queue)
	qpop()
	fmt.Println(queue)
	qpop()
	fmt.Println(queue)
}

var queue []int

func qpush(value int) {
	queue = append(queue, value)
}

func qpop() {
	if len(queue) != 0 {
		queue = queue[1:]
	} else {
		fmt.Println("Queue is empty")
	}

}
