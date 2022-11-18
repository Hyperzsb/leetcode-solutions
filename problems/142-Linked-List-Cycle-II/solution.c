/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode *detectCycle(struct ListNode *head) {
    struct ListNode *current = head;
    int len = 0;
    
    while(current) {
        if(current->val >= 900000) {
            break;
        }
        len++;
        current->val += 1000000;
        current = current->next;
    }
    
    current = head;
    for(int i = 0; i < len; i++) {
        current->val -= 1000000;
        current = current->next;
    }
    
    return current;
}

