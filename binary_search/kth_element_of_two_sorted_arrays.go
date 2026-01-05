package binarysearch

// Time Complexity: Achieved: O(log k), Best: O(log k)
// Space Complexity: Achieved: O(1), Best: O(1)
// KthElementOfTwoSortedArrays finds the k-th smallest element in the union of two sorted arrays.
// nums1, nums2: input sorted arrays
// k: the rank of the element to find
func KthElementOfTwoSortedArrays(nums1 []int, nums2 []int, k int) int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	for {
		if len(nums1) == 0 {
			return nums2[k-1]
		}

		if k == 1 {
			if nums1[0] < nums2[0] {
				return nums1[0]
			}
			return nums2[0]
		}

		i := min(len(nums1), k/2)
		j := k - i

		if nums1[i-1] < nums2[j-1] {
			nums1 = nums1[i:]
			k -= i
		} else {
			nums2 = nums2[j:]
			k -= j
		}

		if len(nums1) > len(nums2) {
			nums1, nums2 = nums2, nums1
		}
	}
}
