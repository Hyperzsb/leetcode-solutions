/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* reverseList(struct ListNode* head) {
    struct ListNode **ptrList, *retHead = head, *current = head;
    int len = 0, index = 0;
    while(current) {
        len++;
        current = current->next;
    }
    
    if (len <= 1) {
        return head;
    }
    
    ptrList = (struct ListNode **)malloc(len * sizeof(struct ListNode *));
    current = head;
    
    while(current) {
        ptrList[index++] = current;
        current = current->next;
    }
    
    retHead = ptrList[len - 1];
    current = retHead;
    for (int i = len - 2; i >= 0; i--) {
        current->next = ptrList[i];
        current = current->next;
    }
    current->next = NULL;

    free(ptrList);

    return retHead;
}

