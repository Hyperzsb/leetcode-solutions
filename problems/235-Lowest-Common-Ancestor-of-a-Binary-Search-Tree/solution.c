/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */

struct TreeNode* lowestCommonAncestor(struct TreeNode* root, struct TreeNode* p, struct TreeNode* q) {
    struct TreeNode *currentNode = root, *temp;
    
    if(p->val > q->val) {
        temp = p;
        p = q;
        q = temp;
    }
    
    while(currentNode->left || currentNode->right) {
        if(currentNode->val == p->val || currentNode->val == q->val) {
            return currentNode;
        }
        if(currentNode->val > p->val && currentNode->val < q->val) {
            return currentNode;
        }
        if(currentNode->val > p->val && currentNode->val > q->val) {
            currentNode = currentNode->left;
        }
        if(currentNode->val < p->val && currentNode->val < q->val) {
            currentNode = currentNode->right;
        }
    }
    
    return NULL;
}

