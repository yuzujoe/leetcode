/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */
/**
 * @param {Node} root
 * @return {number[]}
 */

var preorder = function (root) {
  const array = [];

  const traverseTree = (node) => {
    if (!node) return;

    array.push(node.val);

    node.children.forEach((child) => traverseTree(child));
  }

  traverseTree(root);

  return array;
};
