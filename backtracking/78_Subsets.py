class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        length = len(nums)
        size = pow(2, length)
        result = []
        for i in range(size):
            sub = []

            for j in range(length):
                if i & (1 << j) > 0:
                    sub.append(nums[j])
                    
            result.append(sub)

        return result
