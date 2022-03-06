# LeetCode_1356

## 題目描述

給定一個陣列 nums，寫一個 function 計算出每個元素中比目前元素值還要小的元素個數，這個回傳結果是一個陣列。

## 舉例

### Example 1

```shell
Input: nums = [8, 1, 2, 2, 3]
Output: [4, 0, 1, 1, 3]
```

### Example 2

```shell
Input: nums = [6, 5, 4, 8]
Output: [2, 1, 0, 3]
```

### Example 3

```shell
Input: nums = [7, 7, 7, 7]
Output: [0, 0, 0, 0]
```

## 解法

```golang
func SmallerNumnersThanCurrent(nums []int) []int {
	result := []int{}
	for idx := range nums {
		result = append(result, countSmall(nums, idx))
	}
	return result
}

func countSmall(nums []int, idx int) int {
	result := 0
	for i, val := range nums {
		if i == idx {
			continue
		}
		if val < nums[idx] {
			result += 1
		}
	}
	return result
}
```