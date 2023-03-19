//
// Created by tcy on 2023/3/19.
//
#include <iostream>
#include <map>

class Solution {
private:
    std::map<int ,int> data;
public:
    int climbStairs(int n) {
        if (n==1)
            return 1;
        if (n == 2)
            return 2;
        if (n==3)
            return 3;
        if (data.count(n)){
            return data[n];
        } else{
            data[n] = climbStairs(n-1)+climbStairs(n-2);
        }
        return data[n];
    }
};
int main()
{
    Solution solution;
    std::cout << solution.climbStairs(45);
}
