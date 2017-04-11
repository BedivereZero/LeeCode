/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
bool isBalanced(struct TreeNode* root) {
	if (root == NULL)
		return true;
	if (root->left == NULL && root->right && (root->right->left != NULL || root->right->right != NULL))
		return false;
	if (root->right == NULL && root->left && (root->left->left != NULL || root->left->right != NULL))
		return false;
	return isBalanced(root->left) && isBalanced(root->right);
}
