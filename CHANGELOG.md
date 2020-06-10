# Changelog

| Difficulty | Resolved | Total |
| :--------- | :------- | :---- |
| Easy       | 20       | 492   |
| Medium     | 14       | 837   |
| Hard       | 0        | 328   |

## 0024 Swap nodes in pairs

- Return `head.next` if head is not null
- First's next is third if fourth does not exist

## 0018 - 4Sum

- Two points
- $O(N^3)$

## 0016 - 3Sum closet

- Two points
- [Golang] Sort array via `sort.Ints`

## 0017 - Letter combinations of a phone number

- Strings contain '1' return `[]`

## 0015 - 3Sum

- Sort array
- Two points
- [Goalng] Define 2D array via `var result [][]int`

## 0019 - Remove nth node from end of list

- Two points `a` and `b`. Point `a` is `n` steps further than point `b`.
- If point `a` is None, it means the length of list is `n`. Remove the first node, that is, return `head.next`
- Move two points to next util `a` reach end (`a.next` is `None`).
- Remove point `b.next` when point `a` is `None`.

## 0012 - Integer to roman

- Linear

## 0011 - Container with most water

- Two points
- Update `height` every times

## 0008 - String to integer atoi

- Detect `a + b` out of the range via comparing `a` and `INT_MAX - b`
- Detect `a * b` out of the range via comparing `a` and `INT_MAX / b`
- Integer string startswith `+`, `-` or neither

## 0005 - Longest palindromic substring

- DP: F(i, j) = F(i + 1, j - 1) && s[i] == s[j]. When j - i < 2, F(i, j) = s[i] == s[j]
- [Golang] Create Two-dimensional array via `make([]int, m * n)`
- TODO: Manacher's algorithm

## 0088 - Merge sorted array

- Insert element from tail

## 0083 - Remove duplicates from sorted list

- Skip multi nodes

## 0070 - Climbing stairs

- DP: F(n) = F(n - 1) - F(n - 2)
- TODO: Binet's formula

## 0069 - Sqrtx

- Binary search
- Heron's formula

## 0067 - Add binary

- No reverse
- Full adder

## 0066 - Plus one

- Reverse array
- [Golang] `func reverse(input []int) []int`

## 0058 - Length of last word

- Trailing space

## 0053 - Maximum subarray

- DP: F(n): The maximum value of the subarray ending in the nth element.

## 0038 - Count and say

- [Golang] Use `bytes.Buffer` save temp string.

## 0035 - Search Insert Position

- **Binary Search**

## 0028 - Implement strstr

- [Python3] $O(m * (m - n))$
- [Golang] TODO: KMP

## 0027 - Remove element

- [Python3] Two points, one from head, the other from tail.
- [Python3] No need to exchange variables.

## 0026 - Remove duplicates from sorted array

- [Pytnon3] Two points.

## 0021 - Merge two sorted lists

- [Python3] If one list is empty, then reset is the other
- [Golang] Create point: `var head *ListNode`, `head := &ListNode{}`

## 0020 - Valid parentheses

- [Python3] Use list as stack
- [Golang] Create map: `map[type]type`
- [Golang] Traversing list via `for index, element := range list`
- [Golang] Difference between `byte` and `runc`
- [Golang] Append element to list via `append()`

## 0014 - Longest common prefix

- [Golang] Compare the string and the last of sorted strings?

## 0013 - Roman to integer

- [Python3] Compare next, add or minus

## 0009 - Palindrome number

- [Python3] Numbers less than zero are not palindrome.
- [Python3] Use `list[::-1]` get reversal.
- [Golang] Use modulus to get reversal.
- [Golang] Numbers endswith zero are not palidrome.
- [Golang] Reverse half.

## 0007 - Reverse interger

- [Python3] Transfer int to str.
- [Python3] Use `list[::-1]` to get reverse.
- [Golang] Modulus operation `%` preserves symbols
