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
    int prev;
    int minDiff;

    public int diff(int a, int b) {
        if (a > b) {
            return a - b;
        } else {
            return b - a;
        }
    }
    
    public void traverse(TreeNode root) {
        if (root == null) {
            return;
        }

        traverse(root.left);

        if (prev == -1) {
            prev = root.val;
        } else {
            minDiff = minDiff < diff(root.val, prev) ? minDiff : diff(root.val, prev);
            prev = root.val;
        }

        traverse(root.right);
    }

    public int getMinimumDifference(TreeNode root) {
        prev = -1;
        minDiff = 100000;
        traverse(root);

        return minDiff;
    }
}

