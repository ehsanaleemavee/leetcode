# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        
        if head == None:
            return head
        
        my_dic = {}
        temp = head

        while temp != None:            
            if my_dic.get(temp.val, -1) == -1 :               
                my_dic[temp.val] = temp
            
            temp = temp.next

        result = None
        prev = None
        
        for val in my_dic.values(): #for i, val in my_dic.items():
            if result == None:
                prev = val
                result = prev
                continue

            prev.next = val
            prev = prev.next
            
        prev.next = None    

        return result