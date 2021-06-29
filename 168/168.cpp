#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    string convertToTitle(int columnNumber) {
        string ans;
        while (columnNumber) {
            int a0 = (columnNumber - 1) % 26 + 1;
            ans += a0 + 'A' - 1;
            columnNumber = (columnNumber - a0) / 26;
        }
        reverse(ans.begin(), ans.end());
        return ans;
    }
};

int main () {
    Solution a;
    cout << a.convertToTitle(701) << endl;

    return 0;
}