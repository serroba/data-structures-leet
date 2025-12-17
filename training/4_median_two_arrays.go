package training

import "slices"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	combined := make([]int, 0, len(nums1)+len(nums2))
	for i := 0; i < len(nums1); i++ {
		j := 0
		for len(nums2) > 0 && j < len(nums2) {
			if nums1[i] < nums2[j] {
				combined = append(combined, nums1[i])
				j++
				break
			} else {
				combined = append(combined, nums2[j])
				nums2 = nums2[1:]
				j = 0
			}
		}
		if len(nums2) == 0 {
			combined = slices.Concat(combined, nums1[i:])
			break
		}

	}
	if len(nums2) > 0 {
		combined = slices.Concat(combined, nums2)
	}
	if len(combined)%2 == 0 {
		return float64(combined[len(combined)/2-1]+combined[len(combined)/2]) / 2.0
	}
	return float64(combined[len(combined)/2])
}
