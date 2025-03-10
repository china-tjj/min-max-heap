// Copyright © 2025 tjj
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package min_max_heap

import (
	"math/rand"
	"testing"
)

func TestMinMaxHeap(t *testing.T) {
	h := NewMinMaxHeap[int](func(a, b int) bool { return a < b })
	var arr []int
	for i := 0; i < 100; i++ {
		arr = append(arr, i)
		arr = append(arr, 1000-i)
	}
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	for _, i := range arr {
		h.Push(i)
	}
	for i := 0; i < 100; i++ {
		if h.PopMin() != i {
			t.Fatalf("PopMin() should be %d, but is %d", i, h.PopMin())
		}
		if h.Len() != 200-2*i-1 {
			t.Fatalf("Len() should be %d, but is %d", len(arr)-2*i-1, h.Len())
		}
		if h.PopMax() != 1000-i {
			t.Fatalf("PopMax() should be %d, but is %d", 1000-i, h.PopMax())
		}
		if h.Len() != 200-2*i-2 {
			t.Fatalf("Len() should be %d, but is %d", len(arr)-2*i-2, h.Len())
		}
	}
}
