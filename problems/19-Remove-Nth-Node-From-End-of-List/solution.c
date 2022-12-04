/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* removeNthFromEnd(struct ListNode* head, int n){
    struct ListNode *current = head, *previous = NULL;
    
    int size = 0;
    while(current != NULL) {
        size++;
        current = current->next;
    }
    
    if(size - n == 0) {
        current = head->next;
    } else if(size - n == size - 1) {
        current = head;
        while(current->next->next != NULL) {
            current = current->next;
        }
        current->next = NULL;
        current = head;
    } else {
        int index = 0;
        previous = head;
        current = head->next;
        while(index < size - n - 1) {
            previous = current;
            current = current->next;
            index++;
        }
        previous->next = current->next;
        current = head;
    }
    
    return current;
}

