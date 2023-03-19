//
// Created by tcy on 2023/3/19.
//
#include <string>
#include <iostream>
using namespace std;
/**
 * 这是一个 C++ 的 Solution 类，其中包含了一个名为 lengthOfLongestSubstring 的函数，函数的作用是在给定的字符串中寻找不包含重复字符的最长子串，并返回子串的长度。

实现的思路是使用双指针的方法，一个指针 i 指向当前子串的起始位置，
 另一个指针 rk 指向当前子串的末尾位置。同时使用一个长度为 255 的布尔数组 flag 来表示每个字符是否在当前子串中出现过。
 在遍历字符串的过程中，使用 i 和 rk 分别维护当前子串的起始和末尾位置，并动态更新布尔数组 flag。

具体实现方式是，首先将指针 rk 初始值设为 -1，表示当前子串为空。
 然后开始遍历字符串，对于每个位置 i，将 flag[s[i-1]] 设置为 0（这里需要注意 i=0 的情况需要特判，
 因为此时 si-1 是越界的）。然后使用 while 循环，将指针 rk 向右移动，
 直到当前子串中出现了重复字符或者末尾指针 rk 到达字符串的末尾位置，同时在每次移动时更新布尔数组 flag。
 然后将当前子串的长度 rk-i+1 与之前计算的最长子串长度 result 比较，取其中的最大值，并将其存入 result 中。
 最终遍历完成后，函数返回 result。

需要注意的是，在每次移动末尾指针 rk 时，需要使用数组 flag 来判断当前字符是否出现过。
 如果 flag[s[rk+1]] == 0，说明当前字符还没有出现过，可以将其加入到当前子串中；否则说明当前字符已经出现过，
 需要缩小子串的长度，移动指针 i 来去除子串的起始位置的重复字符。
 */
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int result = 0;
        int rk = -1;
        int flag[255];
        size_t n = s.size();
        for (int & i : flag) {
            i = 0;
        }
        for (int i = 0; i < s.size(); ++i) {
            if (i != 0){
                flag[s[i-1]] = 0;
            }
            while (rk+1 < n && flag[s[rk+1]] == 0){
                flag[s[rk+1]] = 1;
                rk++;
            }
            result = max(result, rk-i+1);
        }
        return result;
    }
};

int main()
{
    string s = "bbbbb";
    Solution solution;
    cout << solution.lengthOfLongestSubstring(s);
}
