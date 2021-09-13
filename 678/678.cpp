#include<iostream>
#include<stack>
#include<list>
using namespace std;

class Solution {
public:
    bool checkValidString(string s) {
        int l = 0, r = 0;
        for(int i=0;i<s.size();i++)
            if(s[i] == '(') l++;
            else if(s[i] == ')') r++;
        list<int> L;
        for(int i=0;i<s.size();i++)
            if(s[i] == '*') L.push_back(i);

        // 使得l = r
        for(int i=0;i<l-r and !L.empty();i++) {
            s[L.back()] = ')'; L.pop_back();
        }
        for(int i=0;i<r-l and !L.empty();i++) {
            s[L.front()] = '('; L.pop_front();
        }
        // 给*进行处理
        while(L.size() > 1) {
            s[L.front()] = '(';
            s[L.back()] = ')';
            L.pop_front(); L.pop_back();
        }

        int res = 0;
        for(int i=0;i<s.size();i++)
            if(s[i] == '(') res++;
            else if (s[i] == ')') {
                if(res <= 0) {
                    return false;
                }
                res--;
            }

        return res == 0;
    }
};

int main () {
    Solution a;
    cout << a.checkValidString("(") << endl;
    return 0;
}