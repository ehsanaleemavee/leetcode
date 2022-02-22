/**
 * @param {number[]} nums
 * @return {number}
 */
 var singleNumber = function(nums) {
    let length = nums.length;
    mergesort(nums, 0, length - 1);

    let i = 0;
    while( i + 1 <= length){
        if(nums[i] != nums[i+1] )
            break;

        i = i + 2;
    }

    return nums[i]

};

function mergesort(arr, left, right){
    
    if(left >= right){
        return 
    }

    let mid = left + parseInt((right-left)/2);
    mergesort(arr, left, mid);
    mergesort(arr, mid + 1, right);
    sort(arr, left, mid, right);

}

function sort(arr, left, mid, right){

    let n1 = mid - left + 1;
    let n2 = right - mid;
    let left_side = []
    let right_side = []

    let l = 0;
    let r = 0;   
    
    while(l < n1){
        left_side[l] = arr[left + l];
        l++;       
    }
   
    while(r < n2){
        right_side[r] = arr[mid + r + 1];
        r++;     
    }
     
    let i = 0;
    let j = 0;
    
    let k = left;

    while(i < n1  &&  j < n2){

       if(left_side[i] > right_side[j]){           
            arr[k] = right_side[j];           
            j++;
        }        
        else{       
            arr[k] = left_side[i];
            i++;
        }
        k++;
    }

    while(i < n1){
        arr[k] = left_side[i];
        k++;
        i++;
    }

    while(j < n2){
        arr[k] = right_side[j];
        k++;
        j++;
    }    
}