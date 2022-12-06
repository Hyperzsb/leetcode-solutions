/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
int height(struct TreeNode *root, int current, bool *isBal) {
    if(!(*isBal)) {
        return current;
    }
    
    if(root->left == NULL && root->right == NULL) {
        return current;
    }
    
    if(root->left != NULL && root->right == NULL) {
        int leftHeight = height(root->left, current + 1, isBal);
        if (leftHeight > current + 1) {
            *isBal = false;
        }
        return leftHeight;
    } else if(root->left == NULL && root->right != NULL) {
        int rightHeight = height(root->right, current + 1, isBal);
        if(rightHeight > current + 1) {
            *isBal = false;
        }
        return rightHeight;
    } else {
        int leftHeight = height(root->left, current + 1, isBal);
        int rightHeight = height(root->right, current + 1, isBal);
        if(leftHeight - rightHeight > 1 || rightHeight - leftHeight > 1){
            *isBal = false;
        }
        return leftHeight >= rightHeight ? leftHeight : rightHeight;
    }
}

bool isBalanced(struct TreeNode* root){
    if(root == NULL) {
        return true;
    }
    
    if(root->right == NULL && root->left == NULL) {
        return true;
    }
    
    bool isBal = true;
    int rootHeight = height(root, 0, &isBal);
    
    return isBal;
}

