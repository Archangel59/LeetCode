#include<iostream>
using namespace std;

class Solution {
public:
    int next[50010],n,m;
    void getNext(string str) {
        int i = -1,j = 0;
        next[0] = -1;
        while(j < m) {
            if(i == -1 or str[i] == str[j]) {
                i++,j++;
                next[j] = i;
            } else {
                i = next[i];
            }
        }
    }
    int strStr(string haystack, string needle) {
        n = haystack.size(),m = needle.size();
        if(!m)  return 0;

        getNext(needle);
        // kmp
        int i = 0,j = 0;
        while(j < n) {
            if(i == -1 or needle[i] == haystack[j]) {
                i++,j++;
            } else {
                i = next[i];
            }
            if (i == m) {
                return j-i;
            }
        }
        return -1;
    }
};

int main () {
    Solution a;
    cout << a.strStr("hello", "ll") << endl;
}