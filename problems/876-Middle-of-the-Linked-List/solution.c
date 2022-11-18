/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* middleNode(struct ListNode* head){
    struct ListNode *current = head;
    int len = 0;
    
    while(current) {
        len++;
        current = current->next;
    }
    
    current = head;
    for(int i = 0; i < len / 2; i++) {
        current = current->next;
    }
    
    return current;
}

