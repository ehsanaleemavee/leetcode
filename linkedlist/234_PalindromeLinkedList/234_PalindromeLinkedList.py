# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:        
        
        if head.next == None:
            return True

        first_pointer = head
        second_pointer = head       

        while first_pointer != None and second_pointer != None:                     

            if second_pointer.next == None:
                break
            else:                
                second_pointer = second_pointer.next.next

            first_pointer = first_pointer.next
        
        mid = first_pointer       
           
        if second_pointer != None:
            mid = mid.next        
           
        reverse = self.reverseLinkedList(mid)      

        while head != None and reverse != None:         
            if head.val != reverse.val:
                return False
            
            head = head.next
            reverse = reverse.next            
       
        return True
    
    def pushToTop(self, val, head: Optional[ListNode]) -> ListNode:
        node = ListNode(val)
        node.next = head
        
        return node

    def reverseLinkedList(self, head: Optional[ListNode]) -> ListNode:
        reverse = None

        while head != None:                        
            reverse = self.pushToTop(head.val, reverse)
            head = head.next

        return reverse