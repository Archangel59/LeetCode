#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<int> v[20];
    struct node {
        int num, idx;
    };
    int numWays(int n, vector<vector<int>>& relation, int k) {
        for(int i=0;i<relation.size();i++) {
            v[relation[i][0]].push_back(relation[i][1]);
        }
        int ans = 0;
        node a;a.idx = 0, a.num = 0;
        queue<node> q;
        q.push(a);
        while(!q.empty()) {
            a = q.front();q.pop();
            if(a.idx > k) break;
            if(a.idx == k && a.num == n-1) ans++;
            for(int i=0;i<v[a.num].size();i++) {
                node b;
                b.idx = a.idx + 1;
                b.num = v[a.num][i];
                q.push(b);
            }
        }

        return ans;
    }
};

int main () {
    Solution a;
    vector<vector<int> > v = {{0,2},{2,1},{3,4},{2,3},{1,4},{2,0},{0,4}};
    cout << a.numWays(5, v, 3) << endl;


    return 0;
}
