var connect = function (root) {
  if (!root) {
    return;
  }

  if (root.left) {
    root.left.next = root.right;
  }

  // using example from above, think about the node for (5). it does not have anything to the right of it but
  // it needs to be connected to (6). to do this, when we're processing node (2), we have access to node (2)'s "next", which is
  // node (3). we'll point (5) at (3)'s left, which is (6).
  if (root.right) {
    root.right.next = root.next ? root.next.left : null;
  }

  connect(root.left);
  connect(root.right);
};
