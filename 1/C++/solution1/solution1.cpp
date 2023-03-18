//
// Created by tcy on 2023/3/18.
//

#include "solution1.h"
#include <iostream>
#include <map>
#include <vector>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        int temp = 0;
        vector<int> result;
        map<int, int> signal;
        for (int i = 0; i < nums.size(); ++i) {
            temp = target - nums[i];
            if (signal.find(nums[i]) != signal.end()){ // 已经找到了有效答案，可以放入结果进行返回
                result.insert(result.end(),i);
                result.insert(result.end(),signal[nums[i]]);
                break;
            }
            if (!signal[temp]){ // 不是有效答案，进行附值记录
                signal[temp] = i;
            }
        }
        return result;
    }
};

int main(){
    Solution solution;
    vector<int> nums = {2, 7, 11, 15};
    int target = 9;

    auto result = solution.twoSum(nums,target);

    for (auto temp:result) {
        cout << temp << " ";
    }
}
