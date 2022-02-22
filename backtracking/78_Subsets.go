func subsets(nums []int) [][]int {
	length := len(nums)
	size := int(math.Pow(2, float64(length)))
	result := make([][]int, size)

	for i := 0; i < size; i++ {
		s := make([]int, 0, length)

		for j := 0; j < length; j++ {

			if (i & (1 << j)) > 0 {
				s = append(s, nums[j])
			}
		}

		result[i] = s
	}

	return result
}