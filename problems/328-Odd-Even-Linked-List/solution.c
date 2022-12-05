/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* oddEvenList(struct ListNode* head){
    if(head == NULL) {
        return NULL;
    }
    
    if(head->next == NULL) {
        return head;
    }
    
    struct ListNode *oddHead = head, *oddCurrent = oddHead, 
    *evenHead = head->next, *evenCurrent = evenHead, *current = head->next->next;
    
    oddCurrent->next = NULL;
    evenCurrent->next = NULL;
    struct ListNode *temp = NULL;
    
    int isOdd = 1;
    while(current) {
        if(isOdd) {
            temp = current;
            current = current->next;
            
            oddCurrent->next = temp;
            oddCurrent = oddCurrent->next;
            oddCurrent->next = NULL;
            
            isOdd = 0;
        } else {
            temp = current;
            current = current->next;
            
            evenCurrent->next = temp;
            evenCurrent = evenCurrent->next;
            evenCurrent->next = NULL;
            
            isOdd = 1;
        }
    }
    
    oddCurrent->next = evenHead;
    
    return oddHead;
}

