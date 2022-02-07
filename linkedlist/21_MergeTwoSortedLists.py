# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        head = None
        node = None

        while list1 != None or list2 != None:
            data1 = 101
            data2 = 101
            pick = 101

            if list1 != None:
                data1 = list1.val
            
            if list2 != None:
                data2 = list2.val
            
            if data1 > data2:
                pick = data2
                list2 = list2.next
            else:
                pick = data1
                list1 = list1.next
            
            if head == None:
                node = ListNode(pick)
                head = node
                continue

            node.next = ListNode(pick)
            node = node.next

        return head