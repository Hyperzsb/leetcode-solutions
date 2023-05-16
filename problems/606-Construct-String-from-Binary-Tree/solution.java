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
    public String tree2str(TreeNode root) {
        StringBuilder result = new StringBuilder(Integer.toString(root.val));

        if (root.left == null && root.right == null) {
            ;
        } else if (root.left == null && root.right != null) {
            result.append("()")
                .append("(").append(tree2str(root.right)).append(")");
        } else if (root.left != null && root.right == null) {
            result.append("(").append(tree2str(root.left)).append(")");
        } else {
            result.append("(").append(tree2str(root.left)).append(")")
                .append("(").append(tree2str(root.right)).append(")");
        }

        return result.toString();
    }
}

