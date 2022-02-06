package main

import "fmt";

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    carry := 0
    var head *ListNode
    var node *ListNode
    
    for l1 != nil || l2 != nil {        
        sum := 0
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        sum += carry
        carry = 0;

        if sum > 9 {
            carry = 1
            sum = sum % 10
        }
        
        if head == nil { 
            node = new(ListNode)
            node.Val = sum
            head = node            
            continue
        }
        
        node.Next = new(ListNode)
        node = node.Next        
        node.Val = sum
        sum=0               
    }

    if carry > 0 {
        node.Next = new(ListNode)
        node.Next.Val = carry
    }
   
    return head;
}

func createLinkedList(arr []int) *ListNode {
    //*ListNode head :=  nil;
    //*ListNode node :=  nil;

    var head *ListNode
    var node *ListNode
    

    for i := 0; i < len(arr); i++ {
        
        if i == 0 {            
            node = new(ListNode)
            node.Val = arr[i]
            head = node           
            continue
        }
        
        node.Next = new(ListNode) //{Val:arr[i]}
        node = node.Next
        node.Val = arr[i]  

    }
    
    return head;
}

func traverse(l1 *ListNode){
    
    for l1 != nil {
        fmt.Println(l1.Val)
        l1 = l1.Next
    }
}

func main (){

    // arr1 := []int{2,4,3}
    // arr2 := []int{5,6,4}

    // arr1 := []int{9,9,9,9,9,9,9}
    // arr2 := []int{9,9,9,9}

    // arr1 := []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
    // arr2 := []int{5,6,4}

    // arr1 := []int{0}
    // arr2 := []int{0}

    arr1 := []int{0}
    arr2 := []int{1}

    l1 := createLinkedList(arr1);
    l2 := createLinkedList(arr2);
    

    //traverse(l1)
    //traverse(l2)

    traverse(addTwoNumbers(l1, l2))

}

//traverse(addTwoNumbers(l1, l2));