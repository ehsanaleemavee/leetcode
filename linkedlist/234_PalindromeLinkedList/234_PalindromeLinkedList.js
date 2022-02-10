/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {boolean}
 */
 var isPalindrome = function(head) {
    let node = head;
    let prevnode = head;
		
    while(node != null){
        node.prev = prevnode; 
        prevnode = node      
        node = node.next;		  	
    }	

    tail = prevnode; 
	
    while(head && tail && (head.next != tail) && ( head != tail.prev)){
				
        if(tail.val != head.val)
            return false

        head = head.next;
        tail = tail.prev;
    }
    
    if(head && tail && head.val != tail.val)
        return false

    return true;
};