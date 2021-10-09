# [9.Palindrome Number(回文数)](https://leetcode-cn.com/problems/palindrome-number/)

*难度：简单*

## 题目

给你一个整数 `x` ，如果 x 是一个回文整数，返回 `true` ；否则，返回 `false` 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，`121` 是回文，而 `123` 不是。

 
**示例 1：**
```
输入：x = 121
输出：true
```

**示例 2：**
```
输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
```

**示例 3：**
```
输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
```

**例 4：**
```
输入：x = -101
输出：false
```

**提示：**

- `-231 <= x <= 231 - 1`
 

**进阶**：你能不将整数转为字符串来解决这个问题吗？

## 解题思路

- 判断一个整数是不是回文数。
- 简单题。注意会有负数的情况，负数，个位数， `10` 都不是回文数。其他的整数再按照回文的规则判断。

## 代码

```go
package leetcode

import "strconv"

// 解法一
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	sz := len(arr)
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

// 解法二 数字转字符串
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

```
