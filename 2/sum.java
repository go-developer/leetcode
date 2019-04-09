/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
 }
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        boolean isHasExtraAdd = false;
        ListNode result = new ListNode(0);
        ListNode first = result;
        int tmp = 0;
        while (true) {
            if (null == l1 && null == l2) {
                break;
            }
            if (null == l1) {
                l1 = new ListNode(0);
            }

            if (null == l2) {
                l2 = new ListNode(0);
            }

            tmp = l1.val + l2.val;
            if (isHasExtraAdd) {
                tmp = tmp + 1;
                isHasExtraAdd = false;
            }


            if (tmp >= 10) {
                isHasExtraAdd = true;
                tmp = tmp - 10;
            }

            result.next = new ListNode(tmp);
            result = result.next;

            l1 = l1.next;
            l2 = l2.next;
        }

        if (isHasExtraAdd) {
            result.next = new ListNode(1);
        }

        return first.next;
    }
}