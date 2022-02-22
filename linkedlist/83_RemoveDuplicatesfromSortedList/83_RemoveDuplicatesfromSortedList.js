/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function(head) {
    if(head == null)
        return head;

    let prev = head;    
    let temp = head;       
    temp = temp.next;    

    while(temp != null){    
        
        if(prev.val == temp.val){
            let temp2 = temp;    
            temp = temp.next;
            delete temp2;           
        }
        else{ 
           prev.next = temp;
           prev = prev.next;
           temp = temp.next
       }
    }

    if(prev != null && prev.next != null)
        prev.next = null;
    
    return head;
    
};