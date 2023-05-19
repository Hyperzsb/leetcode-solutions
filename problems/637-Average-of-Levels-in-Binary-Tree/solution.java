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
    public List<Double> averageOfLevels(TreeNode root) {
        Queue<TreeNode> nodeQ = new LinkedList<>();
        nodeQ.add(root);

        List<Double> result = new LinkedList<>();
        while (nodeQ.size() > 0) {
            int count = nodeQ.size();
            double sum = 0;
            for (int i = 0; i < count; i++) {
                TreeNode curr = nodeQ.remove();
                sum += (double)curr.val;

                if (curr.left != null) {
                    nodeQ.add(curr.left);
                }

                if (curr.right != null) {
                    nodeQ.add(curr.right);
                }
            }

            result.add(sum / count);
        }

        return result;
    }
}

