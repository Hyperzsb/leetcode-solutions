/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
bool isValidBST(struct TreeNode* root){
    if(root == NULL) {
        return true;
    }
    
    if(root->left == NULL && root->right == NULL) {
        return true;
    }
    
    struct TreeNode **nodes = 
        (struct TrreeNode **)malloc(10010 * sizeof(struct TreeNode *));
    memset(nodes, 0, 10010 * sizeof(struct TreeNode *));
    
    int *sides = (int *)malloc(10010 * sizeof(int));
    memset(sides, 0, 10010 * sizeof(int));
    
    int currentHeight = 0;
    nodes[currentHeight] = root;
    sides[currentHeight] = 0;
    
    while(currentHeight >= 0) {
        if(sides[currentHeight] == 0) {
            sides[currentHeight] = 1;
            if(nodes[currentHeight]->left != NULL) {
                currentHeight++;
                nodes[currentHeight] = nodes[currentHeight - 1]->left;
                if(nodes[currentHeight]->val >= nodes[currentHeight - 1]->val) {
                    return false;
                }
                for(int i = 2; i <= currentHeight; i++) {
                    if(sides[currentHeight - i] == 2) {
                        if(nodes[currentHeight]-> val <= nodes[currentHeight - i]->val) {
                            return false;
                        }
                    }
                }
            }
        } else if (sides[currentHeight] == 1) {
            sides[currentHeight] = 2;
            if(nodes[currentHeight]->right != NULL) {
                currentHeight++;
                nodes[currentHeight] = nodes[currentHeight - 1]->right;
                if(nodes[currentHeight]->val <= nodes[currentHeight - 1]->val) {
                    return false;
                }
                for(int i = 2; i <= currentHeight; i++) {
                    if(sides[currentHeight - i] == 1) {
                        if(nodes[currentHeight]-> val >= nodes[currentHeight - i]->val) {
                            return false;
                        }
                    }
                }
            }
        } else {
            sides[currentHeight] = 0;
            currentHeight--;
        }
    }
    
    return true;
}

