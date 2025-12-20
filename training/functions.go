package training

import (
	"strconv"
	"strings"
)

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		_, ok := seen[complement]
		if ok {
			return []int{seen[complement], i}
		}

		seen[num] = i
	}

	return []int{}
}

type List struct {
	head *ListNode
	tail *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(numbers ...int) *ListNode {
	l := &ListNode{}
	for i, number := range numbers {
		l.Val = number
		if i < len(numbers)-1 {
			l.Next = &ListNode{}
		}
	}
	return l
}

// func NewListFromNumber(number int) List {
//	list := List{}
//	for number > 0 {
//		list.head = ListNode{}
//	}
//	return
//}

func (l *ListNode) String() string {
	if l == nil {
		return "[]"
	}

	var vals []string

	for l != nil {
		vals = append(vals, strconv.Itoa(l.Val))
		l = l.Next
	}

	return "[" + strings.Join(vals, ", ") + "]"
}

func (l *ListNode) Append(number int) {
	current := l
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{Val: number}
}

func (l *ListNode) AppendMany(numbers ...int) *ListNode {
	for _, number := range numbers {
		l.Append(number)
	}
	return l
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	result := &ListNode{}
	head := result

	var carry int
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		result.Val = sum % 10
		if l1 != nil || l2 != nil || carry > 0 {
			result.Next = &ListNode{}
			result = result.Next
		}
	}

	return head
}

func LengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	left := 0
	maxLen := 0

	for right, char := range s {
		if i, ok := seen[char]; ok && i >= left {
			left = i + 1
		}

		seen[char] = right

		windowLen := right - left + 1
		if windowLen > maxLen {
			maxLen = windowLen
		}
	}

	return maxLen
}

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	magnitude := 1
	for x/magnitude >= 10 {
		magnitude *= 10
	}

	for x > 0 {
		l := x / magnitude

		r := x % 10
		if l != r {
			return false
		}

		x = (x - l*magnitude) / 10
		magnitude /= 100
	}

	return true
}

func RomanToInt(s string) int {
	numbers := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0

	for i := range s {
		current := numbers[s[i]]
		if i+1 < len(s) && current < numbers[s[i+1]] {
			result -= current
		} else {
			result += current
		}
	}

	return result
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var prefix []byte
	//{name: "example 1", args: args{strs: []string{"flower", "flow", "flight"}}, want: "fl"},
	for i, v := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return string(prefix)
			}
		}
		prefix = append(prefix, byte(v))
	}

	return string(prefix)
}

func IsValidParens(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var stack []rune
	for _, paren := range s {
		if isLeftParen(paren) {
			if len(stack) == 0 {
				stack = append(stack, paren)
			} else {
				top := stack[len(stack)-1]
				if isLeftParen(top) {
					stack = append(stack, paren)
				}
			}
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if !matchType(top, paren) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func isLeftParen(r rune) bool {
	return r == '(' || r == '[' || r == '{'
}

func matchType(left, right rune) bool {
	return left == '(' && right == ')' || left == '[' && right == ']' || left == '{' && right == '}'
}

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
