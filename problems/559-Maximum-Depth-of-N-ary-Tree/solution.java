/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/

class Solution {
    public int maxDepth(Node root) {
        if (root == null) {
            return 0;
        }
        
        Queue<Node> queue = new LinkedList<>();
        queue.add(root);
        int result = 0;

        while (!queue.isEmpty()) {
            result++;

            int length = queue.size();
            for (int i = 0; i < length; i++) {
                Node curr = queue.remove();

                for (Node child : curr.children) {
                    queue.add(child);
                }
            }
        }

        return result;
    }
}

