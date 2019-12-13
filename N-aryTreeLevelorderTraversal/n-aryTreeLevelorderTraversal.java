/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val,List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/
class Solution {
  private List<List<Integer>> list = new ArrayList<>();

  public List<List<Integer>> levelOrder(Node root) {
    traverse(root, 1);
    return list;
  }

  private void traverse(Node node, int level) {
    if (node == null)
      return;
    List<Integer> l;
    if (list.size() < level) {
      l = new ArrayList<>();
      list.add(l);
    } else {
      l = list.get(level - 1);
    }
    l.add(node.val);
    for (Node n : node.children) {
      traverse(n, level + 1);
    }
  }
}
