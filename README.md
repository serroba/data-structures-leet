# ds

Personal data structures and algorithms practice in Go.

## Structure

```
ds/
├── graph/      # Graph representations and algorithms
├── list/       # Linked list implementations
├── queue/      # Queue data structure
├── stack/      # Stack data structure
├── tree/       # Binary tree utilities
├── trie/       # Trie with wildcard pattern matching
└── training/   # LeetCode problem solutions
```

## Packages

### graph
Graph data structure with adjacency list representation.

### list
Singly linked list implementation.

### queue
Generic queue with enqueue/dequeue operations.

### stack
Generic stack with push/pop operations.

### tree
Binary tree node definition and utilities.

### trie
Trie (prefix tree) with support for:
- Word insertion and search
- Prefix matching
- Wildcard patterns (`?` and `*`)

## Training

LeetCode solutions organized by problem number:
- Tree traversals (94, 102, 103, 144, 145)
- Binary tree properties (100, 101, 104, 110, 111, 112)
- Tree construction (105, 108, 109)
- Arrays and strings (121, 125, 136)
- Linked lists (141)
- And more...

## Run Tests

```bash
go test ./...
```
