/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
int countNodes(struct TreeNode* root){
    if(root == NULL) {
        return 0;
    }
    
    struct TreeNode **nodes = (struct TreeNode **)malloc(50010 * sizeof(struct TreeNode *));
    int index = 0, len = 0;
    
    nodes[index] = root;
    len++;
    
    while(nodes[index]->left != NULL && nodes[index]->right != NULL) {
        nodes[len++] = nodes[index]->left;
        nodes[len++] = nodes[index]->right;
        index++;
    }
    
    if(nodes[index]->left != NULL) {
        len++;
    }
    
    return len;
}

