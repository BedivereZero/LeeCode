/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
bool isSymmetric(struct TreeNode* root) {
	if (root->left == NULL && root->right == NULL)
		return true;
	if (root->left == NULL || root->right == NULL)
		return false;
	return (root->left->val == root->right->val) && isSymmetric(root->left) && isSymmetric(root->right);
}
