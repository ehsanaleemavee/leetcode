class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        my_dict = {}
        result = 0
        for i in nums:           
            my_dict[i] = 2 if my_dict.get(i, -1) > -1 else 1           
       
        for i in my_dict.keys():
            
            if my_dict[i] == 1:
                result = i                
                break

            i = i + 1

        return result