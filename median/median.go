package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sizeIseven := (len(nums1)+len(nums2))%2 == 0
	medianIndex := (len(nums1) + len(nums2)) / 2
	//to do, use 2 variables to store the two elements needed to compute
	//the median
	mergedArray := make([]int, medianIndex+1, medianIndex+1)
	i := 0
	j := 0
	k := 0
	for k < medianIndex+1 {
		var first *int
		var second *int
		if i < len(nums1) {
			first = &nums1[i]
		}
		if j < len(nums2) {
			second = &nums2[j]
		}
		if first != nil && second != nil {
			if *first < *second {
				mergedArray[k] = *first
				i++
			} else {
				mergedArray[k] = *second
				j++
			}
		} else if first != nil && second == nil {
			mergedArray[k] = *first
			i++
		} else if first == nil && second != nil {
			mergedArray[k] = *second
			j++
		}
		k++
	}
	if sizeIseven {
		return float64(mergedArray[medianIndex-1]+mergedArray[medianIndex]) / 2.0
	} else {
		return float64(mergedArray[medianIndex])
	}
}
