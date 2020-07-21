# Changelog

## 0004 - Median Of Two Sorted Arrays

- Division and conquer

## 0005 - Longest palindromic substring

- DP: F(i, j) = F(i + 1, j - 1) && s[i] == s[j]. When j - i < 2, F(i, j) = s[i] == s[j]
- [Golang] Create Two-dimensional array via `make([]int, m * n)`
- TODO: Manacher's algorithm

## 0007 - Reverse interger

- [Python3] Transfer int to str.
- [Python3] Use `list[::-1]` to get reverse.
- [Golang] Modulus operation `%` preserves symbols

## 0008 - String to integer atoi

- Detect `a + b` out of the range via comparing `a` and `INT_MAX - b`
- Detect `a * b` out of the range via comparing `a` and `INT_MAX / b`
- Integer string startswith `+`, `-` or neither

## 0009 - Palindrome number

- [Python3] Numbers less than zero are not palindrome.
- [Python3] Use `list[::-1]` get reversal.
- [Golang] Use modulus to get reversal.
- [Golang] Numbers endswith zero are not palidrome.
- [Golang] Reverse half.

## 0010 - Regular Expression Matching

- Division and conquer
- $f(x,y)$ is result of matching `p[:x]` and `s[:y]`
  - $f(0,0)$ is `True`
  - $f(x,y)$ can be calculated from $f(x-1, y-1)$
    - When $p_{x-1}$ in alphabet, $f(x,y) = f(x-1,y-1) \land p_{x-1} = s_{y-1}$
    - When $p_{x-1}$ is `.`, $f(x,y) = f(x-1,y-1)$
    - When $p_{x-1}$ is `*`
      - When $p_{x-2}$ is `.`, $f(x,y) =\prod_{i=0}^y f(x-2, y-i)$
      - When $p_{x-2}$ is alphabet, $f(x,y) = \prod_{i=0}^y f(x-2, y-i) \land g(s, y, i, p_{x-2})$
        - $g(s, y, i, p_{x-2})$ String `s`, from `y-i` to `y`, whether it contains only the letter $p_{x-2}$

## 0011 - Container with most water

- Two points
- Update `height` every times

## 0012 - Integer to roman

- Linear

## 0013 - Roman to integer

- [Python3] Compare next, add or minus

## 0014 - Longest common prefix

- [Golang] Compare the string and the last of sorted strings?

## 0015 - 3Sum

- Sort array
- Two points
- [Goalng] Define 2D array via `var result [][]int`

## 0016 - 3Sum closet

- Two points
- [Golang] Sort array via `sort.Ints`

## 0017 - Letter combinations of a phone number

- Strings contain '1' return `[]`

## 0018 - 4Sum

- Two points
- $O(N^3)$

## 0019 - Remove nth node from end of list

- Two points `a` and `b`. Point `a` is `n` steps further than point `b`.
- If point `a` is None, it means the length of list is `n`. Remove the first node, that is, return `head.next`
- Move two points to next util `a` reach end (`a.next` is `None`).
- Remove point `b.next` when point `a` is `None`.

## 0020 - Valid parentheses

- [Python3] Use list as stack
- [Golang] Create map: `map[type]type`
- [Golang] Traversing list via `for index, element := range list`
- [Golang] Difference between `byte` and `runc`
- [Golang] Append element to list via `append()`

## 0021 - Merge two sorted lists

- [Python3] If one list is empty, then reset is the other
- [Golang] Create point: `var head *ListNode`, `head := &ListNode{}`

## 0022 - Generate parentheses

- DFS

## 0023 - Merge k Sorted Lists

- Division and conquer

## 0024 - Swap nodes in pairs

- Return `head.next` if head is not null
- First's next is third if fourth does not exist

## 0025 - Reverse Nodes in k-Group

