/**
 * @param {number} num
 * @return {number}
 */
 var addDigits = function(num) {
    
    if(num <10)
         return num;
    
    if(num == 10)
        return 1;
    
     while(num >= 10){
        let length = num.toString().length; 
        length = 10 ** (length-1); 
        let quotient = Math.floor(num / length);
        let reminder = num % length; 
        
        num = quotient + reminder;
    }

    return num;
};
