/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public int getMinimumDifference(TreeNode root) {
        TreeNode pre = null;
        int ret = Integer.MAX_VALUE;
        Deque<TreeNode> stack = new ArrayDeque<>();
        for (TreeNode p = root; p != null; p = p.left) {
            stack.push(p);
        }
        while(!stack.isEmpty()) {
            TreeNode p = stack.pop();
            // stack.pop();
            if (pre != null) {
                ret = Math.min(ret, Math.abs(p.val-pre.val));
            }
            pre = p;
            if (p.right != null) {
                p = p.right;
                for (TreeNode k = p; k != null; k = k.left) {
                    stack.push(k);
                }
            }
        }
        return ret;
    }
}
