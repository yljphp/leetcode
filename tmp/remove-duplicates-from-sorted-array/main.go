package main

/**
26. 删除排序数组中的重复项
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。

@see https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
 */
func main() {

	nums := []int{
		1,1,2,
	}

	removeDuplicates(nums)
}


func removeDuplicates(nums []int) int {

	if len(nums) < 2 {
		return  len(nums)
	}

	resNum := 0

	for i:=0; i<len(nums); i++ {

		if i == len(nums) -1 {
			break
		}

		if nums[i] != nums[i+1]{


			resNum++
		}else{

			nums[i+1] = nums[i]
		}

	}


	return resNum
}