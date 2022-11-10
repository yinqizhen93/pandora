package utils

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type TaskFunc func(interface{}) error

// Task is a wrapper that contains function and its argument.
type Task struct {
	Run TaskFunc
	Arg interface{}
}

var taskPool = sync.Pool{New: func() interface{} { return new(Task) }}

// GetTask gets a cached Task from pool.
func GetTask() *Task {
	return taskPool.Get().(*Task)
}

// PutTask puts the trashy Task back in pool.
func PutTask(task *Task) {
	task.Run, task.Arg = nil, nil
	taskPool.Put(task)
}

// AsyncTaskQueue is a queue storing asynchronous tasks.
type AsyncTaskQueue interface {
	Enqueue(*Task)
	Dequeue() *Task
	IsEmpty() bool
}

// lock_free_queue 无锁队列

// lockFreeQueue is a simple, fast, and practical non-blocking and concurrent queue with no lock.
type lockFreeQueue struct {
	head   unsafe.Pointer
	tail   unsafe.Pointer
	length int32
}

type node struct {
	value *Task
	next  unsafe.Pointer
}

// NewLockFreeQueue instantiates and returns a lockFreeQueue.
func NewLockFreeQueue() AsyncTaskQueue {
	n := unsafe.Pointer(&node{})
	return &lockFreeQueue{head: n, tail: n}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *lockFreeQueue) Enqueue(task *Task) {
	n := &node{value: task}
retry:
	tail := load(&q.tail)
	next := load(&tail.next)
	// Are tail and next consistent?
	if tail == load(&q.tail) {
		if next == nil {
			// Try to link node at the end of the linked list.
			if cas(&tail.next, next, n) { // enqueue is done.
				// Try to swing tail to the inserted node.
				cas(&q.tail, tail, n)
				atomic.AddInt32(&q.length, 1)
				return
			}
		} else { // tail was not pointing to the last node
			// Try to swing tail to the next node.
			cas(&q.tail, tail, next)
		}
	}
	goto retry
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *lockFreeQueue) Dequeue() *Task {
retry:
	head := load(&q.head)
	tail := load(&q.tail)
	next := load(&head.next)
	// Are head, tail, and next consistent?
	if head == load(&q.head) {
		// Is queue empty or tail falling behind?
		if head == tail {
			// Is queue empty?
			if next == nil {
				return nil
			}
			cas(&q.tail, tail, next) // tail is falling behind, try to advance it.
		} else {
			// Read value before CAS, otherwise another dequeue might free the next node.
			task := next.value
			if cas(&q.head, head, next) { // dequeue is done, return value.
				atomic.AddInt32(&q.length, -1)
				return task
			}
		}
	}
	goto retry
}

// IsEmpty indicates whether this queue is empty or not.
func (q *lockFreeQueue) IsEmpty() bool {
	return atomic.LoadInt32(&q.length) == 0
}

func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *node) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
