/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
bool isPalindrome(struct ListNode* head){
    struct ListNode *current = head;
    
    int size = 0;
    while(current != NULL) {
        size++;
        current = current->next;
    }
    
    if(size == 1) {
        return true;
    }
    
    int index = 0;
    struct ListNode *current1 = head, *current2 = head, *previous2 = NULL, *next2 = NULL;
    
    while(index < size / 2){
        previous2 = current2;
        current2 = current2->next;
        index++;
    }

    
    previous2->next = NULL;
    if(size % 2 == 1) {
         current2 = current2->next;
    }

    while(current2 != NULL) {
        next2 = current2->next;
        current2->next = previous2;
        previous2 = current2;
        current2 = next2;
    }

    current2 = previous2;
    while(current1 != NULL && current2 != NULL) {
        if(current1->val != current2->val) {
            return false;
        }
        current1 = current1->next;
        current2 = current2->next;
    }
    
    return true;
}

