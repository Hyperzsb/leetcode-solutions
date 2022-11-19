/**
 * Definition for a Node.
 * struct Node {
 *     int val;
 *     int numChildren;
 *     struct Node** children;
 * };
 */

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* preorder(struct Node* root, int* returnSize) {
    if(root == NULL) {
        *returnSize = 0;
        return NULL;
    }
    
    int *indexes = (int *)malloc(10010 * sizeof(int));
    memset(indexes, 0, 1004 * sizeof(int));
    
    struct Node **nodes = (struct Node **)malloc(10020 * sizeof(struct Node *));
    memset(nodes, 0, 10005 * sizeof(struct Node *));
    
    int *preorderVals = (int *)malloc(10030 * sizeof(int));
    memset(preorderVals, 0, 10005 * sizeof(int));
    
    *returnSize = 0;
    preorderVals[(*returnSize)++] = root->val;
    
    int currentHeight = 0;
    nodes[currentHeight] = root;
    indexes[currentHeight] = 0;
    
    while(currentHeight >= 0) {
        if(nodes[currentHeight]->numChildren > indexes[currentHeight]) {
            nodes[currentHeight + 1] = 
                nodes[currentHeight]->children[indexes[currentHeight]];
            indexes[currentHeight]++;
            currentHeight++;
        } else {
            currentHeight--;
            continue;
        }
        
        preorderVals[(*returnSize)++] = nodes[currentHeight]->val;
        
        if(nodes[currentHeight]->numChildren > 0) {
            indexes[currentHeight] = 0;
        } else {
            currentHeight--;
        }
    }

    free(indexes);
    free(nodes);

    return preorderVals;
}

