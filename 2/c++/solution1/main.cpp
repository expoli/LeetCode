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

struct ListNode {
    int val;
    ListNode *next;

    ListNode() : val(0), next(nullptr) {}

    ListNode(int x) : val(x), next(nullptr) {}

    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

#include <iostream>

class Solution {
public:
    ListNode *addTwoNumbers(ListNode *l1, ListNode *l2) {
        ListNode *result = nullptr;
        ListNode *rear = result;
        int up = 0;
        while (l1 != nullptr || l2 != nullptr) {
            int sum = 0, r=0;
            if (l1 != nullptr) {
                sum += l1->val;
                l1 = l1->next;
            }
            if (l2 != nullptr) {
                sum += l2->val;
                l2 = l2->next;
            }
            sum += up;
            r = sum % 10;
            up = sum / 10;

            if (result == nullptr) {    // 一开始
                auto *temp = new(ListNode);
                temp->val = r;
                result = temp;
                rear = temp;
                continue;
            } else {
                auto *temp = new(ListNode);
                temp->val = r;
                rear->next = temp;
                rear = rear->next;
            }
        }
        if (up != 0) {
            auto *temp = new(ListNode);
            temp->val = up;
            rear->next = temp;
        }
        return result;
    }
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
        }
        rear->next = temp;
        rear = rear->next;
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

int main() {
//    int nums1[] = {2,4,3};
//    int nums2[] = {5,6,4};

    int nums1[] = {9, 9, 9, 9, 9, 9, 9};
    int nums2[] = {9, 9, 9, 9};

    auto n1 = sizeof(nums1) / sizeof(int);
    auto n2 = sizeof(nums2) / sizeof(int);
    auto list1 = createList(nums1, n1);
    auto list2 = createList(nums2, n2);

    Solution solution;
    auto temp = solution.addTwoNumbers(list1, list2);
    showList(temp);

    delete (list1);
    delete (list2);
    delete (temp);
}
