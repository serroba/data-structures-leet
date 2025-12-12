package stack

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	type args[T any] struct {
		item T
	}
	type testCase[T any] struct {
		name string
		s    Stack[T]
		args args[T]
	}
	tests := []testCase[int]{
		{name: "Add to stack", s: New[int](), args: args[int]{item: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.item)
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type want[T any] struct {
		val T
		ok  bool
		len int
	}
	type testCase[T any] struct {
		name string
		s    Stack[T]
		want want[T]
	}
	tests := []testCase[int]{
		{name: "pop on empty stack", s: New[int](), want: want[int]{}},
		{name: "pop with 1 value", s: New[int](1), want: want[int]{val: 1, ok: true, len: 0}},
		{name: "pop with many", s: New[int](1, 2, 3, 4, 5), want: want[int]{val: 5, ok: true, len: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.s.Pop()
			if !reflect.DeepEqual(got, tt.want.val) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want.val)
			}
			if ok != tt.want.ok {
				t.Errorf("Pop() ok = %v, want %v", ok, tt.want.ok)
			}
			if tt.s.Len() != tt.want.len {
				t.Errorf("Len() len = %v, want %v", tt.s.Len(), tt.want.len)
			}
		})
	}
}
