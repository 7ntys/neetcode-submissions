func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	half := (len(nums1) + len(nums2)) / 2

	i, j := 0, 0
	prev, curr := 0, 0

	for k:=0;k<=half;k++{
		prev = curr

		if i < len(nums1) && j < len(nums2) {
			if nums1[i] <= nums2[j] {
				curr = nums1[i]
				i++
			} else {
				curr = nums2[j]
				j++
			}
		} else {

			if i < len(nums1) {
				curr = nums1[i]
				i++
			} else if j < len(nums2) {
				curr = nums2[j]
				j++
			}
		}

	}
	fmt.Println(prev,curr)
	if (len(nums1) + len(nums2)) % 2 == 1 {return float64(curr)}
	return (float64(prev) + float64(curr)) / 2
}
