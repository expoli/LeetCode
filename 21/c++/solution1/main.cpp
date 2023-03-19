//
// Created by tcy on 2023/3/19.
//

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
#include <iostream>

struct ListNode {
    int val;
    ListNode *next;

    ListNode() : val(0), next(nullptr) {}

    ListNode(int x) : val(x), next(nullptr) {}

    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode *createList(const int nums[], int n) {
    ListNode *result = nullptr;
    ListNode *rear = nullptr;
    for (int i = 0; i < n; ++i) {
        auto *temp = new(ListNode);
        temp->val = nums[i];
        if (result == nullptr) {
            result = temp;
            rear = result;
        } else{
            rear->next = temp;
            rear = rear->next;
        }
    }
    return result;
}

void showList(ListNode *num) {
    ListNode *temp = num;
    while (temp != nullptr) {
        std::cout << temp->val << ", ";
        temp = temp->next;
    }
}

class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        if (!list1){
            return list2;
        }
        if (!list2){
            return list1;
        }
        ListNode *head = list1;
        ListNode *pre = list1;
        while (list1 && list2){
            if (list1->val <= list2->val){
                pre = list1;
                list1 = list1->next;
            }else if (list1 == pre){
                ListNode *temp = list2;
                list2 = list1;
                list1 = temp;
                head = list1;
            } else{
                pre->next = list2;
                list2 = list2->next;
                pre = pre->next;
                pre->next=list1;
            }
        }
        if (list2){
            pre->next = list2;
        }
        return head;
    }
};

int main()
{
    int l1[] = {2};
    int l2[] = {1};

    int num1 = sizeof l1 / sizeof (int );
    int num2 = sizeof l2 / sizeof (int );

    auto list1 = createList(l1,num1);
    auto list2 = createList(l2,num2);

    Solution solution;
    auto result = solution.mergeTwoLists(list1,list2);

    showList(result);
}
