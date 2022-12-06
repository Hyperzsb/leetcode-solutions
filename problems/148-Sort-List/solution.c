/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* sortList(struct ListNode* head){
    if(head == NULL || head->next == NULL) {
        return head;
    }
    
    if(head->next->next == NULL) {
        struct ListNode *temp = head;
        if(head->val > head->next->val) {
            head->next->next = head;
            temp = head->next;
            head->next = NULL;
        }
        
        return temp;
    }
    
    struct ListNode *current = head, *halfHead = head;
    
    int count = 0;
    
    while(current->next) {
        if(count) {
            current = current->next;
            halfHead = halfHead->next;
            count = 0;
        } else {
            current = current->next;
            count = 1;
        }
    }
    
    struct ListNode *temp = halfHead->next;
    halfHead->next = NULL;
    halfHead = temp;
    
    head = sortList(head);
    halfHead = sortList(halfHead);
    
    current = head;
    struct ListNode *halfCurrent = halfHead;
    
    if(current->val <= halfCurrent->val) {
        head = current;
        temp = head;
        current = current->next;
    } else {
        head = halfCurrent;
        temp = head;
        halfCurrent = halfCurrent->next;
    }
    
    while(current && halfCurrent) {
        if(current->val <= halfCurrent->val) {
            temp->next = current;
            temp = temp->next;
            current = current->next;
        } else {
            temp->next = halfCurrent;
            temp = temp->next;
            halfCurrent = halfCurrent->next;
        }
    }
    
    while(current || halfCurrent) {
        if(current) {
            temp->next = current;
            temp = temp->next;
            current = current->next;
        } else {
            temp->next = halfCurrent;
            temp = temp->next;
            halfCurrent = halfCurrent->next;
        }
    }
    
    return head;
}

