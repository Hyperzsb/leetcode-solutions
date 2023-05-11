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
    public int findTilt(TreeNode root) {
        calcTilt(root);
        return sumTilt(root);
    }

    public int calcTilt(TreeNode root) {
        if (root == null) {
            return 0;
        }

        int currVal = root.val;
        int leftVal = calcTilt(root.left);
        int rightVal = calcTilt(root.right);
        root.val = diff(leftVal, rightVal);

        return currVal + leftVal + rightVal;
    }

    public int sumTilt(TreeNode root) {
        if (root == null) {
            return 0;
        }

        int leftSum = sumTilt(root.left);
        int rightSum = sumTilt(root.right);

        return root.val + leftSum + rightSum;
    }

    public int diff(int a, int b) {
        if (a > b) {
            return a - b;
        } else {
            return b - a;
        }
    }
}

