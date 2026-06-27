
// Time Complexity: O(N) 
// Space Complexity: O(H) where H is the height of the tree

function TreeNode(val, left, right) {
  this.val = val === undefined ? 0 : val;
  this.left = left === undefined ? null : left;
  this.right = right === undefined ? null : right;
}

/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
var invertTree = function (root) {
  if (root == null) return null;
  let holder = root.left;
  root.left = root.right;
  root.right = holder;
  invertTree(root.left);
  invertTree(root.right);
  return root;
};