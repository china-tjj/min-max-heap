# min-max-heap
go语言实现的泛型最小最大堆
```go
type MinMaxHeap[T any] struct
```

## 提供了以下的方法

|   方法    |   描述   |  时间复杂度   |
|:-------:|:------:|:--------:|
|  Push   |  插入元素  | O(log n) |
| PopMin  | 弹出最小元素 | O(log n) |
| PopMax  | 弹出最大元素 | O(log n) |
| PeekMin | 查询最小元素 |   O(1)   |
| PeekMin | 查询最大元素 |   O(1)   |
|   Len   | 查询元素个数 |   O(1)   |

## 适用场景
带有容量限制的优先级队列(最小即为最优)，在超出容量限制后删除最大的元素，本项目基于`MinMaxHeap`简单的实现了并发安全的上述优先级队列`SyncHeap`