/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
 var twoSum = function(nums, target) {
    
    let map = new Map();
    let result =[];
    
    for (let i=0;i<nums.length;i++){
        let item = nums[i];
        let diff = target - item;        
        if(map.has(diff)) { 
            result = [map.get(diff), i];
            break;
        }
        map.set(item, i);
    }
    return result;
};