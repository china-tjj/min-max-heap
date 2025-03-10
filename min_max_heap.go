// Copyright © 2025 tjj
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package min_max_heap

/*
为了方便，堆中逻辑索引全部从1开始
*/
// left
func l(i int) int {
	return i << 1
}

// right
func r(i int) int {
	return (i << 1) + 1
}

// parent
func p(i int) int {
	return i >> 1
}

// grandparent
func g(i int) int {
	return i >> 2
}

func isMinLevel(i int) bool {
	// 索引的二进制最高位代表层数，奇数层为min层
	level := 0
	for i > 0 {
		i >>= 1
		level++
	}
	return level%2 == 1
}

type MinMaxHeap[T any] struct {
	arr  []T
	less func(a, b T) bool
}

// NewMinMaxHeap 创建 MinMaxHeap，可以额外传入参数 initCap 指定切片的初始长度
func NewMinMaxHeap[T any](less func(a, b T) bool, initCap ...int) *MinMaxHeap[T] {
	c := 0
	if len(initCap) == 1 {
		c = initCap[0]
	}
	return &MinMaxHeap[T]{
		arr:  make([]T, 0, c),
		less: less,
	}
}

// Push 插入元素，时间复杂度为 O(log n)，其中 n = h.Len()
func (h *MinMaxHeap[T]) Push(value T) {
	h.arr = append(h.arr, value)
	h.pushUp(len(h.arr))
}

// PopMin 弹出最小元素，时间复杂度为 O(log n)，其中 n = h.Len()
func (h *MinMaxHeap[T]) PopMin() T {
	minIdx := 1
	minValue := h.pop(minIdx)
	if minIdx <= len(h.arr) {
		h.pushDownMin(minIdx)
	}
	return minValue
}

// PeekMin 查看最小元素，时间复杂度为 O(1)
func (h *MinMaxHeap[T]) PeekMin() T {
	return h.get(1)
}

// PopMax 弹出最大元素，时间复杂度为 O(log n)，其中 n = h.Len()
func (h *MinMaxHeap[T]) PopMax() T {
	maxIdx := h.getMaxIdx()
	maxValue := h.pop(maxIdx)
	if maxIdx <= len(h.arr) {
		h.pushDownMax(maxIdx)
	}
	return maxValue
}

// PeekMax 查看最大元素，时间复杂度为 O(1)
func (h *MinMaxHeap[T]) PeekMax() T {
	return h.get(h.getMaxIdx())
}

// Len 查询元素个数，时间复杂度为 O(1)
func (h *MinMaxHeap[T]) Len() int {
	return len(h.arr)
}

func (h *MinMaxHeap[T]) get(i int) T {
	return h.arr[i-1]
}

func (h *MinMaxHeap[T]) swap(i, j int) {
	h.arr[i-1], h.arr[j-1] = h.arr[j-1], h.arr[i-1]
}

func (h *MinMaxHeap[T]) pop(i int) T {
	popped := h.arr[i-1]
	last := len(h.arr) - 1
	h.arr[i-1] = h.arr[last]
	h.arr = h.arr[:last]
	return popped
}

func (h *MinMaxHeap[T]) getMaxIdx() int {
	if len(h.arr) <= 2 {
		return len(h.arr)
	} else if h.less(h.get(2), h.get(3)) {
		return 3
	} else {
		return 2
	}
}

func (h *MinMaxHeap[T]) pushUp(i int) {
	parent := p(i)
	if parent <= 0 {
		return
	}
	if isMinLevel(i) {
		if h.less(h.get(parent), h.get(i)) {
			h.swap(parent, i)
			h.pushUpMax(parent)
		} else {
			h.pushUpMin(i)
		}
	} else {
		if h.less(h.get(i), h.get(parent)) {
			h.swap(i, parent)
			h.pushUpMin(parent)
		} else {
			h.pushUpMax(i)
		}
	}
}

func (h *MinMaxHeap[T]) pushUpMin(i int) {
	for {
		grandparent := g(i)
		if grandparent <= 0 || !h.less(h.get(i), h.get(grandparent)) {
			break
		}
		h.swap(i, grandparent)
		i = grandparent
	}
}

func (h *MinMaxHeap[T]) pushUpMax(i int) {
	for {
		grandparent := g(i)
		if grandparent <= 0 || !h.less(h.get(grandparent), h.get(i)) {
			break
		}
		h.swap(i, grandparent)
		i = grandparent
	}
}

func (h *MinMaxHeap[T]) pushDownMin(i int) {
	for {
		minIdx := i
		candidates := [6]int{l(i), r(i), l(l(i)), r(l(i)), l(r(i)), r(r(i))}
		for _, candidate := range candidates {
			if candidate > len(h.arr) {
				break
			}
			if h.less(h.get(candidate), h.get(minIdx)) {
				minIdx = candidate
			}
		}
		if minIdx == i {
			return
		}
		h.swap(minIdx, i)
		if p(minIdx) == i {
			return
		}
		// minIdx为孙节点，在min层
		if h.less(h.get(p(minIdx)), h.get(minIdx)) {
			h.swap(p(minIdx), minIdx)
		}
		i = minIdx
	}
}

func (h *MinMaxHeap[T]) pushDownMax(i int) {
	for {
		maxIdx := i
		candidates := [6]int{l(i), r(i), l(l(i)), r(l(i)), l(r(i)), r(r(i))}
		for _, candidate := range candidates {
			if candidate > len(h.arr) {
				break
			}
			if h.less(h.get(maxIdx), h.get(candidate)) {
				maxIdx = candidate
			}
		}
		if maxIdx == i {
			return
		}
		h.swap(maxIdx, i)
		if p(maxIdx) == i {
			return
		}
		// maxIdx为孙节点，在max层
		if h.less(h.get(maxIdx), h.get(p(maxIdx))) {
			h.swap(maxIdx, p(maxIdx))
		}
		i = maxIdx
	}
}
