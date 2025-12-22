package leet

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	middle := len(nums) / 2

	node := &TreeNode{Val: nums[middle]}
	node.Left = sortedArrayToBST(nums[:middle])
	node.Right = sortedArrayToBST(nums[middle+1:])

	return node
}
