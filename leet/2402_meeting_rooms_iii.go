package leet

import (
	"container/heap"
	"sort"
)

func mostBooked(n int, meetings [][]int) int {
	// Sort meetings by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	// Min-heap of available rooms (by room number)
	available := &roomHeap{}
	for i := 0; i < n; i++ {
		heap.Push(available, i)
	}

	// Min-heap of busy rooms: (endTime, roomNumber)
	busy := &busyHeap{}

	// Count of meetings per room
	count := make([]int, n)

	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]
		duration := end - start

		// Free up rooms that have finished by this meeting's start time
		for busy.Len() > 0 && (*busy)[0].endTime <= start {
			room := heap.Pop(busy).(busyRoom)
			heap.Push(available, room.roomNum)
		}

		var roomNum int
		var actualEnd int

		if available.Len() > 0 {
			// Use the available room with lowest number
			roomNum = heap.Pop(available).(int)
			actualEnd = end
		} else {
			// Wait for the next room to be free
			nextFree := heap.Pop(busy).(busyRoom)
			roomNum = nextFree.roomNum
			actualEnd = nextFree.endTime + duration
		}

		count[roomNum]++
		heap.Push(busy, busyRoom{endTime: actualEnd, roomNum: roomNum})
	}

	// Find room with most meetings (lowest number on tie)
	maxRoom := 0
	for i := 1; i < n; i++ {
		if count[i] > count[maxRoom] {
			maxRoom = i
		}
	}

	return maxRoom
}

// roomHeap is a min-heap of room numbers
type roomHeap []int

func (h roomHeap) Len() int           { return len(h) }
func (h roomHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h roomHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *roomHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *roomHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// busyRoom represents a room that is currently in use
type busyRoom struct {
	endTime int
	roomNum int
}

// busyHeap is a min-heap of busy rooms, ordered by end time then room number
type busyHeap []busyRoom

func (h busyHeap) Len() int { return len(h) }
func (h busyHeap) Less(i, j int) bool {
	if h[i].endTime != h[j].endTime {
		return h[i].endTime < h[j].endTime
	}
	return h[i].roomNum < h[j].roomNum
}
func (h busyHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *busyHeap) Push(x any) {
	*h = append(*h, x.(busyRoom))
}

func (h *busyHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}