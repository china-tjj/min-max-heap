// Copyright © 2025 tjj
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package min_max_heap

import (
	"math"
	"sync"
	"time"
)

// SyncHeap 并发安全的带有容量限制的优先级队列(最小即为最优)，在超出容量限制后删除最大的元素
type SyncHeap[T any] struct {
	h      *MinMaxHeap[T]
	tokens chan struct{}
	mu     sync.Mutex
	cap    int
}

// NewSyncHeap 创建 SyncHeap，如果 cap 小于 1 会 panic
func NewSyncHeap[T any](less func(a, b T) bool, cap int) *SyncHeap[T] {
	if cap < 1 {
		panic("cap must be greater than or equal to 1")
	}
	return &SyncHeap[T]{
		h:      NewMinMaxHeap[T](less, cap),
		tokens: make(chan struct{}, math.MaxInt),
		cap:    cap,
	}
}

// Push 插入元素
func (q *SyncHeap[T]) Push(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.h.Push(item)
	q.tokens <- struct{}{}
}

func (q *SyncHeap[T]) popSlow() T {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.h.Len() >= q.cap-1 {
		<-q.tokens
		q.h.PopMax()
	}
	return q.h.PopMin()
}

// Pop 弹出最小元素，如果没有元素则会阻塞
func (q *SyncHeap[T]) Pop() T {
	<-q.tokens
	return q.popSlow()
}

// PopWithDone 弹出最小元素，如果没有元素且取消后则停止阻塞并返回 false
func (q *SyncHeap[T]) PopWithDone(done chan struct{}) (T, bool) {
	select {
	case <-q.tokens:
		return q.popSlow(), true
	case <-done:
		var zero T
		return zero, false
	}
}

// PopWithTimeout 弹出最小元素，如果没有元素且超时后则停止阻塞并返回 false
func (q *SyncHeap[T]) PopWithTimeout(timeout time.Duration) (T, bool) {
	select {
	case <-q.tokens:
		return q.popSlow(), true
	case <-time.After(timeout):
		var zero T
		return zero, false
	}
}

// TryPop 弹出最小元素，如果没有元素则不会阻塞并返回false
func (q *SyncHeap[T]) TryPop() (T, bool) {
	select {
	case <-q.tokens:
		return q.popSlow(), true
	default:
		var zero T
		return zero, false
	}
}

// Len 查询元素个数
func (q *SyncHeap[T]) Len() int {
	return len(q.tokens)
}
