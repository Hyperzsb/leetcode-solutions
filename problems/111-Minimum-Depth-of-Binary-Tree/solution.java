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
    public int minDepth(TreeNode root) {
        Queue<TreeNode> queue = new LinkedList<TreeNode>();
        
        if (root != null) {
            queue.offer(root);
        }
        
        int result = 0;
        while (!queue.isEmpty()) {
            result++;

            int size = queue.size();
            for (int i = 0; i < size; i++) {
                TreeNode node = queue.poll();

                if (node.left == null && node.right == null) {
                    return result;
                } else if (node.left == null && node.right != null) {
                    queue.offer(node.right);
                } else if (node.left != null && node.right == null) {
                    queue.offer(node.left);
                } else {
                    queue.offer(node.left);
                    queue.offer(node.right);
                }
            }
        }

        return result;
    }
}

