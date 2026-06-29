"use strict";

// Time Complexity: O(h)
// Space Complexity: O(1)


function TreeNode(val) {
  this.val = val;
  this.left = this.right = null;
}

const root = new TreeNode(6);
root.left = new TreeNode(2);
root.right = new TreeNode(8);

root.left.left = new TreeNode(0);
root.left.right = new TreeNode(4);
root.left.right.left = new TreeNode(3);
root.left.right.right = new TreeNode(5);

root.right.left = new TreeNode(7);
root.right.right = new TreeNode(9);

/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
var lowestCommonAncestor = function (root, p, q) {
  let holder = root;
  while (holder) {
    if (p.val < holder.val && q.val < holder.val) {
      holder = holder.left;
    } else if (p.val > holder.val && q.val > holder.val) {
      holder = holder.right;
    } else {
      return holder;
    }
  }
};

console.log("--- Test Vectors ---");

// Scenario A: The Clean Split (Expected: 6)
const resultA = lowestCommonAncestor(root, root.left, root.right);
console.log(`Scenario A LCA Value: ${resultA ? resultA.val : "null"}`);

// Scenario B: The Deep Subtree (Expected: 4)
const resultB = lowestCommonAncestor(root, root.left.right.left, root.left.right.right);
console.log(`Scenario B LCA Value: ${resultB ? resultB.val : 'null'}`);

// Scenario C: The Direct Descendant (Expected: 2)
const resultC = lowestCommonAncestor(root, root.left, root.left.right);
console.log(`Scenario C LCA Value: ${resultC ? resultC.val : 'null'}`);

