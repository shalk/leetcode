package lowestcommonancestor;


/**
 * Definition for a binary tree node.
 * public class TreeNode {
 * int val;
 * TreeNode left;
 * TreeNode right;
 * TreeNode(int x) { val = x; }
 * }
 */

/**
 *  有三种情况:
 *  1. p 和q  分别是w的左和右子树； w是LCA 如果在同一侧，那root的下一级更满足条件。
 *  2. q 在p 的左子树 或者右子树 ； q是LCA
 *  3. p 在q的左子树 或者右子树; p 是LCA
 *  模拟操作：
 *  1. 如果出现了情况2 和3 直接就返回了。
 *  2. 如果发现p, q 是root的同一侧，则找到是左侧还是右侧，进入下一级进行查找。
 */
class Solution1 {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (p == root) {
            return p;
        }
        if (q == root) {
            return q;
        }
        if (root.val < p.val && root.val < q.val) {
            return lowestCommonAncestor(root.right, p, q);
        }
        if (root.val > p.val && root.val > q.val) {
            return lowestCommonAncestor(root.left, p, q);
        }
        return root;
    }
}

class Solution2 {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {

        TreeNode w = root;
        // swap p q, let p.val < q.val
        if (p.val > q.val) {
            TreeNode tmp = p;
            p = q;
            q = tmp;
        }

        while (true) {
            if (w == p || w == q) {
                return w;
            }
            if (w.val > q.val) {
                w = w.left;
            } else if (w.val < p.val) {
                w = w.right;
            } else {
                return w;
            }
        }
    }
}

public class LowestCommonAncestor {


    public static void main(String[] args) {
        System.out.println("Hello");
    }
}
