/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */

class Solution {
    private TreeNode traverse(TreeNode prev, TreeNode curr) {
        if (curr == null) {
            return prev;
        }

        prev = traverse(prev, curr.left);

        prev.right = curr;
        curr.left = null;

        prev = traverse(curr, curr.right);

        return prev;
    }

    public TreeNode increasingBST(TreeNode root) {
        TreeNode newRoot = new TreeNode();
        traverse(newRoot, root);

        return newRoot.right;
    }
}