- Related to [0092 - Reverse Linked List II](#0092---reverse-linked-list-ii)
- Revert link list several times
- Function `reverse(head, tail *ListNode) *ListNode`
  - Link list: `r -> 0 -> 1 -> 2 -> 3 -> 4 -> n`
  - `revert(0, 2)` reverse from 0, before 2, return `r -> 1 -> 0 -> 2 -> ...`
  - `revert(2, 4)` reverse from 2, before 3, return `r -> 3 -> 2 -> 4 -> ...`
  - `revert(head, nil)` reverse all of the link list, return `r -> 4 -> 3 -> 2 -> 1 -> 0 -> n`

## 0026 - Remove duplicates from sorted array

- [Pytnon3] Two points.

## 0027 - Remove element

- [Python3] Two points, one from head, the other from tail.
- [Python3] No need to exchange variables.

## 0028 - Implement strstr

- [Python3] $O(m * (m - n))$
- [Golang] TODO: KMP

## 0029 - Divide Two Integers

- Bit operation

## 0030 - Substring with Concatenation of All Words

- Binary

## 0031 - Next permutation

1. Found first element not increase from tail
2. Swap element if found
3. Sort elements after it

## 0032 - Longest Valid Parentheses

- Dymatic programming
  - s[x] is `(` $f(x) = 0$
  - s[x] is `)`
    - $s_{x-1} = ($
      - $x \in [0, 2): f(x) = 2$
      - $x \in [2, \infty]: f(x) = f(x - 2) + 2$

## 0033 - Search in rotated sorted array

- Get rotate pivot via binary-search
- Binary-search in two parts

## 0034 - Find first and last position of element in sorted array

- Binary search twice

## 0035 - Search Insert Position

- **Binary Search**

## 0036 - Valid Sudoku

- Check rows, columns, blocks separately

## 0037 - Sudoku Solver

- NOTICE: Get candidate from block $board[x/3*3+i/3,y/3*3+y\mod3]$
- TODO: `next()` returns the grid with the least possible

## 0038 - Count and say

- [Golang] Use `bytes.Buffer` save temp string.

## 0039 - Combination sum

- DFS, branchs = candidates
- [Golang] Copy slice via `copy(dst, src)`

## 0040 - Combination sum II

- DFS

## 0041 - First Missing Positive

- Sequences of length N can have up to N positive integers
- When all positive integers from 1 to N appear, the solution is the maximum value, N+1
- The range of solutions is 1 to N + 1
- Modify the sequence $nums$ so that $nums_i$ represents whether the number $i+1$ has appeared
- Check the sequence nums to find the first number that does not appear
- If all appear, the solution is $N + 1$

## 0042 - Trapping Rain Water

- Save height of pool

## 0044 - Wildcard Matching

- Dynatic Programming
- TODO: 2D matrix -> 1D matrix

## 0043 - Multiply Strings

- Multiply bit by bit and add

## 0045 - Jump Game II

- Greedy algorithm

## 0046 - Permutations

- DFS
- [Golang]: Not use `append()` in params, why?

## 0047 - Permutations II

- DFS
- Skip on branches

## 0048 - Rotate image

- Split 4 parts

## 0049 - Group anagrams

- Quick sort
- Transform string to number, sort and base-26 to base-10

## 0050 - Submissions

- Division and conquer

$$Pow({x}, {n})=Pow\left({x}^2, \dfrac{n}{2}\right) * Pow\left(x, n \mod 2\right)$$

## 0051 - N-Queens

- Deep First Search

## 0052 - N-Queens II

- Deep First Search

## 0053 - Maximum subarray

- DP: F(n): The maximum value of the subarray ending in the nth element.

## 0054 - Spiral matrix

- Circle by circle

## 0055 - Jump game

- DP

## 0056 - Merge intervals

- Quicksort with customized compare function

## 0057 - Insert Interval

- Binary Search
  - Find first `L` greater than or equal to `newInternal[0]` from `intervals[n][1]`
  - Find first `R` greater than `newInternal[1]` from `intervals[n][0]`
  - Insert from `L` before `R`
  - If `L` less than `R`, newInternal will merge element

## 0058 - Length of last word

- Trailing space

## 0059 - Spiral matrix II

- Linary

## 0060 - Permutation sequence

- Transform index `1...k` to `0...K`
- `result[n]` is `candidates[K // factorial(len(candidates) - 1)]`
- $K_{n+1} = K_n \mod {factorial\left[len\left(candidates\right) - 1\right]}$

## 0061 - Rotate List

- Circle and break
- Mod length of list

## 0062 - Unique paths

- Combination(n, k)
- Calculating factorial may overflow
- Multiply then divide to avoid overflow
- Divided from small to large to ensure divisible

## 0063 - Unique paths II

- Dynamic programming

## 0064 - Minimum Path Sum

- Dynamic programming

## 0065 - Valid Number

- **SHIT**
- COPY FROM: [Golang的非自动机解法](https://leetcode-cn.com/problems/valid-number/solution/golangde-fei-zi-dong-ji-jie-fa-by-xijiz/)

## 0066 - Plus one

- Reverse array
- [Golang] `func reverse(input []int) []int`

## 0067 - Add binary

- No reverse
- Full adder

## 0068 - Text Justification

- [Golang] `bytes.Buffer.Len` and `bytes.Bufer.WriteString`

## 0069 - Sqrtx

- Binary search
- Heron's formula

## 0070 - Climbing stairs

- DP: F(n) = F(n - 1) - F(n - 2)
- TODO: Binet's formula

## 0071 - Simplyfy path

- TODO: base on string

## 0072 - Edit Distance

- Dynatic programming

## 0073 - Set matrix zeros

- Store x, y to be zeros

## 0074 - Search a 2D matrix

- Binary search, twice
- Transform 2D to 1D, binary search once

## 0075 - Sort colors

- `nums[:a]` are red, `nums[a:b]` are white, `nums[c:]` are blue
- `nums[b:c]` are unknown

## 0076 - Minimum Window Substring

- Two points
- Slice window
- Use `queue` record target character

## 0077 - Combinations

- DFS

## 0078 - Subsets

- Use `make` and `copy` instead of `append`

## 0079 - Word search

- DFS
- Record footprint

## 0080 - Remove duplicates from sorted array II

## 0081 - Search in rotated sorted array II

- Search pivot

## 0082 - Remove duplicates from sorted list II

- Skip duplicates
- Move to next node if not skip

## 0083 - Remove duplicates from sorted list

- Skip multi nodes

## 0084 - Largest Rectangle in Histogram

- Recursion
- Scan one time

## 0085 - Maximal Rectangle

- Transform to histogram
- Monotonic stack

## 0086 - Parition list

- Two lists contains low and high
- Combine two lists

## 0087 - Scramble String

- Recursion
- Pruning

## 0088 - Merge sorted array

- Insert element from tail

## 0089 - Gray Code

- Bit operation
- $O(N)$

## 0090 - Subsets II

## 0091 - Decode ways

- Dynamic programming

## 0092 - Reverse Linked List II

- Define `root` point to `head`
- `func reverse(head *ListNode, tail *ListNode) *ListNode` revert `list[head:tail]`

## 0093 - Restore IP Addresses

- Deep First Search
- Not remove number `0`: `010010` -> `[0.10.0.10, 0.100.1.0]`
- [Golang]: `strconv.ParseUint`, `strconv.FormatUint`

## 0094 - Binary tree inorder traversal

- Deep First Search
- Recursion
- Stack and loops

## 0095 - Uniqute Binary Search Tree II

- Dynamic programming
- $Cache(x, y)$ means BST with length y starting at x
- $Cache(x, 0)$ is ${ nil }$

## 0096 - Unique Binary Search Trees

- Dynamic programming
- $F_{n+1}=\sum_{i=0}^{n} F_i \times F_{n-i}$

## 0097 - Interleaving String

- Recursion

## 0098 - Validate Binary Search Tree

- Recursion

## 0099 - Recover Binary Search Tree

- Scan twice: from head, from tail

## 0100 - Same Tree

- Recursion

## 0101 - Symmetric Tree

- Recursion

## 0102 - Binary Tree Level Order Traversal

- Recursion

## 0103 - Binary Tree Zigzag Level Order Traversal

- Recursion
- Orientation

## 0104 - Maximum Depth of Binary Tree

- Recursion

## 0105 - Construct Binary Tree from Preorder and Inorder Traversal

- Recursion
- `preorder[0]` is root
- Find the root in array `inorder`, `inorder[:root]` is left tree, `inorder[root+1:]` is right tree

## 0106 - Construct Binary Tree from Inorder and Postorder Traversal

- Recursion
- `postorder[len(postorder)-1]` is root
- Find the root in array `inorder`, `inorder[:root]` is left tree, `inorder[root+1:]` is right tree

## 0107 - Binary Tree Level Order Traversal II

- Recursion

## 0108 - Convert Sorted Array to Binary Search Tree

- Recursion

## 0109 - Convert Sorted List to Binary Search Tree

- Recursion

## 0110 - Balanced Binary Tree

- Recursion

## 0111 - Minimum Depth of Binary Tree

- Recursion
- NOTICE: Tree with null is not 1

## 0112 - Path Sum

- Recursion

## 0113 - Path Sum II

- Deep First Search

## 0114 -  Flatten Binary Tree to Linked List

- Recursion

## 0118 - Pascal's Triangle

- Nothing

## 0119 - Pascal's Triangle II

- Nothing

## 0120 - Triangle

- Dynamic Programming

## 0121 - Best Time to Buy and Sell Stock

- Dynamtic Programing
- $f_i$ means the minimum in the frist i days

## 0122 - Best Time to Buy and Sell Stock II

- Linearity

## 0124 - Binary Tree Maximum Path Sum

- Recursion
- Update `Node.Val` to the maximum path starting from `Node`
- `maxPathSum(Node)` is the maximum of:
  - `maxPathSum(Node.Left)`
  - `maxPathSum(Node.Right)`
  - `Node.Val`
  - `Node.Val + Node.Left.Val`
  - `Node.Val + Node.Right.Val`
  - `Node.Val + Node.Left.Val + Node.Right.Val`

## 0125 - Valid Palindrome

- Two Points

## 0129 - Sum Root to Leaf Numbers

- Recursion

## 0130 - Surrounded Regions

- Deep First Search
- Mark 'O' from bounds
- Change unmarked 'O' to 'X'

## 0131 - Palindrome Partitioning

- Deep First Search

## 0134 - Gas Station

- Generate array `R`, $R_i$ means relative gas in tank when the car arrive next gas station
- `R[i] = R[i] + R[i - 1]`, means the i th gas station is last, relative gas in tank against starting.

## 0135 - Candy

- Traversal once
- Two points, traversal equal, increase and decrease then loop
- The peek is the max of increase and decrease parts

## 0136 - Single Number

- Exclusive Or
- Twice = nothing
- It is not necessary to record what numbers not single

## 0137 - Single Number II

- Use 2 bit record the number of occurrences of 1: `00`, `01`, `10`
- Use **Karnaugh map** to facilitate the simplification of Boolean algebra

## 0141 - Linked List Cycle

- Two points: slow and fast

## 0142 - Linked List Cycle II

- Three points: x1, x2, x3 speed
- When `x2 == x3`, The distance x1 travelsal is equal to the length of the cycle
- Create another point `st` from head, move `st` and `x1` until `st == x1`

## 0143 - Reorder List

- Traversal once
- Use **Stack** record second half

## 0160 - Intersection of Two Linked Lists

- Use stack storage node

## 0167 - Two Sum II - Input array is sorted

- Map

## 0168 - Excel Sheet Column Title

- Base 26

## 0169 - Majority Element

- Map
- TODO: Boyer Moore

## 0171 - Excel Sheet Column Number

- Linary

## 0172 - Factorial Trailing Zeroes

- Count factor 2 & 5

## 0189 - Rotate Array

- Rotate k elements, one time

## 0190 - Reverse Bits

- Linary

## 0415 - Add strings

- Full adder
- [Golang] Create array via `make([]int, n)`
- [Golang] Combine bytes to string via `types.Buffer`
  - `Buffer.WriteByte`: Write byte to end
  - `Buffer.String`: Get as string
