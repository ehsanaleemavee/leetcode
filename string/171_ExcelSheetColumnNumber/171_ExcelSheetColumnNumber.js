/**
 * @param {string} columnTitle
 * @return {number}
 */
 var titleToNumber = function(columnTitle) {
    let length = columnTitle.length;    
    let counter = length;
    let val =  0;

    while(counter > 0){
        let pos = length - counter;
        val = val * 26 + (columnTitle.charCodeAt(pos) - 64);
        counter = counter -1;
    }
    
    return val;
};