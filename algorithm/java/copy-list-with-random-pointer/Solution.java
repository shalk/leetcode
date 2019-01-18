
/**
 * Definition for singly-linked list with a random pointer.
 * class RandomListNode {
 *     int label;
 *     RandomListNode next, random;
 *     RandomListNode(int x) { this.label = x; }
 * };
 */
public class Solution {
    public RandomListNode copyRandomList(RandomListNode head) {
        if (head == null) {
            return null;
        }

        RandomListNode dummy = new RandomListNode(0);
        RandomListNode q = dummy;

        HashMap<RandomListNode,  RandomListNode> m = new HashMap<>();
        // create new List
        for(RandomListNode p = head; p != null; p = p.next, q = q.next) {
            RandomListNode n;
            if (!m.containsKey(p)) {
                n = new RandomListNode(p.label);
                m.put(p, n);
            } else {
                n = m.get(p);
            }

            RandomListNode r;
            if (p.random == null) {
                r = null;
            } else if (m.containsKey(p.random)) {
                r = m.get(p.random);
            } else {
                r = new RandomListNode(p.random.label);
                m.put(p.random, r);
            }
            n.random = r;
            q.next = n;

        }



        return dummy.next;
    }
}
