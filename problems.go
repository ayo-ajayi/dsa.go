// moderate
package main

import (
	"math"
	"github.com/ayo-ajayi/dsa.go/stack"

)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	current := &ListNode{}
	dummy := current

	for l1 != nil || l2 != nil || carry != 0 {

		var x int
		var y int
		var sum int
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum = x + y + carry
		carry = sum / 10
		node := &ListNode{Val: sum % 10}
		current.Next = node
		current = current.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	return dummy.Next
}
func longestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}
	for i := len(s); i > 0; i-- {
		for j := 0; j <= len(s)-i; j++ {
			if isPalindrome(s[j : j+i]) {
				return s[j : j+i]
			}
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	for i, j := len(s)-1, 0; i > j; i, j = i-1, j+1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func reverse(x int) int {
	r := 0
	for x != 0 {
		r = r*10 + x%10
		if r > 2147483647 || r < -2147483648 {
			return 0
		}
		x = x / 10
	}

	return r
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1
	res := 0
	i := 0

	for i < len(s) && s[i] == ' ' {
		i++

	}
	if i < len(s) {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}
	for i < len(s) {
		if s[i] < '0' || s[i] > '9' {
			return sign * res
		}
		digit := int(s[i] - '0')
		if res > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		res = res*10 + digit
		i++
	}
	res = sign * res

	return res
}

func maxArea(height []int) int {
	n := len(height)
	if n > 1e5 || n < 2 {
		return 0
	}
	maxArea := 0
	for j, i := n-1, 0; i < j; {
		area := 0
		if height[i] > height[j] {
			area = (j - i) * height[j]
			j--
		} else {
			area = (j - i) * height[i]
			i++
		}
		if area > maxArea {
			maxArea = area
		}

	}

	return maxArea
}

func intToRoman(num int) string {
	return getRoman(num, "")
}

func getRoman(num int, roman string) string {
	if num > 4000 {
		return roman
	}
	if num == 0 {
		return roman
	}
	if num-1000 >= 0 {
		return getRoman(num-1000, roman+"M")
	}
	if num-900 >= 0 {
		return getRoman(num-900, roman+"CM")
	}
	if num-500 >= 0 {
		return getRoman(num-500, roman+"D")
	}
	if num-400 >= 0 {
		return getRoman(num-400, roman+"CD")
	}
	if num-100 >= 0 {
		return getRoman(num-100, roman+"C")
	}
	if num-90 >= 0 {
		return getRoman(num-90, roman+"XC")
	}
	if num-50 >= 0 {
		return getRoman(num-50, roman+"L")
	}
	if num-40 >= 0 {
		return getRoman(num-40, roman+"XL")
	}
	if num-10 >= 0 {
		return getRoman(num-10, roman+"X")
	}
	if num-9 >= 0 {
		return getRoman(num-9, roman+"IX")
	}

	if num-5 >= 0 {
		return getRoman(num-5, roman+"V")
	}

	if num-4 >= 0 {
		return getRoman(num-4, roman+"IV")
	}

	if num-1 >= 0 {
		return getRoman(num-1, roman+"I")
	}
	return roman
}

// not submitted: make it better
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 || n > 3000 {
		return [][]int{}
	}
	arr := make([][]int, 0, n*3)

	x := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				x = nums[i] + nums[j] + nums[k]
				if x == 0 {
					arr = append(arr, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return arr

}


func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	prev, curr := dummy, head

	for curr != nil && curr.Next != nil {
		nxtPair := curr.Next.Next
		second := curr.Next

		second.Next = curr
		curr.Next = nxtPair
		prev.Next = second

		prev = curr
		curr = nxtPair
	}
	return dummy.Next
}


func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	wordList := []rune(word)
	w := 0
	_i, _j := 0, 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if wordList[w] == rune(board[i][j]) {
				if w == 0 || (_i == i || _j == j) {
					w++
					_i = i
					_j = j

				}
				if w == len(wordList) {
					return true
				}
			}
		}
	}
	return false
}


func insertNodeAtPosition(llist *ListNode, data int, position int) *ListNode {
	newNode := &ListNode{Val: data}

	if position == 0 {
		newNode.Next = llist
		return newNode
	}

	current := llist
	for i := 0; i < position-1; i++ {
		if current == nil || current.Next == nil {
			return llist
		}
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode

	return llist
}


func isBalanced(s string) string {
	if len(s) == 0 {
		return "YES"
	}
	stack := stack.NewStack()
	for _, v := range s {
		switch v {
		case '{':
			stack.Push('}')
		case '[':
			stack.Push(']')
		case '(':
			stack.Push(')')
		default:
			if stack.IsEmpty() || stack.Peek() != v {
				return "NO"
			}
			stack.Pop()

		}

	}
	if stack.IsEmpty() {
		return "YES"
	} else {
		return "NO"
	}
}




func SolveBoxes(clawPos int, boxes []int, boxInClaw bool) string {
	totalBoxes := 0
	for _, boxCount := range boxes {
		totalBoxes += boxCount
	}
	numStacks := len(boxes)
	targetHeight := totalBoxes / numStacks
	// extraBoxes := totalBoxes % numStacks
	PICK := "PICK"
	RIGHT := "RIGHT"
	LEFT := "LEFT"
	PLACE := "PLACE"
	i := clawPos
	if boxInClaw {
		if boxes[i] < targetHeight {
			return PLACE
		}
		if i-1 >= 0 {
			if boxes[i-1] < targetHeight {
				return LEFT
			}
		}
		if i+1 < numStacks {
			if boxes[i+numStacks] < targetHeight {
				return RIGHT
			}
		}
	} else {
		if boxes[i] > targetHeight {
			if i+1 < numStacks {
				if boxes[i+1] > targetHeight {
					return RIGHT
				}
			}
			if i-1 >= 0 {
				if boxes[i-1] < targetHeight || boxes[i-1] < boxes[i] {
					return LEFT
				} else {
					return PICK
				}
			}
		}
	}
	return ""
}
