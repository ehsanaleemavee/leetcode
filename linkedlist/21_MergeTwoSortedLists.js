/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
 var mergeTwoLists = function(list1, list2) {
    var head;
    var node;
    console.log(list1)
    while(list1 != null || list2 != null){

        let data1 = list1?.val ?? 101;
        let data2 = list2?.val ?? 101;

        let pick = -101;

        if(data1 > data2){
            pick = data2;
            list2 = list2.next
        }
        else{
            pick = data1;
            list1 = list1.next
        }

        if(head == null){
            node = new ListNode(pick)
            head = node
        }
        else{
            node.next = new ListNode(pick)
            node = node.next
        }
    }
    if(head == null)
        head = ListNode()
    return head;
};