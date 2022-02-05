class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        my_dict = {}
        result = []
        for i in range(0, len(nums)):
            
            item = nums[i]
            diff = target - nums[i]
           
            if my_dict.get(diff, -1) > -1:
                result.append(i)
                result.append(my_dict[diff])

            my_dict[item] = i 
        
        return result