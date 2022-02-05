func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

   	result := []int {-1,-1}

   	for i := 0; i < len(nums); i++ {

	   item := nums[i]
	   var diff int = target - item
   
	   if m[diff] > 0 {
		   result[0] = m[diff] - 1
		   result[1]=  i
		   break;
	   }

	   m[item] = i+1;
   	}
   
   	return result
}