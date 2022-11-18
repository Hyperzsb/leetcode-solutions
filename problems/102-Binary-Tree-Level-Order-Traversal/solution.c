/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int** levelOrder(struct TreeNode* root, int* returnSize, int** returnColumnSizes){
    if(root == NULL) {
        *returnSize = 0;
        return NULL;
    }
    
    struct TreeNode **nodes = 
        (struct TreeNode **)malloc(2010 * sizeof(struct TreeNode *));
    int *heightSize = (int *)malloc(2010 * sizeof(int));
    memset(heightSize, 0, 2010 * sizeof(int));
    
    int currentHeight = 0;
    int index = 0;
    int len = 0;
    
    nodes[len++] = root;
    nodes[index]->val += 1000;
    
    while(index < len) {
        currentHeight = nodes[index]->val / 10000;
        heightSize[currentHeight]++;
        if(nodes[index]->left != NULL) {
            nodes[index]->left->val += (currentHeight + 1) * 10000 + 1000;
            nodes[len++] = nodes[index]->left;
        }
        if(nodes[index]->right != NULL) {
            nodes[index]->right->val += (currentHeight + 1) * 10000 + 1000;
            nodes[len++] = nodes[index]->right;
        }
        index++;
    }
    
    *returnSize = nodes[len - 1]->val / 10000 + 1;
    
    *returnColumnSizes = (int *)malloc(*returnSize * sizeof(int));
    int **returnVals = (int **)malloc(*returnSize * sizeof(int *));
    
    index = 0;
    for(int i = 0; i < *returnSize; i++) {
        (*returnColumnSizes)[i] = heightSize[i];
        returnVals[i] = (int *)malloc(heightSize[i] * sizeof(int));
        
        for(int j = 0; j < heightSize[i]; j++) {
            returnVals[i][j] = nodes[index + j]->val % 10000 - 1000;
        }
        
        index += heightSize[i];
    }
    
    return returnVals;
}

