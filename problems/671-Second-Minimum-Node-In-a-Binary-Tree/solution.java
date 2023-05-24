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
    public int findSecondMinimumValue(TreeNode root) {
        HashSet<Integer> hs = new HashSet<>();
        traverse(root, hs);

        Integer[] nums = new Integer[hs.size()];
        nums = hs.toArray(nums);

        Arrays.sort(nums);

        if (nums.length < 2) {
            return -1;
        } else {
            return nums[1];
        }
    }

    public void traverse(TreeNode root, HashSet<Integer> hs) {
        if (root == null) {
            return;
        }

        traverse(root.left, hs);
        hs.add(root.val);
        traverse(root.right, hs);
    }
}

