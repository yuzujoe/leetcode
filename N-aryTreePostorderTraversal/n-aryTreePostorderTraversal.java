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
    return postorder(root, new ArrayList<>());
  }

  private List<Integer> postorder(Node curr, List<Integer> nodes) {
    if (curr == null)
      return nodes;
    for (Node node : curr.children)
      postorder(node, nodes);
    nodes.add(curr.val);
    return nodes;
  }
}
