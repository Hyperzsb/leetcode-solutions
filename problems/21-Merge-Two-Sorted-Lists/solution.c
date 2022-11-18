/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* mergeTwoLists(struct ListNode* list1, struct ListNode* list2){
    struct ListNode* node1 = list1, *node2 = list2, *head = NULL, *current = NULL;
    
    if (!node1 && !node2) {
        return NULL;
    } else if (!node1) {
        return node2;
    } else if (!node2) {
        return node1;
    } else {
        if (node1->val <= node2->val) {
            head = node1;
            current = head;
            node1 = node1->next;
        } else {
            head = node2;
            current = head;
            node2 = node2->next;
        }
    }
    
    while(node1 || node2) {
        if (!node1) {
            current->next = node2;
            return head;
        } else if (!node2) {
            current->next = node1;
            return head;
        } else {
            if (node1->val <= node2->val) {
                current->next = node1;
                current = current->next;
                node1 = node1->next;
            } else {
                current->next = node2;
                current = current->next;
                node2 = node2->next;
            }
        }
    }
    
    return head;
}

