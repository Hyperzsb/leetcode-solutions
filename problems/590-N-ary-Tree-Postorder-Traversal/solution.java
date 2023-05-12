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
    public List<Integer> postorder(Node root) {
        LinkedList<Integer> result = new LinkedList<>();

        if (root == null) {
            return result;
        }

        for (Node child : root.children) {
            result.addAll(postorder(child));
        }

        result.add(root.val);

        return result;
    }
}

