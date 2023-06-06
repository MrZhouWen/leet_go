package 二分查找

/**
题目：给定一个数组，寻找其中的极大值的下标，如果存在多个，则返回其中一个即可，假定 nums[-1] = nums[n] = -∞

二分的本质是分界：
	总结来说就是要向上走，因为题目假设了两头都是最小的，所以一直往上走总能找到峰值
	如果处于下降区间就把右边界左移，向左边的大数靠；
	如果处于上升区间就把左边界右移，向右边的大数靠；
*/
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if (mid-1 < 0 || nums[mid] > nums[mid-1]) && (mid+1 == len(nums) || nums[mid] > nums[mid+1]) { //处于其中一个峰值处
			return mid
		} else if (mid-1 < 0 || nums[mid] >= nums[mid-1]) && (mid+1 == len(nums) || nums[mid] <= nums[mid+1]) { //上升
			left = mid + 1
		} else { //下降
			right = mid - 1
		}
	}
	return 0
}
