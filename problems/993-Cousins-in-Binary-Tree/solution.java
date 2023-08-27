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
    public boolean isCousins(TreeNode root, int x, int y) {
        LinkedList<TreeNode> nodes = new LinkedList<>();
        nodes.add(root);

        while (!nodes.isEmpty()) {
            int length = nodes.size();
            boolean xFlag = false, yFlag = false;

            for (int i = 0; i < length; i++) {
                TreeNode node = nodes.getFirst();
                nodes.removeFirst();

                if (node.val == x) {
                    if (yFlag) {
                        return true;
                    } else {
                        xFlag = true;
                    }
                }

                if (node.val == y) {
                    if (xFlag) {
                        return true;
                    } else {
                        yFlag = true;
                    }
                }

                if (node.left != null && node.right != null) {
                    if ((node.left.val == x && node.right.val == y) ||
                        (node.left.val == y && node.right.val == x)) {
                        return false;
                    } else {
                        nodes.add(node.left);
                        nodes.add(node.right);
                    }
                } else if (node.left != null) {
                    nodes.add(node.left);
                } else if (node.right != null) {
                     nodes.add(node.right);
                }
            }
        }

        return false;
    }
}

