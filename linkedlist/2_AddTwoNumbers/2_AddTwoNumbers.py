# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        head = None
        node = None
        carry = 0
        
        while l1 != None or l2 != None:     
            sum = 0
           
            if l1 != None:
                sum = l1.val
                l1 = l1.next
            
            if l2 != None:
                sum += l2.val
                l2 = l2.next
            
            sum += carry
            carry = 0

            if sum > 9 :
                carry = sum // 10
                sum = sum % 10
            
            if head == None:
                node = ListNode(sum)
                head = node
                continue

            node.next = ListNode(sum)
            node = node.next

        if carry > 0:
            node.next = ListNode(carry)

        return head