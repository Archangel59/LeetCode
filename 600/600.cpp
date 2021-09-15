#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    int findIntegers(int n) {
        int f[34];
        f[0] = 1,f[1] = 2;
        for(int i=2;i<=32;i++) f[i] = f[i-1] + f[i-2];

        int ans = 0;
        for(int i=31;i>=0;i--) {
            if((n & (1<<i)) == 0) continue;
            ans += f[i];
            if(n & (1<<i+1)) return ans;
        }
        return ans+1;
    }
};
int main () {
    Solution a;
    cout << a.findIntegers(13) << endl;
    return 0;
}