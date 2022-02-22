/**
 * @param {number[]} nums
 * @return {number[][]}
 */
 var subsets = function(nums) {
    let length = nums.length;
    let result = [];

    for(let i=0; i < 2 ** length; i++){    
        let sub = []
        for(let j =0 ; j < length; j++){
            if((i & (1 << j) ) > 0){               
                sub.push(nums[j])
            }
        }
        result.push(sub)
    }     
    
    return result
};