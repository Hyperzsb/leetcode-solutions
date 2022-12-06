/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
struct TreeNode* invertTree(struct TreeNode* root){
    if(root == NULL || (root->left == NULL && root->right == NULL)) {
        return root;
    }
    
    struct TreeNode *left = root->left, *right = root->right;
    
    left = invertTree(left);
    right = invertTree(right);
    
    root->left = right;
    root->right = left;
    
    return root;
}

