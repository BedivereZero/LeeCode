/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
bool hasPathSum(struct TreeNode* root, int sum) {
	if (root == NULL)
		return false;
	if (root->right == NULL && root->left == NULL)
		return root->val == sum;
	return hasPathSum(root->left, sum - root->val) || hasPathSum(root->right, sum - root->val);
}
