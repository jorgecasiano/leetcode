/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        var result = new ListNode();
        var current = result;
        var acc = 0;
        
        while (l1 != null || l2 != null) {
            var num1 = l1 != null ? l1.val : 0;
            var num2 = l2 != null ? l2.val : 0;
            var sum = num1 + num2 + acc;
            
            acc = sum < 10 ? 0 : 1;
            sum = acc == 1 ? sum - 10 : sum;
            
            var newNode = new ListNode(sum);
            current.next = newNode;
            current = newNode;
            
            l1 = l1 != null ? l1.next : null;
            l2 = l2 != null ? l2.next : null;
        }
        
        if (acc == 1) {
            var newNode = new ListNode(1);
            current.next = newNode;
        }
        
        return result.next;
    }
}