#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    int maxIceCream(vector<int>& costs, int coins) {
        int vis[100010];
        memset(vis,0,sizeof(vis));
        for(int i=0;i<costs.size();i++)
            vis[costs[i]]++;

        int ans = 0;
        for(int i=1;i<=100000;i++) {
            if(i > coins) break;
            if(vis[i]) {
                int num = min(coins/i, vis[i]);
                ans += num;
                coins -= num * i;
            }
        }
        return ans;
    }
};

int main () {
    Solution a;
    vector<int> v = {1,1,2,3,4};
    cout << a.maxIceCream(v, 7) << endl;

    return 0;
}