/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
struct TreeNode* invertTree(struct TreeNode* root) {
	struct TreeNode* l = NULL;
	struct TreeNode* r = NULL;
	if(root == NULL)
		return NULL;
	l = invertTree(root->right);
	r = invertTree(root->left);
	root->right = r;
	root->left = l;
	return root;
}
